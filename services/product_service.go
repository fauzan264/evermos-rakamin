package services

import (
	"fmt"
	"os"
	"time"

	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"github.com/fauzan264/evermos-rakamin/repositories"
	"github.com/gosimple/slug"
)

type ProductService interface {
	GetProductByID(requestID request.GetByProductIDRequest) (response.ProductResponse, error)
	CreateProduct(requestUser request.GetByUserIDRequest, requestData request.ProductRequest) (response.ProductResponse, error)
	UpdateProduct(requestUser request.GetByUserIDRequest,  requestID request.GetByProductIDRequest, requestData request.ProductRequest) (response.ProductResponse, error)
	DeleteProduct(requestUser request.GetByUserIDRequest,  requestID request.GetByProductIDRequest) error
}

type productService struct {
	repository repositories.ProductRepository
	tokoRepository repositories.TokoRepository
	categoryRepository repositories.CategoryRepository
}

func NewProductService(
	repository repositories.ProductRepository,
	tokoRepository repositories.TokoRepository,
	categoryRepository repositories.CategoryRepository,
) *productService {
	return &productService{repository, tokoRepository, categoryRepository}
}

func (s *productService) GetProductByID(requestID request.GetByProductIDRequest) (response.ProductResponse, error) {
	getProduct, err := s.repository.GetProductByID(requestID.ID)
	if err != nil {
		return response.ProductResponse{}, err
	}

	tokoResponse := response.TokoResponse{
		ID : getProduct.Toko.ID,
		NamaToko : getProduct.Toko.NamaToko,
		URLFoto : getProduct.Toko.URLFoto,
	}

	categoryResponse := response.CategoryResponse{
		ID : getProduct.Category.ID,
		NamaCategory : getProduct.Category.NamaCategory,
	}

	productResponse := response.ProductResponse{
		ID : getProduct.ID,
		NamaProduk : getProduct.NamaProduk,
		Slug : getProduct.Slug,
		HargaReseller : getProduct.HargaReseller,
		HargaKonsumen : getProduct.HargaKonsumen,
		Stok : getProduct.Stok,
		Deskripsi : getProduct.Deskripsi,
		Toko : tokoResponse,
		Category : categoryResponse,
	}

	getPhotos, err := s.repository.GetPhotosProductByProductID(requestID.ID)
	if err != nil {
		return response.ProductResponse{}, err
	}

	var photos []response.PhotoProductResponse
	for _, getPhoto := range getPhotos {
		photo := response.PhotoProductResponse{
			ID : getPhoto.ID,
			IDProduk : getPhoto.IDProduk,
			URL : getPhoto.URL,
		}
		photos = append(photos, photo)
	}

	productResponse.Photos = photos

	return productResponse, nil
}

