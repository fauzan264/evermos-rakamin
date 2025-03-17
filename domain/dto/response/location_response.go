package response

type ProvinsiResponse struct {
	ID					string
	Nama				string
}

type KotaResponse struct {
	ID					string
	IDProvinsi			string
	Nama				string
}
