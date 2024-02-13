package tests

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/stretchr/testify/assert"
)

func getChats(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	CreateChat(db)
	t.Run("Retrieve Chats", func(t *testing.T) {
		url := baseUrl
		req := httptest.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken(db)))
		res, _ := app.Test(req)

		CreateCity(db)
		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "success", body["status"])
		assert.Equal(t, "Chats fetched", body["message"])
		data, _ := json.Marshal(body["data"])
		assert.Equal(t, true, (len(data) > 0))
	})
}


func TestChat(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TESTING")
	app := fiber.New()
	db := Setup(t, app)
	BASEURL := "/api/v4/chats"

	// Run Chat Endpoint Tests
	getChats(t, app, db, BASEURL)


	// Drop Tables and Close Connectiom
	DropData(db)
	CloseTestDatabase(db)
}
