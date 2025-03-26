package request

type GetByUserIDRequest struct {
	ID		int		`params:"id" validate:"required"`
}