package model

import "time"

type Produk struct {
	ID 				uint64		`gorm:"primaryKey;autoIncrement"`
	NamaProduk 		string		`gorm:"type:varchar(255);not null"`
	Slug 			string		`gorm:"type:varchar(255);not null"`
	HargaReseller 	string		`gorm:"type:varchar(255);not null"`
	HargaKonsumen 	string		`gorm:"type:varchar(255);not null"`
	Stok 			int			`gorm:"type:int;not null;default:0"`
	Deskripsi 		string		`gorm:"type:text;not null"`
	CreatedAt 		time.Time	`gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt 		time.Time	`gorm:"type:timestamp"`
	IDToko 			uint64		`gorm:"type:int;not null"`
	IDCategory 		uint64		`gorm:"type:int;not null"`

	Toko			Toko		`gorm:"foreignKey:IDToko;references:ID"`
	Category		Category 	`gorm:"foreignKey:IDCategory;references:ID"`
}

type LogProduk struct {
	ID 				int			`gorm:"type:int;primaryKey;autoIncrement"`
	IDProduk 		int			`gorm:"type:int;not null"`
	NamaProduk 		string		`gorm:"type:varchar(255);not null"`
	Slug 			string		`gorm:"type:varchar(255);not null"`
	HargaReseller 	string		`gorm:"type:varchar(255);not null"`
	HargaKonsumen 	string		`gorm:"type:varchar(255);not null"`
	Stock			int			`gorm:"type:int;not null;default:0"`
	Deskripsi 		string		`gorm:"type:text;not null"`
	CreatedAt 		time.Time	`gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt 		time.Time	`gorm:"type:timestamp"`
	IDToko 			int			`gorm:"type:int;not null"`
	IDCategory 		int			`gorm:"type:int;not null"`

	Produk			Produk		`gorm:"foreignKey:IDProduk;references:ID"`
	Toko			Toko		`gorm:"foreignKey:IDToko;references:ID"`
	Category		Category	`gorm:"foreignKey:IDCategory;references:ID"`
}

type FotoProduk struct {
	ID 				int			`gorm:"primaryKey;autoIncrement"`
	IDProduk 		int			`gorm:"type:int;not null"`	
	URL 			string		`gorm:"type:varchar(255);not null"`
	CreatedAt 		time.Time	`gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt 		time.Time	`gorm:"type:timestamp"`

	Produk 			Produk		`gorm:"foreignKey:IDProduk;references:ID"`
}

func (Produk) TableName() string {
	return "produk"
}

func (LogProduk) TableName() string {
	return "log_produk"
}

func (FotoProduk) TableName() string {
	return "foto_produk"
}