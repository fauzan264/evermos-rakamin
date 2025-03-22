package response

type TokoResponse struct {
	ID       uint64 `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
}