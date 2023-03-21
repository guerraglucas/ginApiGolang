package databases

import (
	m "github.com/guerraglucas/ginApi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	pgHost     = "localhost"
	pgPort     = "5432"
	pgUser     = "root"
	pgPassword = "root"
	pgDatabase = "root"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToPostgres() {
	dsn := "host=" + pgHost + " port=" + pgPort + " user=" + pgUser + " password=" + pgPassword + " dbname=" + pgDatabase + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&m.Student{})
}

func CloseConnection() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}
