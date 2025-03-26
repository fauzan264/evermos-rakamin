package request

type GetByUserIDRequest struct {
	ID				int		`params:"id" validate:"required"`
}

type UpdateProfileRequest struct {
	Nama			string 		`json:"nama" validate:"required"`
	KataSandi		string 		`json:"kata_sandi" validate:"required"`	
	NoTelp			string 		`json:"no_telp" validate:"required"`
	TanggalLahir	string 		`json:"tanggal_lahir" validate:"required"`		
	Pekerjaan		string 		`json:"pekerjaan" validate:"required"`	
	Email			string 		`json:"email" validate:"required"`
	IDProvinsi		string 		`json:"id_provinsi" validate:"required"`	
	IDKota			string 		`json:"id_kota" validate:"required"`
}