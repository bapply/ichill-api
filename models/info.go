package models

import "github.com/jinzhu/gorm"

type Info struct {
	gorm.Model
	SongID   int
	Title    string
	Artist   string
	Category string
	Approved bool
}
