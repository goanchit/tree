package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", viper.GetString("db.host"), viper.GetString("db.db_user"), viper.GetString("db.db_password"), viper.GetString("db.db_name"), viper.GetString("db.port"))
}

func OpenDb() *gorm.DB {
	dsn := getDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	return db
}
