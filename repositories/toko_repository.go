package repositories

import (
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"gorm.io/gorm"
)

type tokoRepository struct {
	db *gorm.DB
}

type TokoRepository interface {
	CreateToko(toko model.Toko) error
	GetAllToko() ([]model.Toko, error)
	GetTokoByID(id int) (model.Toko, error)
	GetTokoByUserID(userID int) (model.Toko, error)
	UpdateToko(toko model.Toko) (model.Toko, error)
}

func NewTokoRepository(db *gorm.DB) *tokoRepository {
	return &tokoRepository{db}
}

func (r *tokoRepository) CreateToko(toko model.Toko) error {
	err := r.db.Create(&toko).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *tokoRepository) GetAllToko() ([]model.Toko, error) {
	var allToko []model.Toko

	err := r.db.Find(&allToko).Error
	if err != nil {
		return allToko, err
	}

	return allToko, nil
}

func (r *tokoRepository) GetTokoByID(id int) (model.Toko, error) {
	var toko model.Toko
	err := r.db.Where("id = ?", id).Find(&toko).Error
	if err != nil {
		return toko, err
	}

	return toko, nil
}

func (r *tokoRepository) GetTokoByUserID(userID int) (model.Toko, error) {
	var toko model.Toko
	err := r.db.Where("id_user = ", userID).Find(&toko).Error
	if err != nil {
		return toko, err
	}
	
	return toko, nil
}

func (r *tokoRepository) UpdateToko(toko model.Toko) (model.Toko, error) {
	err := r.db.Save(&toko).Error
	if err != nil {
		return toko, err
	}

	return toko, nil
}