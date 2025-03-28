package model

import "time"

type TRX struct {
	ID 					int				`gorm:"type:int;primaryKey;autoIncrement"`
	IDUser 				int				`gorm:"type:int;not null"`
	AlamatPengiriman 	int				`gorm:"type:int;not null"`
	HargaTotal 			int				`gorm:"type:int;not null"`
	KodeInvoice 		string			`gorm:"type:varchar(255);not null"`
	MethodBayar 		string			`gorm:"type:varchar(255);not null"`
	CreatedAt 			time.Time		`gorm:"type:timestamp;not null:default:current_timestamp"`
	UpdatedAt 			time.Time		`gorm:"type:timestamp"`

	Alamat				Alamat			`gorm:"foreignKey:AlamatPengiriman;references:ID"`
	User				User			`gorm:"foreignKey:IDUser;references:ID"`
	DetailTRX   		[]DetailTRX 	`gorm:"foreignKey:IDTrx;references:ID"`
}

type DetailTRX struct {
	ID 					int				`gorm:"type:int;primaryKey;autoIncrement"`
	IDTrx 				int				`gorm:"type:int;not null"`
	IDLogProduk 		int				`gorm:"type:int;not null"`
	IDToko 				int				`gorm:"type:int;not null"`
	Kuantitas 			int				`gorm:"type:int;not null"`
	HargaTotal 			int				`gorm:"type:int;not null"`
	CreatedAt 			time.Time		`gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt 			time.Time		`gorm:"type:timestamp"`

	LogProduct			LogProduct		`gorm:"foreignKey:IDLogProduk;references:ID"`
	Toko				Toko			`gorm:"foreignKey:IDToko;references:ID"`
}

func (TRX) TableName() string {
	return "trx"
}

func (DetailTRX) TableName() string {
	return "detail_trx"
}