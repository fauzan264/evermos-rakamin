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
	var updateProducts []model.Product
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

			Toko: product.Toko,
			Category: product.Category,
			Produk: product,
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
			Toko: product.Toko,
		}

		product.Stok -= detailTRX.Kuantitas

		detailTRXs = append(detailTRXs, detailTRX)
		updateProducts = append(updateProducts, product)
	}

	invoiceNumber := helpers.GenerateInvoiceNumber()

	shippingAddress, err := s.addressRepository.GetAlamatByID(requestData.AlamatPengiriman)
	if err != nil {
		return response.TRXResponse{}, err
	}

	trx := model.TRX{
		IDUser: requestUser.ID,
		AlamatPengiriman: requestData.AlamatPengiriman,
		HargaTotal: totalPrice,
		KodeInvoice: invoiceNumber,
		MethodBayar: requestData.MethodBayar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		Alamat: shippingAddress,
		DetailTRX: detailTRXs,		
	}

	createTRX, err := s.trxRepository.CreateTRX(trx)
	if err != nil {
		return response.TRXResponse{}, err
	}

	tx := s.productRepository.BeginTransaction()
	for _, updateProduct := range updateProducts {
		_, _ = s.productRepository.UpdateProduct(tx, updateProduct)
	}
	tx.Commit()

	trxResponse := response.TRXResponseFormatter(createTRX)

	return trxResponse, nil
}