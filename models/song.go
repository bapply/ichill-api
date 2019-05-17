package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Song struct {
	gorm.Model
	FilePath string
	InfoID   int
}

func SongHandler(c *gin.Context) {
	song := Song{}

	switch c.Request.Method {
	case "GET":
		song.Index(c)
	}
}
