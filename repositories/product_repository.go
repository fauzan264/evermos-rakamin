package repositories

import (
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	BeginTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB)
	RollbackTransaction(tx *gorm.DB)
	GetListProduct(requestSearch request.ProductListRequest) ([]model.Product, error)
	GetProductByID(userID int, id int) (model.Product, error)
	GetPhotosProductByProductID(id int) ([]model.PhotoProduct, error)
	CreateProduct(tx *gorm.DB, product model.Product) (model.Product, error)
	CreatePhotosProduct(tx *gorm.DB, photos []model.PhotoProduct) ([]model.PhotoProduct, error)
	CreateLogProduct(logProduct model.LogProduct) (model.LogProduct, error)
	UpdateProduct(tx *gorm.DB, product model.Product) (model.Product, error)
	DeleteProduct(tx *gorm.DB, id int) error
	DeleteAllPhotosByProductID(tx *gorm.DB, id int) error
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *productRepository) CommitTransaction(tx *gorm.DB) {
	tx.Commit()
}

func (r *productRepository) RollbackTransaction(tx *gorm.DB) {
	tx.Rollback()
}

func (r *productRepository) GetListProduct(requestSearch request.ProductListRequest) ([]model.Product, error) {
	var listProduct []model.Product

	offset := (requestSearch.Page - 1) * requestSearch.Limit

	query := r.db.Model(&model.Product{})

    if requestSearch.NameProduct != "" {
		query = query.Where("nama_produk LIKE ?", "%"+ requestSearch.NameProduct +"%")
	}

    if requestSearch.IDCategory != nil {
		query = query.Where("id_category = ?",  requestSearch.IDCategory)
	}

    if requestSearch.ShopID != nil {
		query = query.Where("id_toko = ?", requestSearch.ShopID)
	}

    if requestSearch.MinPrice != nil {
		query = query.Where("harga_konsumen >= ?", requestSearch.MinPrice)
	}

    if requestSearch.MaxPrice != nil {
		query = query.Where("harga_konsumen <= ?", requestSearch.MaxPrice)
	}

	query = query.Preload("Toko").
				Preload("Category").
				Preload("PhotosProduct")

	err := query.Limit(requestSearch.Limit).
				Offset(offset).
				Find(&listProduct).Error


	if err != nil {
		return listProduct, err
	}

	return listProduct, nil
}

func (r *productRepository) GetProductByID(userID int, id int) (model.Product, error) {
	var product model.Product

	err := r.db.Preload("Toko", "id_user = ?", userID).
				Preload("Category").
				Preload("PhotosProduct").
				Where("id = ?", id).
				First(&product).
				Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) GetPhotosProductByProductID(produkID int) ([]model.PhotoProduct, error) {
	var photosProduct []model.PhotoProduct

	err := r.db.Where("id_produk = ?", produkID).
				Find(&photosProduct).
				Error

	if err != nil {
		return photosProduct, err
	}

	return photosProduct, nil
}

func (r *productRepository) CreateProduct(tx *gorm.DB, product model.Product) (model.Product, error) {
	err := tx.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) CreatePhotosProduct(tx *gorm.DB, photos []model.PhotoProduct) ([]model.PhotoProduct, error) {
	err := tx.Create(&photos).Error
	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (r *productRepository) CreateLogProduct(logProduct model.LogProduct) (model.LogProduct, error) {
	err := r.db.Create(&logProduct).Error
	if err != nil {
		return logProduct, err
	}

	return logProduct, nil
}

func (r *productRepository) UpdateProduct(tx *gorm.DB, product model.Product) (model.Product, error) {
	err := tx.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) DeleteProduct(tx *gorm.DB, id int) error {
	var product model.Product
	err := tx.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) DeleteAllPhotosByProductID(tx *gorm.DB, id int) error {
	var photoProduct model.PhotoProduct
	err := tx.Where("id_produk = ?", id).Delete(&photoProduct).Error
	if err != nil {
		return err
	}

	return nil
}