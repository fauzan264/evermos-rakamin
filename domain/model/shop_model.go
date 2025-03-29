package model

import "time"

type Shop struct {
	ID 			int				`gorm:"type:int;primaryKey;autoIncrement"`
	IDUser 		int				`gorm:"type:int;not null"`
	NamaToko 	string			`gorm:"type:varchar(255)"`
	URLFoto 	string			`gorm:"type:varchar(255)"`
	CreatedAt 	time.Time		`gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt 	time.Time		`gorm:"type:timestamp"`

	User		User			`gorm:"foreignKey:IDUser;references:ID"`
}

func (Shop) TableName() string {
	return "toko"
}