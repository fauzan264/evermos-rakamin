package request

type GetByAddressIDRequest struct {
	ID				int			`params:"id" validate:"required"`
}

type AddressListRequest struct {
	Title			string			`query:"judul_alamat"`
	Page			int				`query:"page"`
	Limit			int				`query:"limit"`
}

type CreateAddressRequest struct {
	JudulAlamat		string 		`json:"judul_alamat" validate:"required"`
	NamaPenerima	string		`json:"nama_penerima" validate:"required"`
	NoTelp			string		`json:"no_telp" validate:"required"`
	DetailAlamat	string		`json:"detail_alamat" validate:"required"`
}

type UpdateAddressRequest struct {
	NamaPenerima	string		`json:"nama_penerima" validate:"required"`
	NoTelp			string		`json:"no_telp" validate:"required"`
	DetailAlamat	string		`json:"detail_alamat" validate:"required"`
}