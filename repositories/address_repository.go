package repositories

import (
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"gorm.io/gorm"
)

type addressRepository struct {
	db *gorm.DB
}

type AddressRepository interface {
	CreateAddress(address model.Address) (model.Address, error)
	GetAddressByID(id int) (model.Address, error)
	GetAddressByUserID(userID int) ([]model.Address, error)
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

func (r *addressRepository) GetAddressByUserID(userID int) ([]model.Address, error) {
	var address []model.Address
	err := r.db.Where("id_user = ?", userID).Find(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
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
