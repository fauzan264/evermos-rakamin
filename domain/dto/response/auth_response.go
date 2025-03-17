package response

import "time"

type LoginResponse struct {
	Nama				string				`json:"nama"`
	NoTelp				string				`json:"no_telp"`
	TanggalLahir		time.Time			`json:"tanggal_lahir"`
	Tentang				string				`json:"tentang"`
	Pekerjaan			string				`json:"pekerjaan"`
	Email				string				`json:"email"`
	ProvinsiResponse	ProvinceResponse	`json:"id_provinsi"`
	KotaResponse		CityResponse		`json:"id_kota"`
	Token				string				`json:"token"`
}