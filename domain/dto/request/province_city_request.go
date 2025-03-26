package request

type GetByProvinceIDRequest struct {
	ProvinceID		int 	`params:"prov_id" validate:"required"`
}

type GetByCityIDRequest struct {
	CityID 			int 	`params:"city_id" validate:"required"`
}