package tests

import (
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestProfiles(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TESTING")
	app := fiber.New()
	db := Setup(t, app)
	// BASEURL := "/api/v4/profiles"

	// Run Profiles Endpoint Tests

	// Drop Tables and Close Connectiom
	DropData(db)
	CloseTestDatabase(db)
}
