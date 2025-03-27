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
	alamatRepository repositories.AlamatRepository
}

type UserService interface {
	GetUserByID(request request.GetByUserIDRequest) (response.UserResponse, error)
	UpdateUser(requestID request.GetByUserIDRequest, requestData request.UpdateProfileRequest) (response.UserResponse, error)
	GetMyAlamat(requestUser request.GetByUserIDRequest) ([]response.AddressResponse, error)
	GetAlamatUserByID(requestUser request.GetByUserIDRequest, requestID request.GetByAddressIDRequest) (response.AddressResponse, error)
	CreateAlamatUser(requestUser request.GetByUserIDRequest, requestData request.CreateAddressRequest) (response.AddressResponse, error)
	UpdateAlamatUser(requestUser request.GetByUserIDRequest,  requestID request.GetByAddressIDRequest, requestData request.UpdateAddressRequest) (response.AddressResponse, error)
	DeleteAlamatUser(requestUser request.GetByUserIDRequest,  requestID request.GetByAddressIDRequest) error

}
func NewUserService(repository repositories.UserRepository, alamatRepository repositories.AlamatRepository) *userService {
	return &userService{repository, alamatRepository}
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

func (s *userService) GetMyAlamat(requestUser request.GetByUserIDRequest) ([]response.AddressResponse, error) {
	getMyAddress, err := s.alamatRepository.GetAlamatByUserID(requestUser.ID)
	if err != nil {
		return []response.AddressResponse{}, err
	}

	var addressesResponse []response.AddressResponse
	for _, address := range getMyAddress {
		addressResponse := response.AddressResponse{
			ID : address.ID,
			JudulAlamat : address.JudulAlamat,
			NamaPenerima : address.NamaPenerima,
			NoTelp : address.NoTelp,
			DetailAlamat : address.DetailAlamat,
		}
		addressesResponse = append(addressesResponse, addressResponse)
	}

	return addressesResponse, nil
}

func (s *userService) GetAlamatUserByID(requestUser request.GetByUserIDRequest, requestID request.GetByAddressIDRequest) (response.AddressResponse, error) {
	address, err := s.alamatRepository.GetAlamatUserByID(requestUser.ID, requestID.ID)
	if err != nil {
		return response.AddressResponse{}, err
	}

	addressResponse := response.AddressResponse{
		ID : address.ID,
		JudulAlamat : address.JudulAlamat,
		NamaPenerima : address.NamaPenerima,
		NoTelp : address.NoTelp,
		DetailAlamat : address.DetailAlamat,
	}
	return addressResponse, nil
}

func (s *userService) CreateAlamatUser(requestUser request.GetByUserIDRequest, requestData request.CreateAddressRequest) (response.AddressResponse, error) {
	address := model.Alamat{
		IDUser: requestUser.ID,
		JudulAlamat: requestData.JudulAlamat,
		NamaPenerima: requestData.NamaPenerima,
		NoTelp: requestData.NoTelp,
		DetailAlamat: requestData.DetailAlamat,
		CreatedAt: time.Now(),
	}

	createAddress, err := s.alamatRepository.CreateAlamat(address)
	if err != nil {
		return response.AddressResponse{}, err
	}

	addressResponse := response.AddressResponse{
		ID : createAddress.ID,
		JudulAlamat : createAddress.JudulAlamat,
		NamaPenerima : createAddress. NamaPenerima,
		NoTelp : createAddress.NoTelp,
		DetailAlamat : createAddress.DetailAlamat,
	}

	return addressResponse, nil
}

func (s *userService) UpdateAlamatUser(requestUser request.GetByUserIDRequest,  requestID request.GetByAddressIDRequest, requestData request.UpdateAddressRequest) (response.AddressResponse, error) {
	getAddress, err := s.alamatRepository.GetAlamatByID(requestID.ID)
	if err != nil {
		return response.AddressResponse{}, err
	}

	if getAddress.IDUser != requestID.ID {
		return response.AddressResponse{}, constants.ErrUnauthorized
	}

	address := model.Alamat{
		ID : requestID.ID,
		IDUser : getAddress.IDUser,
		JudulAlamat : getAddress.JudulAlamat,
		NamaPenerima : requestData.NamaPenerima,
		NoTelp : requestData.NoTelp,
		DetailAlamat : requestData.DetailAlamat,
		CreatedAt : getAddress.CreatedAt,
		UpdatedAt : time.Now(),
	}

	updateAddress, err := s.alamatRepository.UpdateAlamat(address)
	if err != nil {
		return response.AddressResponse{}, err
	}

	addressResponse := response.AddressResponse{
		ID : updateAddress.ID,
		JudulAlamat : updateAddress.JudulAlamat,
		NamaPenerima : updateAddress.NamaPenerima,
		NoTelp : updateAddress.NoTelp,
		DetailAlamat : updateAddress.DetailAlamat,
	}

	return addressResponse, nil
}

func (s *userService) DeleteAlamatUser(requestUser request.GetByUserIDRequest,  requestID request.GetByAddressIDRequest) error {
	getAddress, err := s.alamatRepository.GetAlamatByID(requestID.ID)
	if err != nil {
		return err
	}

	if getAddress.IDUser != requestUser.ID {
		return constants.ErrUnauthorized
	}

	err = s.alamatRepository.DeleteAlamat(getAddress.ID)
	if err != nil {
		return err
	}

	return nil
}