package services

import (
	"time"

	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"github.com/fauzan264/evermos-rakamin/repositories"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repository repositories.UserRepository
	addressRepository repositories.AddressRepository
}

type UserService interface {
	GetUserByID(request request.GetByUserIDRequest) (response.UserResponse, error)
	UpdateUser(requestID request.GetByUserIDRequest, requestData request.UpdateProfileRequest) (response.UserResponse, error)
	GetMyAddress(requestUser request.GetByUserIDRequest, requestSearch request.AddressListRequest) (response.PaginatedResponse, error)
	GetAddressUserByID(requestUser request.GetByUserIDRequest, requestID request.GetByAddressIDRequest) (response.AddressResponse, error)
	CreateAddressUser(requestUser request.GetByUserIDRequest, requestData request.CreateAddressRequest) (response.AddressResponse, error)
	UpdateAddressUser(requestUser request.GetByUserIDRequest,  requestID request.GetByAddressIDRequest, requestData request.UpdateAddressRequest) (response.AddressResponse, error)
	DeleteAddressUser(requestUser request.GetByUserIDRequest,  requestID request.GetByAddressIDRequest) error

}
func NewUserService(repository repositories.UserRepository, addressRepository repositories.AddressRepository) *userService {
	return &userService{repository, addressRepository}
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
	user.NoTelp = requestData.NoTelp
	user.TanggalLahir = tanggalLahir
	user.Pekerjaan = requestData.Pekerjaan
	user.Email = requestData.Email
	user.IDProvinsi = requestData.IDProvinsi
	user.IDKota = requestData.IDKota

	kataSandiHash, err := bcrypt.GenerateFromPassword([]byte(requestData.KataSandi), bcrypt.MinCost)
	if err != nil {
		return response.UserResponse{}, err
	}

	user.KataSandi = string(kataSandiHash)
	
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

func (s *userService) GetMyAddress(requestUser request.GetByUserIDRequest, requestSearch request.AddressListRequest) (response.PaginatedResponse, error) {
	var myAddressResponse response.PaginatedResponse

	listMyAddress, err := s.addressRepository.GetAddressByUserID(requestUser.ID, requestSearch)
	if err != nil {
		return response.PaginatedResponse{}, err
	}

	listMyAddressResponse := response.ListAddressResponseFormatter(listMyAddress)

	if len(listMyAddressResponse) == 0 {
		listMyAddressResponse = []response.AddressResponse{}
	}

	myAddressResponse.Data = listMyAddressResponse
	myAddressResponse.Page = requestSearch.Page
	myAddressResponse.Limit = requestSearch.Limit

	return myAddressResponse, nil
}

func (s *userService) GetAddressUserByID(requestUser request.GetByUserIDRequest, requestID request.GetByAddressIDRequest) (response.AddressResponse, error) {
	address, err := s.addressRepository.GetAddressUserByID(requestUser.ID, requestID.ID)
	if err != nil {
		return response.AddressResponse{}, err
	}

	addressResponse := response.AddressResponseFormatter(address)

	return addressResponse, nil
}

func (s *userService) CreateAddressUser(requestUser request.GetByUserIDRequest, requestData request.CreateAddressRequest) (response.AddressResponse, error) {
	address := model.Address{
		IDUser: requestUser.ID,
		JudulAlamat: requestData.JudulAlamat,
		NamaPenerima: requestData.NamaPenerima,
		NoTelp: requestData.NoTelp,
		DetailAlamat: requestData.DetailAlamat,
		CreatedAt: time.Now(),
	}

	createAddress, err := s.addressRepository.CreateAddress(address)
	if err != nil {
		return response.AddressResponse{}, err
	}

	addressResponse := response.AddressResponseFormatter(createAddress)

	return addressResponse, nil
}

func (s *userService) UpdateAddressUser(requestUser request.GetByUserIDRequest,  requestID request.GetByAddressIDRequest, requestData request.UpdateAddressRequest) (response.AddressResponse, error) {
	getAddress, err := s.addressRepository.GetAddressByID(requestID.ID)
	if err != nil {
		return response.AddressResponse{}, err
	}

	if getAddress.IDUser != requestUser.ID {
		return response.AddressResponse{}, constants.ErrUnauthorized
	}

	address := model.Address{
		ID : requestID.ID,
		IDUser : getAddress.IDUser,
		JudulAlamat : getAddress.JudulAlamat,
		NamaPenerima : requestData.NamaPenerima,
		NoTelp : requestData.NoTelp,
		DetailAlamat : requestData.DetailAlamat,
		CreatedAt : getAddress.CreatedAt,
		UpdatedAt : time.Now(),
	}

	updateAddress, err := s.addressRepository.UpdateAddress(address)
	if err != nil {
		return response.AddressResponse{}, err
	}

	addressResponse := response.AddressResponseFormatter(updateAddress)

	return addressResponse, nil
}

func (s *userService) DeleteAddressUser(requestUser request.GetByUserIDRequest,  requestID request.GetByAddressIDRequest) error {
	getAddress, err := s.addressRepository.GetAddressByID(requestID.ID)
	if err != nil {
		return err
	}

	if getAddress.IDUser != requestUser.ID {
		return constants.ErrUnauthorized
	}

	err = s.addressRepository.DeleteAddress(getAddress.ID)
	if err != nil {
		return err
	}

	return nil
}