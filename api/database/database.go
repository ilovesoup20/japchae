package database

import (
	"context"
	"log"

	"github.com/ilovesoup20/japchae/ent"
)

// InitDB a
func InitDB() (*ent.Client, error) {
	client, err := ent.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		log.Fatalf("Failed opening mysql connection %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed creating schema resources %v", err)
	}

	return client, err
}
