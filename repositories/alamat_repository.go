package repositories

import (
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"gorm.io/gorm"
)

type alamatRepository struct {
	db *gorm.DB
}

type AlamatRepository interface {
	CreateAlamat(alamat model.Alamat) (model.Alamat, error)
	GetAlamatByID(id int) (model.Alamat, error)
	GetAlamatByUserID(userID int) ([]model.Alamat, error)
	GetAlamatUserByID(userID int, id int) (model.Alamat, error)
	UpdateAlamat(alamat model.Alamat) (model.Alamat, error)
	DeleteAlamat(id int) error
}

func NewAlamatRepository(db *gorm.DB) *alamatRepository {
	return &alamatRepository{db}
}

func (r *alamatRepository) CreateAlamat(alamat model.Alamat) (model.Alamat, error) {
	err := r.db.Create(&alamat).Error
	if err != nil {
		return alamat, err
	}

	return alamat, nil
}

func (r *alamatRepository) GetAlamatByID(id int) (model.Alamat, error) {
	var alamat model.Alamat
	err := r.db.Where("id = ?", id).Find(&alamat).Error
	if err != nil {
		return alamat, err
	}

	return alamat, nil
}

func (r *alamatRepository) GetAlamatByUserID(userID int) ([]model.Alamat, error) {
	var alamat []model.Alamat
	err := r.db.Where("id_user = ?", userID).Find(&alamat).Error
	if err != nil {
		return alamat, err
	}

	return alamat, nil
}

func (r *alamatRepository) GetAlamatUserByID(userID int, id int) (model.Alamat, error) {
	var alamat model.Alamat
	err := r.db.Where("id_user = ? and id = ?", userID, id).First(&alamat).Error
	if err != nil {
		return alamat, err
	}

	return alamat, nil
}

func (r *alamatRepository) UpdateAlamat(alamat model.Alamat) (model.Alamat, error) {
	err := r.db.Save(&alamat).Error
	if err != nil {
		return alamat, err
	}

	return alamat, nil
}

func (r *alamatRepository) DeleteAlamat(id int) error {
	var alamat model.Alamat
	err := r.db.Where("id = ?", id).Delete(&alamat).Error
	if err != nil {
		return err
	}

	return nil
}
