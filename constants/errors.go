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

// province city
var (
	ErrFailedFetchListProvince = errors.New("failed to fetch provinces")
	ErrFailedFetchListCity = errors.New("failed to fetch cities")
	ErrFailedFetchDetailProvince = errors.New("failed to fetch detail province")
	ErrFailedFetchDetailCity = errors.New("failed to fetch city")
)

// category
var (
	ErrCategoryNotFound = errors.New("No Data Category")
)