package response

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