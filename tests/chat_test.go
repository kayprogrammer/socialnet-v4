package tests

import (
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestChat(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TESTING")
	app := fiber.New()
	db := Setup(t, app)
	// BASEURL := "/api/v4/chats"

	// Run Chat Endpoint Tests

	// Drop Tables and Close Connectiom
	DropData(db)
	CloseTestDatabase(db)
}
