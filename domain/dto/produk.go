package dto

type CreateProductRequest struct {
	NamaProduk			string				`form:"nama_produk" validate:"required"`
	IDCategory			string				`form:"category_id" validate:"required"`
	HargaReseller		string				`form:"harga_reseller" validate:"required"`
	HargaKonsumen		string				`form:"harga_konsumen" validate:"required"`
	Stok				int					`form:"stok" validate:"required"`
	Deskripsi			string				`form:"deskripsi" validate:"required"`
	Photo				string				`form:"photos" validate:"required"`
}

type UpdateProductRequest struct {
	NamaProduk			string				`form:"nama_produk"`
	IDCategory			string				`form:"category_id"`
	HargaReseller		string				`form:"harga_reseller"`
	HargaKonsumen		string				`form:"harga_konsumen"`
	Stok				int					`form:"stok"`
	Deskripsi			string				`form:"deskripsi"`
	Photo				string				`form:"photos"`
}