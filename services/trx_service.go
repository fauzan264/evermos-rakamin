package services

import (
	"strconv"
	"time"

	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"github.com/fauzan264/evermos-rakamin/helpers"
	"github.com/fauzan264/evermos-rakamin/repositories"
)

type TRXService interface {
	GetListTRX(requestUser request.GetByUserIDRequest, requestData request.TRXListRequest) ([]response.TRXResponse, error)
	GetDetailTRX(requestUser request.GetByUserIDRequest, requestID request.GetByTRXIDRequest) (response.TRXResponse, error)
	CreateTRX(requestUser request.GetByUserIDRequest, requestData request.CreateTrxRequest) (response.TRXResponse, error)
}

type trxService struct {
	trxRepository repositories.TRXRepository
	productRepository repositories.ProductRepository
	addressRepository repositories.AlamatRepository
	shopRepository repositories.TokoRepository
	categoryRepository repositories.CategoryRepository
}

func NewTRXService(
	trxRepository repositories.TRXRepository,
	productRepository repositories.ProductRepository,
	addressRepository repositories.AlamatRepository,
	shopRepository repositories.TokoRepository,
	categoryRepository repositories.CategoryRepository,
) *trxService {
	return &trxService{
		trxRepository,
		productRepository,
		addressRepository,
		shopRepository,
		categoryRepository,
	}
}

func (s *trxService) GetListTRX(requestUser request.GetByUserIDRequest, requestData request.TRXListRequest) ([]response.TRXResponse, error) {
	page := requestData.Page
	limit := requestData.Limit
	name := requestData.Search
	
	listTRX, err := s.trxRepository.GetTRXByUserID(requestUser.ID, page, limit, name)
	if err != nil {
		return []response.TRXResponse{}, err
	}

	if len(listTRX) == 0 {
		return []response.TRXResponse{}, nil
	}

	listTRXResponse := response.ListTRXResponseFormatter(listTRX)

	return listTRXResponse, nil
}

func (s *trxService) GetDetailTRX(requestUser request.GetByUserIDRequest, requestID request.GetByTRXIDRequest) (response.TRXResponse, error) {
	trx, err := s.trxRepository.GetTRXUserByID(requestUser.ID, requestID.ID)
	if err != nil {
		return response.TRXResponse{}, err
	}

	trxResponse := response.TRXResponseFormatter(trx)

	return trxResponse, nil
}

func (s *trxService) CreateTRX(requestUser request.GetByUserIDRequest, requestData request.CreateTrxRequest) (response.TRXResponse, error) {
	var totalPrice int
	var logProducts []model.LogProduct
	var detailTRXs []model.DetailTRX
	for _, detailTRX := range requestData.DetailTrxRequest {
		product, err := s.productRepository.GetProductByID(requestUser.ID, detailTRX.IDProduk)
		if err != nil {
			return response.TRXResponse{}, err
		}

		if product.Stok < detailTRX.Kuantitas {
			return response.TRXResponse{}, constants.ErrInsufficient
		}
		
		logProduct := model.LogProduct{
			IDProduk: product.ID,
			NamaProduk: product.NamaProduk,
			Slug: product.Slug,
			HargaReseller: product.HargaReseller,
			HargaKonsumen: product.HargaKonsumen,
			Stock: detailTRX.Kuantitas,
			Deskripsi: product.Deskripsi,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IDToko: product.IDToko,
			IDCategory: product.IDCategory,
		}

		logProducts = append(logProducts, logProduct)
		hargaKonsumenInt, _ := strconv.Atoi(product.HargaKonsumen)
		price := hargaKonsumenInt * detailTRX.Kuantitas
		totalPrice += price

		detailTRX := model.DetailTRX{
			IDLogProduk: logProduct.ID,
			IDToko: logProduct.IDToko,
			Kuantitas: detailTRX.Kuantitas,
			HargaTotal: price,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),

			LogProduct: logProduct,
		}

		detailTRXs = append(detailTRXs, detailTRX)
	}

	invoiceNumber := helpers.GenerateInvoiceNumber()

	trx := model.TRX{
		IDUser: requestUser.ID,
		AlamatPengiriman: requestData.AlamatPengiriman,
		HargaTotal: totalPrice,
		KodeInvoice: invoiceNumber,
		MethodBayar: requestData.MethodBayar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		DetailTRX: detailTRXs,
	}

	createTRX, err := s.trxRepository.CreateTRX(trx)
	if err != nil {
		return response.TRXResponse{}, err
	}

	trxResponse := response.TRXResponseFormatter(createTRX)
	if trxResponse.ShippingAddress.ID == 0 {
		shippingAddress, _ := s.addressRepository.GetAlamatByID(trx.AlamatPengiriman)
		trxResponse.ShippingAddress.ID = shippingAddress.ID
		trxResponse.ShippingAddress.JudulAlamat = shippingAddress.JudulAlamat
		trxResponse.ShippingAddress.NamaPenerima = shippingAddress.NamaPenerima
		trxResponse.ShippingAddress.NoTelp = shippingAddress.NoTelp
		trxResponse.ShippingAddress.DetailAlamat = shippingAddress.DetailAlamat
	}

	var newDetailTRX []response.DetailTrx
	for _, detailTRX := range trxResponse.DetailTrx {
		if detailTRX.Toko.ID == 0 {
			shop, _ := s.shopRepository.GetTokoByID(trx.AlamatPengiriman)
			detailTRX.Toko.ID = shop.ID
			detailTRX.Toko.NamaToko = shop.NamaToko
			detailTRX.Toko.URLFoto = shop.URLFoto
		}

		if detailTRX.Product.Toko.ID == 0 {
			shop, _ := s.shopRepository.GetTokoByID(trx.AlamatPengiriman)
			detailTRX.Product.Toko.ID = shop.ID
			detailTRX.Product.Toko.NamaToko = shop.NamaToko
			detailTRX.Product.Toko.URLFoto = shop.URLFoto
		}

		// if detailTRX.Product.Category.ID == 0 {
		// 	category, _ := s.categoryRepository.GetCategoryByID(trx.DetailTRX.)
		// 	detailTRX.Product.Category.ID = category.ID
		// 	detailTRX.Product.Category.NamaCategory = category.NamaCategory
		// }
		// newDetailTRX = append(newDetailTRX, detailTRX)
	}
	trxResponse.DetailTrx = newDetailTRX

	return trxResponse, nil
}