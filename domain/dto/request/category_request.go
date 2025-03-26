package request

import "github.com/fauzan264/evermos-rakamin/domain/model"

type GetByCategoryIDRequest struct {
	ID					int 			`params:"id" validate:"required"`
}

type CategoryRequest struct {
	NamaCategory		string			`json:"nama_category" validate:"required"`
	User				model.User
}