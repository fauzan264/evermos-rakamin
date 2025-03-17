package request

type UpdateProfileShopRequest struct {
	Nama			string		`form:"nama" validate:"required"`
	Photo			string		`form:"photo" validate:"required"`
}