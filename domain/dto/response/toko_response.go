package response

type TokoResponse struct {
	ID       	int 		`json:"id"`
	NamaToko 	string 		`json:"nama_toko"`
	URLFoto  	string 		`json:"url_foto"`
}