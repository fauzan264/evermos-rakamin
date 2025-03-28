package services

import (
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/repositories"
)

type TRXService interface {
	GetListTRX(requestUser request.GetByUserIDRequest, requestData request.TRXListRequest) ([]response.TRXResponse, error)
	GetDetailTRX(requestUser request.GetByUserIDRequest, requestID request.GetByTRXIDRequest) (response.TRXResponse, error)
}

type trxService struct {
	repository repositories.TRXRepository
}

func NewTRXService(repository repositories.TRXRepository) *trxService {
	return &trxService{repository}
}

func (s *trxService) GetListTRX(requestUser request.GetByUserIDRequest, requestData request.TRXListRequest) ([]response.TRXResponse, error) {
	page := requestData.Page
	limit := requestData.Limit
	name := requestData.Search
	
	listTRX, err := s.repository.GetTRXByUserID(requestUser.ID, page, limit, name)
	if err != nil {
		return []response.TRXResponse{}, err
	}

	if len(listTRX) == 0 {
		return []response.TRXResponse{}, nil
	}

	var responseListTRX []response.TRXResponse
	for _, trx := range listTRX {
		trxAlamat := response.AddressResponse{
			ID: trx.Alamat.ID,
			JudulAlamat: trx.Alamat.JudulAlamat,
			NamaPenerima: trx.Alamat.NamaPenerima,
			NoTelp: trx.Alamat.NoTelp,
			DetailAlamat: trx.Alamat.DetailAlamat,
		}

		var listDetailTRX []response.DetailTrx
		for _, detailTRX := range trx.DetailTRX {

			productShop := response.TokoResponse{
				ID: detailTRX.LogProduct.Toko.ID,
				NamaToko: detailTRX.LogProduct.Toko.NamaToko,
				URLFoto: detailTRX.LogProduct.Toko.URLFoto,
			}

			productCategory := response.CategoryResponse{
				ID: detailTRX.LogProduct.Category.ID,
				NamaCategory: detailTRX.LogProduct.Category.NamaCategory,
				CreatedAt: &detailTRX.LogProduct.Category.CreatedAt,
				UpdatedAt: &detailTRX.LogProduct.Category.UpdatedAt,
			}

			var productPhotos []response.PhotoProductResponse
			for _, photo := range detailTRX.LogProduct.Produk.PhotosProduct {
				productPhoto := response.PhotoProductResponse{
					ID: photo.ID,
					IDProduk: photo.IDProduk,
					URL: photo.URL,
				}

				productPhotos = append(productPhotos, productPhoto)
			}

			product := response.LogProductResponse{
				ID: detailTRX.LogProduct.ID,
				NamaProduk: detailTRX.LogProduct.NamaProduk,
				Slug: detailTRX.LogProduct.Slug,
				HargaReseller: detailTRX.LogProduct.HargaReseller,
				HargaKonsumen: detailTRX.LogProduct.HargaKonsumen,
				Deskripsi: detailTRX.LogProduct.Deskripsi,
				Toko: productShop,
				Category: productCategory,
				Photos: productPhotos,
			}

			shop := response.TokoResponse{
				ID: detailTRX.Toko.ID,
				NamaToko: detailTRX.Toko.NamaToko,
				URLFoto: detailTRX.Toko.URLFoto,
			}

			dataDetailTRX := response.DetailTrx{
				Product: product,
				Toko: shop,
				Kuantitas: detailTRX.Kuantitas,
				HargaTotal: detailTRX.HargaTotal,
			}

			listDetailTRX = append(listDetailTRX, dataDetailTRX)
		}


		responseTRX := response.TRXResponse{
			ID: trx.ID,
			HargaTotal: trx.HargaTotal,
			KodeInvoice: trx.KodeInvoice,
			MethodBayar: trx.MethodBayar,
			ShippingAddress: trxAlamat,
			DetailTrx: listDetailTRX,
		}

		responseListTRX = append(responseListTRX, responseTRX)
	}

	return responseListTRX, nil
}

func (s *trxService) GetDetailTRX(requestUser request.GetByUserIDRequest, requestID request.GetByTRXIDRequest) (response.TRXResponse, error) {
	trx, err := s.repository.GetTRXUserByID(requestUser.ID, requestID.ID)
	if err != nil {
		return response.TRXResponse{}, err
	}

	trxAlamat := response.AddressResponse{
		ID: trx.Alamat.ID,
		JudulAlamat: trx.Alamat.JudulAlamat,
		NamaPenerima: trx.Alamat.NamaPenerima,
		NoTelp: trx.Alamat.NoTelp,
		DetailAlamat: trx.Alamat.DetailAlamat,
	}

	var listDetailTRX []response.DetailTrx
	for _, detailTRX := range trx.DetailTRX {

		productShop := response.TokoResponse{
			ID: detailTRX.LogProduct.Toko.ID,
			NamaToko: detailTRX.LogProduct.Toko.NamaToko,
			URLFoto: detailTRX.LogProduct.Toko.URLFoto,
		}

		productCategory := response.CategoryResponse{
			ID: detailTRX.LogProduct.Category.ID,
			NamaCategory: detailTRX.LogProduct.Category.NamaCategory,
			CreatedAt: &detailTRX.LogProduct.Category.CreatedAt,
			UpdatedAt: &detailTRX.LogProduct.Category.UpdatedAt,
		}

		var productPhotos []response.PhotoProductResponse
		for _, photo := range detailTRX.LogProduct.Produk.PhotosProduct {
			productPhoto := response.PhotoProductResponse{
				ID: photo.ID,
				IDProduk: photo.IDProduk,
				URL: photo.URL,
			}

			productPhotos = append(productPhotos, productPhoto)
		}

		product := response.LogProductResponse{
			ID: detailTRX.LogProduct.ID,
			NamaProduk: detailTRX.LogProduct.NamaProduk,
			Slug: detailTRX.LogProduct.Slug,
			HargaReseller: detailTRX.LogProduct.HargaReseller,
			HargaKonsumen: detailTRX.LogProduct.HargaKonsumen,
			Deskripsi: detailTRX.LogProduct.Deskripsi,
			Toko: productShop,
			Category: productCategory,
			Photos: productPhotos,
		}

		shop := response.TokoResponse{
			ID: detailTRX.Toko.ID,
			NamaToko: detailTRX.Toko.NamaToko,
			URLFoto: detailTRX.Toko.URLFoto,
		}

		dataDetailTRX := response.DetailTrx{
			Product: product,
			Toko: shop,
			Kuantitas: detailTRX.Kuantitas,
			HargaTotal: detailTRX.HargaTotal,
		}

		listDetailTRX = append(listDetailTRX, dataDetailTRX)
	}


	responseTRX := response.TRXResponse{
		ID: trx.ID,
		HargaTotal: trx.HargaTotal,
		KodeInvoice: trx.KodeInvoice,
		MethodBayar: trx.MethodBayar,
		ShippingAddress: trxAlamat,
		DetailTrx: listDetailTRX,
	}


	return responseTRX, nil
}