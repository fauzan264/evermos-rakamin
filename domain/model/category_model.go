package model

import "time"

type Category struct {
	ID 				int				`gorm:"type:int;primaryKey;autoIncrement"`
	NamaCategory 	string			`gorm:"type:varchar(255);not null"`
	CreatedAt 		time.Time		`gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt 		time.Time		`gorm:"type:timestamp"`
}

func (Category) TableName() string {
	return "category"
}