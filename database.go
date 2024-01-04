package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kayprogrammer/socialnet-v4/ent"
	_ "github.com/lib/pq"
	"github.com/kayprogrammer/socialnet-v4/config"
)

func ConnectDb() {
	cfg := config.GetConfig()

	dbUrlTemplate := "host=%s port=%s user=%s dbname=%s password=%s"

	dbUrl := fmt.Sprintf(
		dbUrlTemplate,
		cfg.PostgresServer,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresDB,
		cfg.PostgresPassword,
	)

	client, err := ent.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
