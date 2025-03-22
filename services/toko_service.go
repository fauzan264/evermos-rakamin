package services

import (
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/repositories"
)

type TokoService interface {
	GetListToko(request request.TokoListRequest) (response.PaginatedResponse, error)
}

type tokoService struct {
	tokoRepository repositories.TokoRepository
}

func NewTokoService(tokoRepository repositories.TokoRepository) *tokoService {
	return &tokoService{tokoRepository}
}

func (s *tokoService) GetListToko(request request.TokoListRequest) (response.PaginatedResponse, error) {
	page := request.Page
	limit := request.Limit
	name := request.Name

	getListToko, err := s.tokoRepository.GetListToko(page, limit, name)

	var tokoResponse response.PaginatedResponse

	if err != nil {
		return tokoResponse, err
	}

	listToko := make([]response.TokoResponse, 0)
	for _, toko := range getListToko {
		dataToko := response.TokoResponse{
			ID: uint64(toko.ID),
			NamaToko: toko.NamaToko,
			UrlFoto: toko.URLFoto,
		}

		listToko = append(listToko, dataToko)
	}

	tokoResponse.Data = listToko
	tokoResponse.Page = request.Page
	tokoResponse.Limit = request.Limit

	return tokoResponse, nil
}
