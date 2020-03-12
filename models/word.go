package models

import (
	"github.com/jinzhu/gorm"
)

type Word struct {
	Word        string `json:"word"`
	Translation string `json:"translation"`
}

type Translation struct {
	gorm.Model
	Id int
	En string
	Ru string
}
