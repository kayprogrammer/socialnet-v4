package database

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/config"
	"github.com/kayprogrammer/socialnet-v4/ent"
	_ "github.com/lib/pq"
)
var Database *ent.Client

func ConnectDb() *ent.Client {
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
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func DatabaseMiddleware(c *fiber.Ctx) error {
	Database = ConnectDb()
	c.Locals("db", Database)
	defer Database.Close()
	return c.Next()
}