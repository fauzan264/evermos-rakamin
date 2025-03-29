package response

import "github.com/fauzan264/evermos-rakamin/domain/model"

type ProductResponse struct {
	ID				int						`json:"id"`
	NamaProduk		string					`json:"nama_produk"`
	Slug			string					`json:"slug"`
	HargaReseller	string					`json:"harga_reseller"`
	HargaKonsumen	string					`json:"harga_konsumen"`
	Stok			int						`json:"stok"`
	Deskripsi		string					`json:"deskripsi"`
	Toko			ShopResponse			`json:"toko"`
	Category		CategoryResponse		`json:"category"`
	Photos			[]PhotoProductResponse	`json:"photos"`
}

type PhotoProductResponse struct {
	ID 				int						`json:"id"`
	IDProduk 		int						`json:"product_id"`
	URL 			string					`json:"url"`
}

type LogProductResponse struct {
	ID				int						`json:"id"`
	NamaProduk		string					`json:"nama_produk"`
	Slug			string					`json:"slug"`
	HargaReseller	string					`json:"harga_reseller"`
	HargaKonsumen	string					`json:"harga_konsumen"`
	Deskripsi		string					`json:"deskripsi"`
	Toko			ShopResponse			`json:"toko"`
	Category		CategoryResponse		`json:"category"`
	Photos			[]PhotoProductResponse	`json:"photos"`
}

func ProductResponseFormatter(product model.Product) ProductResponse {
	productShop := ShopResponseFormatter(product.Toko)
	
	productCategory := CategoryResponseFormatter(product.Category)
	
	var productPhotos []PhotoProductResponse
	for _, photo := range product.PhotosProduct {
		photoProduct := PhotoProductResponse{
			ID: photo.ID,
			IDProduk: photo.IDProduk,
			URL: photo.URL,
		}

		productPhotos = append(productPhotos, photoProduct)
	}

	productResponse := ProductResponse{
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

	return productResponse
}

func ListProductResponseFormatter(listProduct []model.Product) []ProductResponse {
	var listProductResponse []ProductResponse
	for _, product := range listProduct {
		productResponseFormatter := ProductResponseFormatter(product)
		listProductResponse = append(listProductResponse, productResponseFormatter)
	}

	return listProductResponse
}
