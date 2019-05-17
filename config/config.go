package config

import (
	"github.com/bapply/ichill-api/database"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AppConfig struct {
	Name   string
	DB     *gorm.DB
	Router *gin.Engine
}

var App = AppConfig{
	Name:   "iChill",
	DB:     database.Connect(),
	Router: gin.Default(),
}
