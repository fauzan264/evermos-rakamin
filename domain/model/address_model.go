package model

import "time"

type Address struct {
	ID 				int				`gorm:"type:int;primaryKey;autoIncrement"`
	IDUser 			int				`gorm:"type:int"`
	JudulAlamat 	string			`gorm:"type:varchar(255);not null"`
	NamaPenerima 	string			`gorm:"type:varchar(255);not null"`
	NoTelp 			string			`gorm:"type:varchar(255);not null"`
	DetailAlamat 	string			`gorm:"type:varchar(255);not null"`
	CreatedAt 		time.Time		`gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt 		time.Time		`gorm:"type:timestamp"`

	User			User			`gorm:"foreignKey:IDUser;references:ID"`
}

func (Address) TableName() string {
	return "alamat"
}