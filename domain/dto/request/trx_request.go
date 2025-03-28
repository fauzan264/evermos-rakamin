package request

type GetByTRXIDRequest struct {
	ID				int		`params:"id" validate:"required"`	
}

type TRXListRequest struct {
	Page			int			`query:"page"`
	Limit			int			`query:"limit"`
	Search			string		`query:"search"`
}

type CreateTrxRequest struct {
	MethodBayar			string				`json:"method_bayar" binding:"required"`
	AlamatPengiriman	int					`json:"alamat_kirim" binding:"required"`
	DetailTrxRequest 	[]DetailTrxRequest	`json:"detail_trx" binding:"required"`
}

type DetailTrxRequest struct {
	IDProduk			int					`json:"product_id" binding:"required"`
	Kuantitas			int					`json:"kuantitas" binding:"required"`
}