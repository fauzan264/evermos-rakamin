package constants

import "errors"

var (
	ErrWrongUserOrPassword = errors.New("No Telp atau kata sandi salah")
	ErrRecordNotFound = errors.New("record not found")
	ErrStoreNotFound = errors.New("Toko tidak ditemukan")
	ErrProductNotFound = errors.New("No Data Product")
	ErrTrxNotFound = errors.New("No Data Trx")
	ErrInvalidToken = errors.New("Invalid token")
	ErrUnauthorized = errors.New("Unauthorized")
	ErrInvalidDateFormat = errors.New("Invalid date format. Use DD/MM/YYYY")
)