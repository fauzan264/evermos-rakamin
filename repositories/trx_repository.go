package repositories

import (
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"gorm.io/gorm"
)

type trxRepository struct {
	db *gorm.DB
}

type TRXRepository interface {
	GetTRXByUserID(userID, page, limit int, search string) ([]model.TRX, error)
	GetTRXUserByID(userID, id int) (model.TRX, error)
	CreateTRX(trx model.TRX) (model.TRX, error)
}

func NewTRXRepository(db *gorm.DB) *trxRepository {
	return &trxRepository{db}
}

func (r *trxRepository) GetTRXByUserID(userID, page, limit int, search string) ([]model.TRX, error) {
	var listTRX []model.TRX

	offset := (page - 1) * limit

	query := r.db.Model(&model.TRX{})

	if search != "" {
		query = query.Where("kode_invoice LIKE ?", "%"+ search +"%")
	}

	err := query.Preload("Alamat").
				Preload("DetailTRX.LogProduct.Toko").
				Preload("DetailTRX.LogProduct.Category").
				Preload("DetailTRX.LogProduct.Produk.PhotosProduct").
				Preload("DetailTRX.Toko").
				Limit(limit).
				Offset(offset).
				Find(&listTRX).Error

	if err != nil {
		return listTRX, err
	}

	return listTRX, nil
}

func (r *trxRepository) GetTRXUserByID(userID, id int) (model.TRX, error) {
	var trx model.TRX
	err := r.db.Preload("Alamat").
				Preload("DetailTRX.LogProduct.Toko").
				Preload("DetailTRX.LogProduct.Category").
				Preload("DetailTRX.LogProduct.Produk.PhotosProduct").
				Preload("DetailTRX.Toko").
				First(&trx).Error

	if err != nil {
		return trx, err
	}

	return trx, nil
}

func (r *trxRepository) CreateTRX(trx model.TRX) (model.TRX, error) {
	err := r.db.Create(&trx).Error
	if err != nil {
		return trx, err
	}

	return trx, nil
}