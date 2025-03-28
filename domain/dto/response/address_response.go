package response

import "github.com/fauzan264/evermos-rakamin/domain/model"

type AddressResponse struct {
	ID 				int				`json:"id"`
	JudulAlamat 	string			`json:"judul_alamat"`
	NamaPenerima 	string			`json:"nama_penerima"`
	NoTelp 			string			`json:"no_telp"`
	DetailAlamat 	string			`json:"detail_alamat"`
}



func AddressResponseFormatter(address model.Alamat) AddressResponse {
	addressResponse := AddressResponse{
		ID : address.ID,
		JudulAlamat : address.JudulAlamat,
		NamaPenerima : address.NamaPenerima,
		NoTelp : address.NoTelp,
		DetailAlamat : address.DetailAlamat,
	}

	return addressResponse
}

func ListAddressResponseFormatter(listAddress []model.Alamat) []AddressResponse {
	var listAddressResponse []AddressResponse
	for _, address := range listAddress {
		addressResponseFormatter := AddressResponseFormatter(address)
		listAddressResponse = append(listAddressResponse, addressResponseFormatter)
	}

	return listAddressResponse
}