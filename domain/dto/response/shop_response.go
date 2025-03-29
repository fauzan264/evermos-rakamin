package response

import (
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"github.com/fauzan264/evermos-rakamin/helpers"
)

type TokoResponse struct {
	ID       	int 		`json:"id"`
	NamaToko 	string 		`json:"nama_toko"`
	URLFoto  	string 		`json:"url_foto"`
}


func ShopResponseFormatter(shop model.Toko) TokoResponse {
	shopResponse := TokoResponse{
		ID: shop.ID,
		NamaToko: shop.NamaToko,
		URLFoto: helpers.GetImageURL(shop.URLFoto),
	}

	return shopResponse
}

func ListShopResponseFormatter(listShop []model.Toko) []TokoResponse {
	var listShopResponse []TokoResponse
	for _, shop := range listShop {
		shopResponseFormatter := ShopResponseFormatter(shop)
		listShopResponse = append(listShopResponse, shopResponseFormatter)
	}

	return listShopResponse
}