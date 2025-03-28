package response

import "github.com/fauzan264/evermos-rakamin/domain/model"

type TRXResponse struct {
	ID          		int         				`json:"id"`
	HargaTotal  		int         				`json:"harga_total"`
	KodeInvoice 		string      				`json:"kode_invoice"`
	MethodBayar 		string      				`json:"method_bayar"`
	ShippingAddress 	AddressResponse   			`json:"alamat_kirim"`
	DetailTrx   		[]DetailTrx 				`json:"detail_trx"`
}

type DetailTrx struct {
	Product    			LogProductResponse 			`json:"product"`
	Toko       			TokoResponse    			`json:"toko"`
	Kuantitas  			int     					`json:"kuantitas"`
	HargaTotal 			int     					`json:"harga_total"`
}

func TRXResponseFormatter(trx model.TRX) TRXResponse {
	trxAlamat := AddressResponse{
		ID: trx.Alamat.ID,
		JudulAlamat: trx.Alamat.JudulAlamat,
		NamaPenerima: trx.Alamat.NamaPenerima,
		NoTelp: trx.Alamat.NoTelp,
		DetailAlamat: trx.Alamat.DetailAlamat,
	}

	var listDetailTRX []DetailTrx
	for _, detailTRX := range trx.DetailTRX {

		productShop := TokoResponse{
			ID: detailTRX.LogProduct.Toko.ID,
			NamaToko: detailTRX.LogProduct.Toko.NamaToko,
			URLFoto: detailTRX.LogProduct.Toko.URLFoto,
		}

		productCategory := CategoryResponse{
			ID: detailTRX.LogProduct.Category.ID,
			NamaCategory: detailTRX.LogProduct.Category.NamaCategory,
			CreatedAt: &detailTRX.LogProduct.Category.CreatedAt,
			UpdatedAt: &detailTRX.LogProduct.Category.UpdatedAt,
		}

		var productPhotos []PhotoProductResponse
		for _, photo := range detailTRX.LogProduct.Produk.PhotosProduct {
			productPhoto := PhotoProductResponse{
				ID: photo.ID,
				IDProduk: photo.IDProduk,
				URL: photo.URL,
			}

			productPhotos = append(productPhotos, productPhoto)
		}

		product := LogProductResponse{
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

		shop := TokoResponse{
			ID: detailTRX.Toko.ID,
			NamaToko: detailTRX.Toko.NamaToko,
			URLFoto: detailTRX.Toko.URLFoto,
		}

		dataDetailTRX := DetailTrx{
			Product: product,
			Toko: shop,
			Kuantitas: detailTRX.Kuantitas,
			HargaTotal: detailTRX.HargaTotal,
		}

		listDetailTRX = append(listDetailTRX, dataDetailTRX)
	}

	trxResponse := TRXResponse{
		ID: trx.ID,
		HargaTotal: trx.HargaTotal,
		KodeInvoice: trx.KodeInvoice,
		MethodBayar: trx.MethodBayar,
		ShippingAddress: trxAlamat,
		DetailTrx: listDetailTRX,
	}
	return trxResponse
}

func ListTRXResponseFormatter(listTRX []model.TRX) []TRXResponse {
	var listTRXResponse []TRXResponse
	for _, trx := range listTRX {
		trxResponseFormatter := TRXResponseFormatter(trx)
		listTRXResponse = append(listTRXResponse, trxResponseFormatter)
	}

	return listTRXResponse
}