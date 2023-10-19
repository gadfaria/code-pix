package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gadfaria/codepix/domain/model"

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

	var dialector gorm.Dialector
	var dsn string

	if env != "test" {
		dsn = os.Getenv("dsn")
		dialector = postgres.Open(dsn)
	} else {
		dsn = os.Getenv("dsnTest")
		dialector = sqlite.Open(dsn)
	}

	db, err = gorm.Open(dialector, &gorm.Config{})

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
