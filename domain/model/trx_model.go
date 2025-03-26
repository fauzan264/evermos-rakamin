package model

import "time"

type Trx struct {
	ID 					int				`gorm:"type:int;primaryKey;autoIncrement"`
	IDUser 				int				`gorm:"type:int;not null"`
	AlamatPengiriman 	int				`gorm:"type:int;not null"`
	HargaTotal 			int				`gorm:"type:int;not null"`
	KodeInvoice 		string			`gorm:"type:varchar(255);not null"`
	MethodBayar 		string			`gorm:"type:varchar(255);not null"`
	CreatedAt 			time.Time		`gorm:"type:timestamp;not null:default:current_timestamp"`
	UpdatedAt 			time.Time		`gorm:"type:timestamp"`

	User				User			`gorm:"foreignKey:IDUser;references:ID"`
}

type DetailTrx struct {
	ID 					int				`gorm:"type:int;primaryKey;autoIncrement"`
	IDTrx 				int				`gorm:"type:int;not null"`
	IDLogProduk 		int				`gorm:"type:int;not null"`
	IDToko 				int				`gorm:"type:int;not null"`
	Kuantitas 			int				`gorm:"type:int;not null"`
	HargaTotal 			int				`gorm:"type:int;not null"`
	CreatedAt 			time.Time		`gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt 			time.Time		`gorm:"type:timestamp"`

	Trx					Trx 			`gorm:"foreignKey:IDTrx;references:ID"`
	LogProduct			LogProduct		`gorm:"foreignKey:IDLogProduk;references:ID"`
	Toko				Toko			`gorm:"foreignKey:IDToko;references:ID"`
}

func (Trx) TableName() string {
	return "trx"
}

func (DetailTrx) TableName() string {
	return "detail_trx"
}