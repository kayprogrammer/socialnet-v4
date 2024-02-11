package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/config"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/migrate"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/routes"
	"github.com/stretchr/testify/assert"
)

func CreateTables(db *ent.Client) {
	if err := db.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func DropData(db *ent.Client) {
	// Delete all entities
	db.Country.Delete().Exec(managers.Ctx)
	db.Region.Delete().Exec(managers.Ctx)
	db.City.Delete().Exec(managers.Ctx)
	db.User.Delete().Exec(managers.Ctx)
	db.Otp.Delete().Exec(managers.Ctx)

	db.Friend.Delete().Exec(managers.Ctx)
	db.Notification.Delete().Exec(managers.Ctx)

	db.SiteDetail.Delete().Exec(managers.Ctx)

	db.Chat.Delete().Exec(managers.Ctx)
	db.Message.Delete().Exec(managers.Ctx)

	db.Post.Delete().Exec(managers.Ctx)
	db.Comment.Delete().Exec(managers.Ctx)
	db.Reply.Delete().Exec(managers.Ctx)
	db.Reaction.Delete().Exec(managers.Ctx)
}

func SetupTestDatabase() *ent.Client {
	cfg := config.GetConfig()
	dbUrl := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		cfg.PostgresServer,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.TestPostgresDB,
		cfg.PostgresPassword,
	)
	client, err := ent.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	return client
}

func CloseTestDatabase(db *ent.Client) {
	if err := db.Close(); err != nil {
		log.Fatal("Failed to close database connection: " + err.Error())
	}
}

func Setup(t *testing.T, app *fiber.App) *ent.Client {
	// Set up the test database
	db := SetupTestDatabase()

	// Inject your test database and environment text into the Fiber app's context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		c.Locals("env", "test")
		return c.Next()
	})
	routes.SetupRoutes(app)
	DropData(db)
	CreateTables(db)
	return db
}

func ParseResponseBody(t *testing.T, b io.ReadCloser) interface{} {
	body, _ := io.ReadAll(b)
	// Parse the response body as JSON
	responseBody := make(map[string]interface{})
	err := json.Unmarshal(body, &responseBody)
	if err != nil {
		t.Errorf("error parsing response body as JSON: %s", err)
	}
	return responseBody
}

func ProcessTestBody(t *testing.T, app *fiber.App, url string, method string, body interface{}, access ...string) *http.Response {
	// Marshal the test data to JSON
	requestBytes, err := json.Marshal(body)
	requestBody := bytes.NewReader(requestBytes)
	assert.Nil(t, err)
	req := httptest.NewRequest(method, url, requestBody)
	req.Header.Set("Content-Type", "application/json")
	if access != nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access[0]))
	}
	res, err := app.Test(req)
	if err != nil {
		log.Println(err)
	}
	return res
}
