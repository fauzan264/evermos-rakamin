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
	tx := s.repository.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()


	getProduct, err := s.repository.GetProductByID(requestID.ID)
	if err != nil {
		return response.ProductResponse{}, err
	}
	
	if getProduct.Toko.IDUser != requestUser.ID {
		return response.ProductResponse{}, constants.ErrUnauthorized
	}

	if requestData.NamaProduk == "" {
		requestData.NamaProduk = getProduct.NamaProduk
	}
	
	if requestData.IDCategory == nil {
		requestData.IDCategory = &getProduct.IDCategory
	}
	
	if requestData.HargaReseller == "" {
		requestData.HargaReseller = getProduct.HargaReseller
	}
	
	if requestData.HargaKonsumen == "" {
		requestData.HargaKonsumen = getProduct.HargaKonsumen
	}
	
	if requestData.Stok == nil {
		requestData.Stok = &getProduct.Stok
	}
	
	if requestData.Deskripsi == "" {
		requestData.Deskripsi = getProduct.Deskripsi
	}
	

	getPhotos, err := s.repository.GetPhotosProductByProductID(getProduct.ID)
	if err != nil {
		return response.ProductResponse{}, err
	}

	getToko, err := s.tokoRepository.GetTokoByUserID(requestUser.ID)
	if err != nil {
		tx.Rollback()
		return response.ProductResponse{}, err
	}

	product := model.Product{
		ID: getProduct.ID,
		NamaProduk: requestData.NamaProduk,
		HargaReseller: requestData.HargaReseller,
		HargaKonsumen: requestData.HargaKonsumen,
		Stok: *requestData.Stok,
		Deskripsi: requestData.Deskripsi,
		CreatedAt: getProduct.CreatedAt,
		UpdatedAt: time.Now(),
		IDToko: getProduct.IDToko,
		IDCategory: *requestData.IDCategory,
	}

	slugProduct := fmt.Sprintf("%s-%d", requestData.NamaProduk, time.Now().UnixNano()/int64(time.Millisecond))
	product.Slug = slug.Make(slugProduct)

	updateProduct, err := s.repository.UpdateProduct(tx, product)
	if err != nil {
		tx.Rollback()
		return response.ProductResponse{}, err
	}

	var photos []model.PhotoProduct
	for _, photo := range requestData.Photos {
		photoProduct := model.PhotoProduct{
			IDProduk: getProduct.ID,
			URL: photo.URL,
			CreatedAt: getPhotos[0].CreatedAt,
			UpdatedAt: time.Now(),
		}
		photos = append(photos, photoProduct)
	}

	err = s.repository.DeleteAllPhotosByProductID(tx, product.ID)
	if err != nil {
		tx.Rollback()
		return response.ProductResponse{}, err
	}

	var createPhotos []model.PhotoProduct
	if len(photos) > 0 {
		createPhotos, err = s.repository.CreatePhotosProduct(tx, photos)
		if err != nil {
			tx.Rollback()
			return response.ProductResponse{}, err
		}

		for _, existingPhoto := range getPhotos {
			fmt.Println(existingPhoto.URL)
			os.Remove(existingPhoto.URL)
		}
	}

	tx.Commit()

	logProduct := model.LogProduct{
		IDProduk: getProduct.ID,
		NamaProduk:    requestData.NamaProduk,
		Slug: slugProduct,
		HargaReseller: requestData.HargaReseller,
		HargaKonsumen: requestData.HargaKonsumen,
		Stock:          *requestData.Stok,
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

	getCategory, _ := s.categoryRepository.GetCategoryByID(updateProduct.IDCategory)

	categoryResponse := response.CategoryResponse{
		ID:           getCategory.ID,
		NamaCategory: getCategory.NamaCategory,
	}

	productResponse := response.ProductResponse{
		ID:            updateProduct.ID,
		NamaProduk:    updateProduct.NamaProduk,
		Slug:          updateProduct.Slug,
		HargaReseller: updateProduct.HargaReseller,
		HargaKonsumen: updateProduct.HargaKonsumen,
		Stok:          updateProduct.Stok,
		Deskripsi:     updateProduct.Deskripsi,
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
		fmt.Println(existingPhoto.URL)
		os.Remove(existingPhoto.URL)
	}

	err = s.repository.DeleteProduct(tx, getProduct.ID)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}