package storage

import (
	models "admin/model"
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	user     = "postgres"
	password = "postgrespw"
	host     = "localhost"
	port     = 32769
	dbname   = "postgres"
)

var doOnce sync.Once
var singleton *gorm.DB

func GetConnection() *gorm.DB {

	doOnce.Do(func() {

		connURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			user,
			password,
			host,
			port,
			dbname,
		)

		db, err := gorm.Open(postgres.Open(connURL), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			panic("failed to connect database")
		}

		singleton = db
	})

	return singleton
}

func init() {
	GetConnection().AutoMigrate(&models.Administrator{})
}
