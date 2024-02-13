package tests

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
)

func getChats(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	CreateChat(db)
	t.Run("Retrieve Chats", func(t *testing.T) {
		url := baseUrl
		req := httptest.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken(db)))
		res, _ := app.Test(req)

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

func sendMessage(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	// Drop and Create Chat
	chatManager.DropData(db)
	chat := CreateChat(db)
	sender := chat.Edges.Owner
	t.Run("Send Message", func(t *testing.T) {
		url := baseUrl
		invalidUUID := uuid.New()
		text := "JESUS is KING"
		messageData := schemas.MessageCreateSchema{ChatID: &invalidUUID, Text: &text}
		// Test for valid response for invalid chat id
		res := ProcessTestBody(t, app, url, "POST", messageData, AccessToken(db))
		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "User has no chat with that ID", body["message"])

		// Test for valid response for valid entry
		messageData.ChatID = &chat.ID
		res = ProcessTestBody(t, app, url, "POST", messageData, AccessToken(db))
		// Assert Status code
		assert.Equal(t, 201, res.StatusCode)
		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Message sent",
			"data": map[string]interface{}{
				"id":      body["data"].(map[string]interface{})["id"],
				"chat_id": chat.ID,
				"sender": map[string]interface{}{
					"name":     fmt.Sprintf("%s %s", sender.FirstName, sender.LastName),
					"username": sender.Username,
					"avatar":   nil,
				},
				"text":             messageData.Text,
				"created_at":       body["data"].(map[string]interface{})["created_at"],
				"updated_at":       body["data"].(map[string]interface{})["updated_at"],
				"file_upload_data": nil,
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func TestChat(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TESTING")
	app := fiber.New()
	db := Setup(t, app)
	BASEURL := "/api/v4/chats"

	// Run Chat Endpoint Tests
	getChats(t, app, db, BASEURL)
	sendMessage(t, app, db, BASEURL)

	// Drop Tables and Close Connectiom
	DropData(db)
	CloseTestDatabase(db)
}
