package request

type ShippingAddressRequest struct {
	JudulAlamat		string 	`json:"judul_alamat" validate:"required"`
	NamaPenerima	string	`json:"nama_penerima" validate:"required"`
	NoTelp			string	`json:"no_telp" validate:"required"`
	DetailAlamat	string	`json:"detail_alamat" validate:"required"`
}

type UpdateShippingAddressRequest struct {
	NamaPenerima	string	`json:"nama_penerima" validate:"required"`
	NoTelp			string	`json:"no_telp" validate:"required"`
	DetailAlamat	string	`json:"detail_alamat" validate:"required"`
}