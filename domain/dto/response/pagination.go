package response

type PaginatedResponse struct {
	Data       any 			`json:"data"`
	Page       int      	`json:"page"`
	Limit      int      	`json:"limit"`
}