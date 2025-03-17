package response

type ProvinceResponse struct {
	ID					string	`json:"id"`
	Name				string	`json:"name"`
}

type CityResponse struct {
	ID					string	`json:"id"`
	ProvinceID			string	`json:"province_id"`
	Name				string	`json:"name"`
}