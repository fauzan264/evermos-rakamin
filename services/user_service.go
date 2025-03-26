package services

import (
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/repositories"
)

type userService struct {
	userRepository repositories.UserRepository
}

type UserService interface {
	GetUserByID(request request.GetByUserIDRequest) (response.UserResponse, error)
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) GetUserByID(request request.GetByUserIDRequest) (response.UserResponse, error) {
	getUser, err := s.userRepository.GetUserByID(request.ID)
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