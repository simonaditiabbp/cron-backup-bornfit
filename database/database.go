package database

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db_prod_bornfit   *gorm.DB
	db_backup_bornfit *gorm.DB
)

type Database struct {
	*gorm.DB
}

func SetupProd() {
	username := os.Getenv("PROD_DB_USER")
	if username == "" {
		username = "postgres"
	}
	password := os.Getenv("PROD_DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}
	host := os.Getenv("PROD_DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PROD_DB_PORT")
	if port == "" {
		port = "5432"
	}
	database := os.Getenv("PROD_DB_NAME")
	if database == "" {
		environment := os.Getenv("ENVIRONMENT") == "PRODUCTION"
		if environment {
			database = "bornfit"
		} else {
			database = "bornfit_v4"
		}
	}

	DB, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password)
	if err != nil {
		fmt.Println("failed to connect database prod: ", err)
	}

	DB.DB().SetMaxOpenConns(5)
	DB.DB().SetMaxIdleConns(3)
	DB.DB().SetConnMaxLifetime(1 * time.Minute)

	db_prod_bornfit = DB
}

func SetupBackup() {
	username := os.Getenv("BACKUP_DB_USER")
	if username == "" {
		username = "postgres"
	}
	password := os.Getenv("BACKUP_DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}
	host := os.Getenv("BACKUP_DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("BACKUP_DB_PORT")
	if port == "" {
		port = "5432"
	}
	database := os.Getenv("BACKUP_DB_NAME")
	if database == "" {
		environment := os.Getenv("ENVIRONMENT") == "PRODUCTION"
		if environment {
			database = "bornfit_backup"
		} else {
			database = "bornfit_backup_v4"
		}
	}

	DB, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password)
	if err != nil {
		fmt.Println("failed to connect database backup: ", err)
	}
	DB.DB().SetMaxOpenConns(5)
	DB.DB().SetMaxIdleConns(3)
	DB.DB().SetConnMaxLifetime(1 * time.Minute)

	db_backup_bornfit = DB
}

func GetDBProd() *gorm.DB {
	return db_prod_bornfit
}

func GetDBBackup() *gorm.DB {
	return db_backup_bornfit
}
