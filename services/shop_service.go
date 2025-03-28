package services

import (
	"os"

	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/repositories"
)

type TokoService interface {
	GetMyToko(requestUser request.GetByUserIDRequest) (response.TokoResponse, error)
	GetListToko(request request.TokoListRequest) (response.PaginatedResponse, error)
	GetTokoByID(requestID request.GetTokoByID) (response.TokoResponse, error)
	UpdateToko(requestUser request.GetByUserIDRequest,  requestID request.GetTokoByID, requestData request.UpdateProfileShopRequest) (response.TokoResponse, error)
}

type tokoService struct {
	repository repositories.TokoRepository
}

func NewTokoService(repository repositories.TokoRepository) *tokoService {
	return &tokoService{repository}
}

func (s *tokoService) GetMyToko(requestUser request.GetByUserIDRequest) (response.TokoResponse, error) {
	shop, err := s.repository.GetTokoByUserID(requestUser.ID)
	if err != nil {
		return response.TokoResponse{}, err
	}

	shopResponse := response.ShopResponseFormatter(shop)

	return shopResponse, nil
}

func (s *tokoService) GetListToko(request request.TokoListRequest) (response.PaginatedResponse, error) {
	page := request.Page
	limit := request.Limit
	name := request.Name

	listShop, err := s.repository.GetListToko(page, limit, name)

	var shopResponse response.PaginatedResponse

	if err != nil {
		return shopResponse, err
	}

	listShopFormatter := response.ListShopResponseFormatter(listShop)

	if len(listShop) == 0 {
		listShopFormatter = []response.TokoResponse{}
	}

	shopResponse.Data = listShopFormatter
	shopResponse.Page = request.Page
	shopResponse.Limit = request.Limit

	return shopResponse, nil
}

func (s *tokoService) GetTokoByID(requestID request.GetTokoByID) (response.TokoResponse, error) {
	shop, err := s.repository.GetTokoByID(requestID.ID)
	if err != nil {
		return response.TokoResponse{}, err
	}

	shopResponse := response.ShopResponseFormatter(shop)
	
	return shopResponse, nil
}

func (s *tokoService) UpdateToko(
	requestUser request.GetByUserIDRequest,
	requestID request.GetTokoByID,
	requestData request.UpdateProfileShopRequest,
) (response.TokoResponse, error) {
	toko, err := s.repository.GetTokoByID(requestID.ID)
	if err != nil {
		return response.TokoResponse{}, err
	}

	if toko.IDUser != requestUser.ID {
		return response.TokoResponse{}, constants.ErrUnauthorized
	}

	toko.NamaToko = requestData.Nama

	if requestData.Photo != "" {
		if toko.URLFoto != "" {
			os.Remove(toko.URLFoto)
		}
		toko.URLFoto = requestData.Photo
	}

	updateShop, err := s.repository.UpdateToko(toko)
	if err != nil {
		return response.TokoResponse{}, err
	}

	shopResponse := response.ShopResponseFormatter(updateShop)

	return shopResponse, nil
}