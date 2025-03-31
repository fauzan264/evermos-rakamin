package repositories

import (
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"gorm.io/gorm"
)

type addressRepository struct {
	db *gorm.DB
}

type AddressRepository interface {
	CreateAddress(address model.Address) (model.Address, error)
	GetAddressByID(id int) (model.Address, error)
	GetAddressByUserID(userID int, requestSearch request.AddressListRequest) ([]model.Address, error)
	GetAddressUserByID(userID int, id int) (model.Address, error)
	UpdateAddress(address model.Address) (model.Address, error)
	DeleteAddress(id int) error
}

func NewAddressRepository(db *gorm.DB) *addressRepository {
	return &addressRepository{db}
}

func (r *addressRepository) CreateAddress(address model.Address) (model.Address, error) {
	err := r.db.Create(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}

func (r *addressRepository) GetAddressByID(id int) (model.Address, error) {
	var address model.Address
	err := r.db.Where("id = ?", id).First(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}

func (r *addressRepository) GetAddressByUserID(userID int, requestSearch request.AddressListRequest) ([]model.Address, error) {
	var listAddress []model.Address

	offset := (requestSearch.Page - 1) * requestSearch.Limit

	query := r.db.Model(&model.Address{})

	if requestSearch.Title != "" {
		query = query.Where("judul_alamat = ?", "%"+ requestSearch.Title +"%")
	}

	query = query.Where("id_user = ?", userID)

	err := query.Limit(requestSearch.Limit).
				Offset(offset).
				Find(&listAddress).Error

	if err != nil {
		return listAddress, err
	}

	return listAddress, nil
}

func (r *addressRepository) GetAddressUserByID(userID int, id int) (model.Address, error) {
	var address model.Address
	err := r.db.Where("id_user = ? and id = ?", userID, id).First(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}

func (r *addressRepository) UpdateAddress(address model.Address) (model.Address, error) {
	err := r.db.Save(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}

func (r *addressRepository) DeleteAddress(id int) error {
	var address model.Address
	err := r.db.Where("id = ?", id).Delete(&address).Error
	if err != nil {
		return err
	}

	return nil
}
