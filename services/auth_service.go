package services

import (
	"time"

	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"github.com/fauzan264/evermos-rakamin/repositories"
	"github.com/fauzan264/evermos-rakamin/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(request request.RegisterRequest) error
	LoginUser(request request.LoginRequest) (response.LoginResponse, error)
}

type authService struct {
	repository repositories.UserRepository
	shopRepository repositories.ShopRepository
	provinceCityRepository repositories.ProvinceCityRepository
}

func NewAuthService(
	repository repositories.UserRepository,
	shopRepository repositories.ShopRepository,
	provinceCityRepository repositories.ProvinceCityRepository,
) *authService {
	return &authService{
		repository,
		shopRepository,
		provinceCityRepository,
	}
}

func (s *authService) RegisterUser(request request.RegisterRequest) (error) {
	tanggalLahir, err := time.Parse("02/01/2006", request.TanggalLahir)
	if err != nil {
		return constants.ErrInvalidDateFormat
	}
	
	user := model.User{
		Nama: request.Nama,
		NoTelp: request.NoTelp,
		TanggalLahir: tanggalLahir,
		Pekerjaan: request.Pekerjaan,
		Email: request.Email,
		IDProvinsi: request.IDProvinsi,
		IDKota: request.IDKota,
	}
	
	kataSandiHash, err := bcrypt.GenerateFromPassword([]byte(request.KataSandi), bcrypt.MinCost)
	if err != nil {
		return err
	}
	
	user.KataSandi = string(kataSandiHash)
	newUser, err := s.repository.CreateUser(user)
	if err != nil {
		return err
	}
	
	shop := model.Shop{
		IDUser: newUser.ID,
	}
	
	err = s.shopRepository.CreateShop(shop)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) LoginUser(request request.LoginRequest) (response.LoginResponse, error) {
	noTelp := request.NoTelp
	kataSandi := request.KataSandi

	user, err := s.repository.GetUserByNoTelp(noTelp)
	if err != nil {
		return response.LoginResponse{}, constants.ErrWrongUserOrPassword
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(kataSandi))
	if err != nil {
		return response.LoginResponse{}, constants.ErrWrongUserOrPassword
	}

	dataProvince, err := s.provinceCityRepository.GetDetailProvince(user.IDProvinsi)
	if err != nil {
		return response.LoginResponse{}, err
	}

	dataCity, err := s.provinceCityRepository.GetDetailCity(user.IDKota)
	if err != nil {
		return response.LoginResponse{}, err
	}

	
	token, err := utils.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		return response.LoginResponse{}, err
	}

	loginResponse := response.LoginResponse{
		Nama: user.Nama,
		NoTelp: user.NoTelp,
		TanggalLahir: user.TanggalLahir,
		Tentang: user.Tentang,
		Pekerjaan: user.Pekerjaan,
		Email: user.Email,
		ProvinsiResponse: dataProvince,
		KotaResponse: dataCity,
		Token: token,
	}
	return loginResponse, err
}