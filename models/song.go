package models

import (
	"github.com/jinzhu/gorm"
)

type Song struct {
	gorm.Model
	FilePath string
}
