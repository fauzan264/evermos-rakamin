package services

import (
	"os"

	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/repositories"
)

type ShopService interface {
	GetMyShop(requestUser request.GetByUserIDRequest) (response.ShopResponse, error)
	GetListShop(request request.ShopListRequest) (response.PaginatedResponse, error)
	GetShopByID(requestID request.GetShopByID) (response.ShopResponse, error)
	UpdateShop(requestUser request.GetByUserIDRequest,  requestID request.GetShopByID, requestData request.UpdateProfileShopRequest) (response.ShopResponse, error)
}

type shopService struct {
	repository repositories.ShopRepository
}

func NewShopService(repository repositories.ShopRepository) *shopService {
	return &shopService{repository}
}

func (s *shopService) GetMyShop(requestUser request.GetByUserIDRequest) (response.ShopResponse, error) {
	shop, err := s.repository.GetShopByUserID(requestUser.ID)
	if err != nil {
		return response.ShopResponse{}, err
	}

	shopResponse := response.ShopResponseFormatter(shop)

	return shopResponse, nil
}

func (s *shopService) GetListShop(requestSearch request.ShopListRequest) (response.PaginatedResponse, error) {
	var shopResponse response.PaginatedResponse
	
	listShop, err := s.repository.GetListShop(requestSearch)

	if err != nil {
		return shopResponse, err
	}

	listShopFormatter := response.ListShopResponseFormatter(listShop)

	if len(listShop) == 0 {
		listShopFormatter = []response.ShopResponse{}
	}

	shopResponse.Data = listShopFormatter
	shopResponse.Page = requestSearch.Page
	shopResponse.Limit = requestSearch.Limit

	return shopResponse, nil
}

func (s *shopService) GetShopByID(requestID request.GetShopByID) (response.ShopResponse, error) {
	shop, err := s.repository.GetShopByID(requestID.ID)
	if err != nil {
		return response.ShopResponse{}, err
	}

	shopResponse := response.ShopResponseFormatter(shop)
	
	return shopResponse, nil
}

func (s *shopService) UpdateShop(
	requestUser request.GetByUserIDRequest,
	requestID request.GetShopByID,
	requestData request.UpdateProfileShopRequest,
) (response.ShopResponse, error) {
	shop, err := s.repository.GetShopByID(requestID.ID)
	if err != nil {
		return response.ShopResponse{}, err
	}

	if shop.IDUser != requestUser.ID {
		return response.ShopResponse{}, constants.ErrUnauthorized
	}

	shop.NamaToko = requestData.Nama

	if requestData.Photo != "" {
		if shop.URLFoto != "" {
			os.Remove(shop.URLFoto)
		}
		shop.URLFoto = requestData.Photo
	}

	updateShop, err := s.repository.UpdateShop(shop)
	if err != nil {
		return response.ShopResponse{}, err
	}

	shopResponse := response.ShopResponseFormatter(updateShop)

	return shopResponse, nil
}