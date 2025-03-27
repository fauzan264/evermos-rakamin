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
	getMyToko, err := s.repository.GetTokoByUserID(requestUser.ID)
	if err != nil {
		return response.TokoResponse{}, err
	}

	tokoResponse := response.TokoResponse{
		ID : getMyToko.ID,
		NamaToko: getMyToko.NamaToko,
		URLFoto: getMyToko.URLFoto,
	}

	return tokoResponse, nil
}

func (s *tokoService) GetListToko(request request.TokoListRequest) (response.PaginatedResponse, error) {
	page := request.Page
	limit := request.Limit
	name := request.Name

	getListToko, err := s.repository.GetListToko(page, limit, name)

	var tokoResponse response.PaginatedResponse

	if err != nil {
		return tokoResponse, err
	}

	listToko := make([]response.TokoResponse, 0)
	for _, toko := range getListToko {
		dataToko := response.TokoResponse{
			ID: toko.ID,
			NamaToko: toko.NamaToko,
			URLFoto: toko.URLFoto,
		}

		listToko = append(listToko, dataToko)
	}

	tokoResponse.Data = listToko
	tokoResponse.Page = request.Page
	tokoResponse.Limit = request.Limit

	return tokoResponse, nil
}

func (s *tokoService) GetTokoByID(requestID request.GetTokoByID) (response.TokoResponse, error) {
	toko, err := s.repository.GetTokoByID(requestID.ID)
	if err != nil {
		return response.TokoResponse{}, err
	}

	tokoResponse := response.TokoResponse{
		ID : toko.ID,
		NamaToko : toko.NamaToko,
		URLFoto : toko.URLFoto,
	}
	return tokoResponse, nil
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

	updateToko, err := s.repository.UpdateToko(toko)
	if err != nil {
		return response.TokoResponse{}, err
	}

	responseToko := response.TokoResponse{
		ID: updateToko.ID,
		NamaToko: updateToko.NamaToko,
		URLFoto: updateToko.URLFoto,
	}

	return responseToko, nil
}