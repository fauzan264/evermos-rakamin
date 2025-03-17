package response

type LoginResponse struct {
	Nama				string
	NoTelp				string
	TanggalLahir		string
	Tentang				string
	Pekerjaan			string
	Email				string
	ProvinsiResponse	ProvinsiResponse
	KotaResponse		KotaResponse
	Token				string
}

type ProvinsiResponse struct {
	ID					string
	Nama				string
}

type KotaResponse struct {
	ID					string
	IDProvinsi			string
	Nama				string
}
