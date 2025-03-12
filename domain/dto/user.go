package dto

import "time"

type LoginRequest struct {
	NoTelp			string	`json:"no_telp" validate:"required"`
	KataSandi		string	`json:"kata_sandi" validate:"required"`
}

type RegisterRequest struct {
	Nama			string		`json:"nama" validate:"required"`
	KataSandi		string		`json:"kata_sandi" validate:"required"`
	NoTelp			string		`json:"no_telp" validate:"required"`
	TanggalLahir	time.Time	`json:"tanggal_lahir" validate:"required"`
	Pekerjaan		string		`json:"pekerjaan" validate:"required"`
	Email			string		`json:"email" validate:"required"`
	IDProvinsi		string		`json:"id_provinsi" validate:"required"`
	IDKota			string		`json:"id_kota" validate:"required"`
}

type UpdateProfileRequest struct {
	Nama			string 		`json:"nama" validate:"required"`
	KataSandi		string 		`json:"kata_sandi" validate:"required"`	
	NoTelp			string 		`json:"no_telp" validate:"required"`
	TanggalLahir	string 		`json:"tanggal_Lahir" validate:"required"`		
	Pekerjaan		string 		`json:"pekerjaan" validate:"required"`	
	Email			string 		`json:"email" validate:"required"`
	IDProvinsi		string 		`json:"id_provinsi" validate:"required"`	
	IDKota			string 		`json:"id_kota" validate:"required"`
}