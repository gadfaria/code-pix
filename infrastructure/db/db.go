package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gadfaria/code-pix/domain/model"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB(env string) *gorm.DB {
	var db *gorm.DB
	var err error

	if env != "test" {
		dsn := os.Getenv("dsn")
		var dialector gorm.Dialector
		switch os.Getenv("dbType") {
		case "postgres":
			dialector = postgres.Open(dsn)
		case "sqlite":
			dialector = sqlite.Open(dsn)
		default:
			log.Fatalf("Invalid database type: %s", os.Getenv("dbType"))
		}
		db, err = gorm.Open(dialector, &gorm.Config{})
	} else {
		dsn := os.Getenv("dsnTest")
		var dialector gorm.Dialector
		switch os.Getenv("dbTypeTest") {
		case "postgres":
			println("entrou aqui")
			dialector = postgres.Open(dsn)
		case "sqlite":
			dialector = sqlite.Open(dsn)
		default:
			log.Fatalf("Invalid database type: %s", os.Getenv("dbTypeTest"))
		}
		db, err = gorm.Open(dialector, &gorm.Config{})
	}

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("debug") == "true" {
		db = db.Debug()
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db
}
