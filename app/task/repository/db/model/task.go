package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	//Uid       uint   `gorm:"not null"`
	//Title     string `json:"title"`
	//Content   string `gorm:"type:longtext"`
	//Status    int    `gorm:"default:0"`
	//StartTime int64
	//EndTime   int64
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index; not null"`
	Content   string `gorm:"type:longtext"`
	Status    int    `gorm:"default:0"`
	StartTime int64
	EndTime   int64 `gorm:"default:0"`
}
