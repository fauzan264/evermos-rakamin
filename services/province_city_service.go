package services

import (
	"strconv"

	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/repositories"
)

type ProvinceCityService interface {
	GetListProvince() ([]response.ProvinceResponse, error)
	GetDetailProvince(request request.GetByProvinceIDRequest) (response.ProvinceResponse, error)
	GetListCity(request request.GetByProvinceIDRequest) ([]response.CityResponse, error)
	GetDetailCity(request request.GetByCityIDRequest) (response.CityResponse, error)
}

type provinceCityService struct {
	repository repositories.ProvinceCityRepository
}

func NewProvinceCityService(repository repositories.ProvinceCityRepository) *provinceCityService {
	return &provinceCityService{repository}
}


func (s *provinceCityService) GetListProvince() ([]response.ProvinceResponse, error) {
	getListProvince, err := s.repository.GetListProvince()
	if err != nil {
		return getListProvince, err
	}

	return getListProvince, nil
}

func (s *provinceCityService) GetDetailProvince(request request.GetByProvinceIDRequest) (response.ProvinceResponse, error) {
	provinceID := strconv.Itoa(request.ProvinceID)
	getProvince, err := s.repository.GetDetailProvince(provinceID)
	if err != nil {
		return getProvince, err
	}

	return getProvince, nil
}

func (s *provinceCityService) GetListCity(request request.GetByProvinceIDRequest) ([]response.CityResponse, error) {
	provinceID := strconv.Itoa(request.ProvinceID)
	getListCity, err := s.repository.GetListCity(provinceID)
	if err != nil {
		return getListCity, err
	}

	return getListCity, nil
}

func (s *provinceCityService) GetDetailCity(request request.GetByCityIDRequest) (response.CityResponse, error) {
	cityID := strconv.Itoa(request.CityID)
	getCity, err := s.repository.GetDetailCity(cityID)
	if err != nil {
		return getCity, err
	}

	return getCity, nil
}
