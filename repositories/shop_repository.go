package repositories

import (
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"gorm.io/gorm"
)

type shopRepository struct {
	db *gorm.DB
}

type ShopRepository interface {
	CreateShop(shop model.Shop) error
	GetListShop(requestSearch request.ShopListRequest) ([]model.Shop, error)
	GetShopByID(id int) (model.Shop, error)
	GetShopByUserID(userID int) (model.Shop, error)
	UpdateShop(shop model.Shop) (model.Shop, error)
}

func NewShopRepository(db *gorm.DB) *shopRepository {
	return &shopRepository{db}
}

func (r *shopRepository) CreateShop(shop model.Shop) error {
	err := r.db.Create(&shop).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *shopRepository) GetListShop(requestSearch request.ShopListRequest) ([]model.Shop, error) {
	var listShop []model.Shop

	offset := (requestSearch.Page - 1) * requestSearch.Limit

	query := r.db.Model(&model.Shop{})

	if requestSearch.Name != "" {
		query = query.Where("nama_toko LIKE ?", "%"+requestSearch.Name+"%")
	}

	err := query.Limit(requestSearch.Limit).
				Offset(offset).
				Find(&listShop).Error

	if err != nil {
		return nil, err
	}

	return listShop, nil
}

func (r *shopRepository) GetShopByID(id int) (model.Shop, error) {
	var shop model.Shop
	err := r.db.Where("id = ?", id).First(&shop).Error
	if err != nil {
		return shop, err
	}

	return shop, nil
}

func (r *shopRepository) GetShopByUserID(userID int) (model.Shop, error) {
	var shop model.Shop
	err := r.db.Where("id_user = ?", userID).First(&shop).Error
	if err != nil {
		return shop, err
	}
	
	return shop, nil
}

func (r *shopRepository) UpdateShop(shop model.Shop) (model.Shop, error) {
	err := r.db.Save(&shop).Error
	if err != nil {
		return shop, err
	}

	return shop, nil
}