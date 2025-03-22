package request

type TokoListRequest struct {
	Page	int					`query:"page"`
	Limit	int					`query:"limit"`
	Name	string				`query:"nama"`
}

type UpdateProfileShopRequest struct {
	Nama			string		`form:"nama" validate:"required"`
	Photo			string		`form:"photo" validate:"required"`
}