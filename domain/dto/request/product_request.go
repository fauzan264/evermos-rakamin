package request

type GetByProductIDRequest struct {
	ID					int 					`params:"id" validate:"required"`	
}

type ProductRequest struct {
	NamaProduk			string					`json:"nama_produk" validate:"required"`
	IDCategory			*int						`json:"category_id" validate:"required"`
	HargaReseller		string					`json:"harga_reseller" validate:"required"`
	HargaKonsumen		string					`json:"harga_konsumen" validate:"required"`
	Stok				*int						`json:"stok" validate:"required"`
	Deskripsi			string					`json:"deskripsi" validate:"required"`
	Photos 				[]PhotoProductRequest 	`json:"photos" validate:"required,dive"`
}

type PhotoProductRequest struct {
	URL 				string 					`json:"photos"`
}