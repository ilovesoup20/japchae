package database

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ilovesoup20/japchae/ent"
	"github.com/ilovesoup20/japchae/model"
)

// InitEntDB .
func InitEntDB() (*ent.Client, error) {
	client, err := ent.Open("mysql", "root:password@tcp(localhost:3306)/japchae")
	if err != nil {
		log.Fatalf("Failed opening mysql connection %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed creating schema resources %v", err)
	}

	return client, err
}

// InitGormDB .
func InitGormDB() (*gorm.DB, error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/japchae?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})
	return db, nil
}
