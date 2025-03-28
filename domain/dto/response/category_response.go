package response

import (
	"time"

	"github.com/fauzan264/evermos-rakamin/domain/model"
)

type CategoryResponse struct {
	ID				int			`json:"id"`
	NamaCategory	string 		`json:"nama_category"`
	CreatedAt		*time.Time	`json:"created_at,omitempty"`
	UpdatedAt		*time.Time	`json:"updated_at,omitempty"`
}

func CategoryResponseFormatter(category model.Category) CategoryResponse {
	categoryResponse := CategoryResponse{
		ID: category.ID,
		NamaCategory: category.NamaCategory,
	}

	return categoryResponse
}

func ListCategoryResponseFormatter(listCategory []model.Category) []CategoryResponse {
	var listCategoryResponse []CategoryResponse
	for _, category := range listCategory {
		categoryResponseFormatter := CategoryResponseFormatter(category)
		listCategoryResponse = append(listCategoryResponse, categoryResponseFormatter)
	}

	return listCategoryResponse
}

