package request

type CategoryRequest struct {
	NamaCategory		string		`json:"nama_category" validate:"required"`
}