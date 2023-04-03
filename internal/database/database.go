package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error

	var DATABASE_URI string = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnvWithDefault("MYSQL_USERNAME", "root"),
		getEnvWithDefault("MYSQL_PASSWORD", "Amagi@560076"),
		getEnvWithDefault("MYSQL_HOST", "localhost"),
		getEnvWithDefault("MYSQL_DB", "nordend"),
	)

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	return nil
}

func getEnvWithDefault(key, value string) string {
	str := os.Getenv(key)
	if len(str) > 0 {
		return str
	}

	return value
}
