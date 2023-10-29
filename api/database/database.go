package database

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ilovesoup20/japchae/ent"
)

// InitDB .
func InitDB() (*ent.Client, error) {
	client, err := ent.Open("mysql", "root:password@tcp(localhost:3306)/japchae")
	if err != nil {
		log.Fatalf("Failed opening mysql connection %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed creating schema resources %v", err)
	}

	return client, err
}