func (s *productService) CreateProduct(requestUser request.GetByUserIDRequest, requestData request.ProductRequest) (response.ProductResponse, error) {
	tx := s.repository.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	getToko, err := s.tokoRepository.GetTokoByUserID(requestUser.ID)
	if err != nil {
		tx.Rollback()
		return response.ProductResponse{}, err
	}

	product := model.Product{
		NamaProduk:    requestData.NamaProduk,
		HargaReseller: requestData.HargaReseller,
		HargaKonsumen: requestData.HargaKonsumen,
		Stok:          *requestData.Stok,
		Deskripsi:     requestData.Deskripsi,
		CreatedAt:     time.Now(),
		IDToko:        getToko.ID,
		IDCategory:    *requestData.IDCategory,
	}

	slugProduct := fmt.Sprintf("%s-%d", requestData.NamaProduk, time.Now().UnixNano()/int64(time.Millisecond))
	product.Slug = slug.Make(slugProduct)

	createProduct, err := s.repository.CreateProduct(tx, product)
	if err != nil {
		tx.Rollback()
		return response.ProductResponse{}, err
	}

	var photos []model.PhotoProduct
	for _, photo := range requestData.Photos {
		photoProduct := model.PhotoProduct{
			IDProduk:  createProduct.ID,
			URL:       photo.URL,
			CreatedAt: time.Now(),
		}
		photos = append(photos, photoProduct)
	}

	var createPhotos []model.PhotoProduct
	if len(photos) > 0 {
		createPhotos, err = s.repository.CreatePhotosProduct(tx, photos)
		if err != nil {
			tx.Rollback()
			return response.ProductResponse{}, err
		}
	}

	tx.Commit()

	logProduct := model.LogProduct{
		IDProduk: createProduct.ID,
		NamaProduk:    requestData.NamaProduk,
		Slug: slugProduct,
		HargaReseller: requestData.HargaReseller,
		HargaKonsumen: requestData.HargaKonsumen,
		Stock:         *requestData.Stok,
		Deskripsi:     requestData.Deskripsi,
		CreatedAt:     time.Now(),
		IDToko:        getToko.ID,
		IDCategory:    *requestData.IDCategory,
	}

	_, err = s.repository.CreateLogProduct(logProduct)
	if err != nil {
		return response.ProductResponse{}, err
	}

	var photosResponse []response.PhotoProductResponse
	for _, createPhoto := range createPhotos {
		photosResponse = append(photosResponse, response.PhotoProductResponse{
			ID:       createPhoto.ID,
			IDProduk: createPhoto.IDProduk,
			URL:      createPhoto.URL,
		})
	}

	tokoResponse := response.TokoResponse{
		ID:       getToko.ID,
		NamaToko: getToko.NamaToko,
		URLFoto:  getToko.URLFoto,
	}

	getCategory, _ := s.categoryRepository.GetCategoryByID(createProduct.IDCategory)

	categoryResponse := response.CategoryResponse{
		ID:           getCategory.ID,
		NamaCategory: getCategory.NamaCategory,
	}

	productResponse := response.ProductResponse{
		ID:            createProduct.ID,
		NamaProduk:    createProduct.NamaProduk,
		Slug:          createProduct.Slug,
		HargaReseller: createProduct.HargaReseller,
		HargaKonsumen: createProduct.HargaKonsumen,
		Stok:          createProduct.Stok,
		Deskripsi:     createProduct.Deskripsi,
		Toko:          tokoResponse,
		Category:      categoryResponse,
		Photos:        photosResponse,
	}

	return productResponse, nil
}

