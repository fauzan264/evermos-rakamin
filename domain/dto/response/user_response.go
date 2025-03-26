package response

import "time"

type UserResponse struct {
	ID 				int				`json:"id"`
	Nama 			string			`json:"nama"`
	NoTelp 			string			`json:"no_telp"`
	TanggalLahir 	time.Time		`json:"tanggal_lahir"`
	JenisKelamin 	string			`json:"jenis_kelamin"`
	Tentang 		string			`json:"tentang"`
	Pekerjaan 		string			`json:"pekerjaan"`
	Email 			string			`json:"email"`
	IDProvinsi 		string			`json:"id_provinsi"`
	IDKota 			string			`json:"id_kota"`
	IsAdmin 		bool			`json:"is_admin"`
}