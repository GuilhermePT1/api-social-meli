package database

import (
	"os"

	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=" + os.Getenv("POSTGRES_HOST") +
		" port=" + os.Getenv("POSTGRES_PORT") +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, db.AutoMigrate(&models.User{}, &models.Product{}, &models.Post{}, &models.Follow{})
}
