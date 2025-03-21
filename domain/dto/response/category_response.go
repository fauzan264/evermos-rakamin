package response

import "time"

type CategoryResponse struct {
	ID				int			`json:"id"`
	NamaCategory	string 		`json:"nama_category"`
	CreatedAt		*time.Time	`json:"created_at,omitempty"`
	UpdatedAt		*time.Time	`json:"updated_at,omitempty"`
}