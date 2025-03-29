package response

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