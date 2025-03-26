package services

import (
	"time"

	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/repositories"
)

type userService struct {
	repository repositories.UserRepository
}

type UserService interface {
	GetUserByID(request request.GetByUserIDRequest) (response.UserResponse, error)
	UpdateUser(requestID request.GetByUserIDRequest, requestData request.UpdateProfileRequest) (response.UserResponse, error)
}
func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) GetUserByID(request request.GetByUserIDRequest) (response.UserResponse, error) {
	getUser, err := s.repository.GetUserByID(request.ID)
	if err != nil {
		return response.UserResponse{}, err
	}

	userResponse := response.UserResponse{
		ID: getUser.ID,
		Nama: getUser.Nama,
		NoTelp: getUser.NoTelp,
		TanggalLahir: getUser.TanggalLahir,
		JenisKelamin: getUser.JenisKelamin,
		Tentang: getUser.Tentang,
		Pekerjaan: getUser.Pekerjaan,
		Email: getUser.Email,
		IDProvinsi: getUser.IDProvinsi,
		IDKota: getUser.IDKota,
		IsAdmin: getUser.IsAdmin,
	}

	return userResponse, nil
}

func (s *userService) UpdateUser(requestID request.GetByUserIDRequest, requestData request.UpdateProfileRequest) (response.UserResponse, error) {
	user, err := s.repository.GetUserByID(requestID.ID)
	if err != nil {
		return response.UserResponse{}, err
	}

	tanggalLahir, err := time.Parse("02/01/2006", requestData.TanggalLahir)
	if err != nil {
		return response.UserResponse{}, constants.ErrInvalidDateFormat
	}

	user.Nama = requestData.Nama
	user.KataSandi = requestData.KataSandi
	user.NoTelp = requestData.NoTelp
	user.TanggalLahir = tanggalLahir
	user.Pekerjaan = requestData.Pekerjaan
	user.Email = requestData.Email
	user.IDProvinsi = requestData.IDProvinsi
	user.IDKota = requestData.IDKota
	
	updateUser, err := s.repository.UpdateUser(user)
	if err != nil {
		return response.UserResponse{}, err
	}

	userResponse := response.UserResponse{
		ID: updateUser.ID,
		Nama: updateUser.Nama,
		NoTelp: updateUser.NoTelp,
		TanggalLahir: updateUser.TanggalLahir,
		Pekerjaan: updateUser.Pekerjaan,
		Email: updateUser.Email,
		IDProvinsi: updateUser.IDProvinsi,
		IDKota: updateUser.IDKota,
	}

	return userResponse, nil
}