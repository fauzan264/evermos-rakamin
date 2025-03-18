package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
)

type provinceCityRepository struct {
	apiURL string
}

type ProvinceCityRepository interface {
	GetListProvince() ([]response.ProvinceResponse, error)
	GetDetailProvince(provinceID string) (response.ProvinceResponse, error)
	GetListCity(provinceID string) ([]response.CityResponse, error)
	GetDetailCity(cityID string) (response.CityResponse, error)
}

func NewProvinceCityRepository(apiURL string) *provinceCityRepository {
	return &provinceCityRepository{apiURL}
}

func (r *provinceCityRepository) GetListProvince() ([]response.ProvinceResponse, error) {
	resp, err := http.Get(r.apiURL + "/provinces.json")
	if err != nil {
		return []response.ProvinceResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []response.ProvinceResponse{}, constants.ErrFailedFetchListProvince
	}

	var provinces []response.ProvinceResponse
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return provinces, err
	}

	return provinces, nil
}

func (r *provinceCityRepository) GetDetailProvince(provinceID string) (response.ProvinceResponse, error) {
	url := fmt.Sprintf("%s/province/%s.json", r.apiURL, provinceID)
	resp, err := http.Get(url)
	if err != nil {
		return response.ProvinceResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return response.ProvinceResponse{}, constants.ErrFailedFetchDetailProvince
	}

	var province response.ProvinceResponse
	if err := json.NewDecoder(resp.Body).Decode(&province); err != nil {
		return province, err
	}

	return province, nil
}

func (r *provinceCityRepository) GetListCity(provinceID string) ([]response.CityResponse, error) {
	url := fmt.Sprintf("%s/regencies/%s.json", r.apiURL, provinceID)
	resp, err := http.Get(url)
	if err != nil {
		return []response.CityResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []response.CityResponse{}, constants.ErrFailedFetchListCity
	}

	var cities []response.CityResponse
	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return cities, err
	}

	return cities, nil
}

func (r *provinceCityRepository) GetDetailCity(cityID string) (response.CityResponse, error) {
	url := fmt.Sprintf("%s/regency/%s.json", r.apiURL, cityID)
	resp, err := http.Get(url)
	if err != nil {
		return response.CityResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return response.CityResponse{}, constants.ErrFailedFetchDetailCity
	}

	var city response.CityResponse
	if err := json.NewDecoder(resp.Body).Decode(&city); err != nil {
		return city, err
	}

	return city, nil
}
