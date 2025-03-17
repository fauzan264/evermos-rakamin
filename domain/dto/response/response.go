package response

type Response struct {
	Status		bool 		`json:"status"`
	Message 	string 		`json:"message"`
	Errors		interface{}	`json:"errors"`
	Data 		interface{}	`json:"data"`
}