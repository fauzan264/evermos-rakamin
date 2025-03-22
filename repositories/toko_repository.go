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
	GetListToko(page, limit int, name string) ([]model.Toko, error)
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

func (r *tokoRepository) GetListToko(page, limit int, name string) ([]model.Toko, error) {
	var listToko []model.Toko

	offset := (page - 1) * limit

	query := r.db.Model(&model.Toko{})

	// Tambahkan filter pencarian jika name tidak kosong
	if name != "" {
		query = query.Where("nama_toko LIKE ?", "%"+name+"%")
	}

	// Ambil data toko dengan pagination
	err := query.Limit(limit).
		Offset(offset).
		Find(&listToko).Error

	if err != nil {
		return nil, err // Kembalikan nil jika ada error
	}

	return listToko, nil // Tetap return array kosong jika tidak ada data
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
	err := r.db.Where("id_user = ", userID).First(&toko).Error
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