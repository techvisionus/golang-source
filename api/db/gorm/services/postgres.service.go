package services

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize(connection string, config *gorm.Config) *gorm.DB {
	con, err := gorm.Open(postgres.Open(connection), config)
	if nil != err {
		log.Error("Error connection DB. err=", err)
		panic(err)
	}
	DB = con
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
