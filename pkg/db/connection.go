package db

import (
	"fmt"

	"github.com/msazad/go-Ecommerce/pkg/config"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config)(*gorm.DB,error){
	psqlinfo:=fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s dbname=%s",cfg.DBHost,cfg.DBUser,cfg.DBPassword)
}