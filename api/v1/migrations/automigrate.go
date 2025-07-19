package migrations

import (
	"fmt"
	"marketplace-api/internal/advert"
	"marketplace-api/internal/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Automigrate(dsn string) {
	if dsn == "" {
		panic("DSN is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&advert.Advert{}, &user.User{}); err != nil {
		panic(err)
	}
	fmt.Println("Migrated successfully")
}
