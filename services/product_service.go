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
	GetListProduct(requestUser request.GetByUserIDRequest, requestSearch request.ProductListRequest) ([]response.ProductResponse, error)
	GetProductByID(requestUser request.GetByUserIDRequest, requestID request.GetByProductIDRequest) (response.ProductResponse, error)
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

func (s *productService) GetListProduct(requestID request.GetByUserIDRequest, requestSearch request.ProductListRequest) ([]response.ProductResponse, error) {
	listProducts, err := s.repository.GetProductsByUserID(requestID.ID, requestSearch)
	if err != nil {
		return []response.ProductResponse{}, err
	}

	if len(listProducts) == 0 {
		return []response.ProductResponse{}, nil
	}
	
	var responseListProduct []response.ProductResponse
	for _, product := range listProducts {
		productShop := response.TokoResponse{
			ID: product.Toko.ID,
			NamaToko: product.Toko.NamaToko,
			URLFoto: product.Toko.URLFoto,
		}

		productCategory := response.CategoryResponse{
			ID: product.Category.ID,
			NamaCategory: product.Category.NamaCategory,
		}

		var productPhotos []response.PhotoProductResponse
		for _, photo := range product.PhotosProduct {
			photoProduct := response.PhotoProductResponse{
				ID: photo.ID,
				IDProduk: photo.IDProduk,
				URL: photo.URL,
			}

			productPhotos = append(productPhotos, photoProduct)
		}

		responseProduct := response.ProductResponse{
			ID: product.ID,
			NamaProduk: product.NamaProduk,
			Slug: product.Slug,
			HargaReseller: product.HargaReseller,
			HargaKonsumen: product.HargaKonsumen,
			Stok: product.Stok,
			Deskripsi: product.Deskripsi,
			Toko: productShop,
			Category: productCategory,
			Photos: productPhotos,
		}
	
		responseListProduct = append(responseListProduct, responseProduct)
	}


	return responseListProduct, nil
}

func (s *productService) GetProductByID(requestUser request.GetByUserIDRequest, requestID request.GetByProductIDRequest) (response.ProductResponse, error) {
	product, err := s.repository.GetProductByID(requestUser.ID, requestID.ID)
	if err != nil {
		return response.ProductResponse{}, err
	}

	productShop := response.TokoResponse{
		ID: product.Toko.ID,
		NamaToko: product.Toko.NamaToko,
		URLFoto: product.Toko.URLFoto,
	}

	productCategory := response.CategoryResponse{
		ID: product.Category.ID,
		NamaCategory: product.Category.NamaCategory,
	}

	var productPhotos []response.PhotoProductResponse
	for _, photo := range product.PhotosProduct {
		photoProduct := response.PhotoProductResponse{
			ID: photo.ID,
			IDProduk: photo.IDProduk,
			URL: photo.URL,
		}

		productPhotos = append(productPhotos, photoProduct)
	}

	responseProduct := response.ProductResponse{
		ID: product.ID,
		NamaProduk: product.NamaProduk,
		Slug: product.Slug,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok: product.Stok,
		Deskripsi: product.Deskripsi,
		Toko: productShop,
		Category: productCategory,
		Photos: productPhotos,
	}

	return responseProduct, nil
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
	product, err := s.repository.GetProductByID(requestUser.ID, requestID.ID)
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

	getProduct, err := s.repository.GetProductByID(requestUser.ID, requestID.ID)
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