func (s *productService) UpdateProduct(
	requestUser request.GetByUserIDRequest,
	requestID request.GetByProductIDRequest,
	requestData request.ProductRequest,
) (response.ProductResponse, error) {
	product, err := s.repository.GetProductByID(requestID.ID)
	if err != nil {
		return response.ProductResponse{}, err
	}
	
	if product.Toko.IDUser != requestUser.ID {
		return response.ProductResponse{}, constants.ErrUnauthorized
	}

	if requestData.NamaProduk != "" {
		product.NamaProduk = requestData.NamaProduk
	}
	
	if requestData.IDCategory != nil {
		product.IDCategory = *requestData.IDCategory
	}
	
	if requestData.HargaReseller != "" {
		product.HargaReseller = requestData.HargaReseller
	}
	
	if requestData.HargaKonsumen != "" {
		product.HargaKonsumen = requestData.HargaKonsumen
	}
	
	if requestData.Stok != nil {
		product.Stok = *requestData.Stok
	}
	
	if requestData.Deskripsi != "" {
		product.Deskripsi = requestData.Deskripsi
	}

	product.UpdatedAt = time.Now()
	

	photos, err := s.repository.GetPhotosProductByProductID(product.ID)
	if err != nil {
		return response.ProductResponse{}, err
	}

	getToko, err := s.tokoRepository.GetTokoByUserID(requestUser.ID)
	if err != nil {
		return response.ProductResponse{}, err
	}

	slugProduct := fmt.Sprintf("%s-%d", requestData.NamaProduk, time.Now().UnixNano()/int64(time.Millisecond))
	product.Slug = slug.Make(slugProduct)

	tx := s.repository.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	_, err = s.repository.UpdateProduct(tx, product)
	if err != nil {
		tx.Rollback()
		return response.ProductResponse{}, err
	}

	var createdAt time.Time
	if len(photos) > 0 {
		createdAt = photos[0].CreatedAt
	} else {
		createdAt = time.Now()
	}

	var createPhotos []model.PhotoProduct
	if len(requestData.Photos) > 0 {
		if len(photos) > 0  {
			err = s.repository.DeleteAllPhotosByProductID(tx, product.ID)
			if err != nil {
				tx.Rollback()
				return response.ProductResponse{}, err
			}
	
			for _, existingPhoto := range photos {
				os.Remove(existingPhoto.URL)
			}
		}
	
		var newPhotos []model.PhotoProduct
		for _, photo := range requestData.Photos {
			photoProduct := model.PhotoProduct{
				IDProduk: product.ID,
				URL: photo.URL,
				CreatedAt: createdAt,
				UpdatedAt: time.Now(),
			}
			newPhotos = append(newPhotos, photoProduct)
		}
	
		createPhotos, err = s.repository.CreatePhotosProduct(tx, newPhotos)
		if err != nil {
			tx.Rollback()
			return response.ProductResponse{}, err
		}	
	}
	
	tx.Commit()

	logProduct := model.LogProduct{
		IDProduk: product.ID,
		NamaProduk: product.NamaProduk,
		Slug: slugProduct,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stock: product.Stok,
		Deskripsi: product.Deskripsi,
		CreatedAt: time.Now(),
		IDToko: getToko.ID,
		IDCategory: product.IDCategory,
	}

	_, err = s.repository.CreateLogProduct(logProduct)
	if err != nil {
		return response.ProductResponse{}, err
	}

	var photosResponse []response.PhotoProductResponse
	if len(requestData.Photos) > 0 {
		for _, createPhoto := range createPhotos {
			photosResponse = append(photosResponse, response.PhotoProductResponse{
				ID:       createPhoto.ID,
				IDProduk: createPhoto.IDProduk,
				URL:      createPhoto.URL,
			})
		}
	} else {
		if len(photos) > 0 {
			for _, photo := range photos {
				photosResponse = append(photosResponse, response.PhotoProductResponse{
					ID:       photo.ID,
					IDProduk: photo.IDProduk,
					URL:      photo.URL,
				})
			}
		}
	}

	tokoResponse := response.TokoResponse{
		ID:       getToko.ID,
		NamaToko: getToko.NamaToko,
		URLFoto:  getToko.URLFoto,
	}

	category, _ := s.categoryRepository.GetCategoryByID(product.IDCategory)

	categoryResponse := response.CategoryResponse{
		ID:           category.ID,
		NamaCategory: category.NamaCategory,
	}

	productResponse := response.ProductResponse{
		ID:            product.ID,
		NamaProduk:    product.NamaProduk,
		Slug:          product.Slug,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     product.Deskripsi,
		Toko:          tokoResponse,
		Category:      categoryResponse,
		Photos:        photosResponse,
	}

	return productResponse, nil
}

func (s *productService) DeleteProduct(
	requestUser request.GetByUserIDRequest,
	requestID request.GetByProductIDRequest,
) error {
	tx := s.repository.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	getProduct, err := s.repository.GetProductByID(requestID.ID)
	if err != nil {
		return err
	}

	if getProduct.Toko.IDUser != requestUser.ID {
		return constants.ErrUnauthorized
	}

	getPhotos, err := s.repository.GetPhotosProductByProductID(getProduct.ID)
	if err != nil {
		return  err
	}

	err = s.repository.DeleteAllPhotosByProductID(tx, getProduct.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	
	for _, existingPhoto := range getPhotos {
		os.Remove(existingPhoto.URL)
	}

	err = s.repository.DeleteProduct(tx, getProduct.ID)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}