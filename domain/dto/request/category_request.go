package request

type GetByCategoryIDRequest struct {
	ID					int 		`params:"id" validate:"required"`
}

type CategoryRequest struct {
	NamaCategory		string		`json:"nama_category" validate:"required"`
}