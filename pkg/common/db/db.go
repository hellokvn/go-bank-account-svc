package db

import (
	"log"

	"github.com/hellokvn/go-bank-account-svc/pkg/common/config"
	"github.com/hellokvn/go-bank-account-svc/pkg/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(c *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(c.PostgresUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Account{})

	return db
}
