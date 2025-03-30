package repositories

import (
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"gorm.io/gorm"
)

type trxRepository struct {
	db *gorm.DB
}

type TRXRepository interface {
	GetTRXByUserID(userID int, requestSearch request.TRXListRequest) ([]model.TRX, error)
	GetTRXUserByID(userID, id int) (model.TRX, error)
	CreateTRX(trx model.TRX) (model.TRX, error)
}

func NewTRXRepository(db *gorm.DB) *trxRepository {
	return &trxRepository{db}
}

func (r *trxRepository) GetTRXByUserID(userID int, requestSearch request.TRXListRequest) ([]model.TRX, error) {
	var listTRX []model.TRX

	offset := (requestSearch.Page - 1) * requestSearch.Limit

	query := r.db.Model(&model.TRX{})

	if requestSearch.Search != "" {
		query = query.Where("kode_invoice LIKE ?", "%"+ requestSearch.Search +"%")
	}

	query = query.Where("id_user = ?", userID).
				Preload("Alamat").
				Preload("DetailTRX.LogProduct.Toko").
				Preload("DetailTRX.LogProduct.Category").
				Preload("DetailTRX.LogProduct.Produk.PhotosProduct").
				Preload("DetailTRX.Toko")
				
				
	err := query.Limit(requestSearch.Limit).
				Offset(offset).
				Find(&listTRX).Error

	if err != nil {
		return listTRX, err
	}

	return listTRX, nil
}

func (r *trxRepository) GetTRXUserByID(userID, id int) (model.TRX, error) {
	var trx model.TRX
	err := r.db.Where("id_user = ? and id = ?", userID, id).
				Preload("Alamat").
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