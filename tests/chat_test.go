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
	chat := CreateChat(db)
	sender := chat.Edges.Owner
	token := AccessToken(db)
	t.Run("Send Message", func(t *testing.T) {
		url := baseUrl
		invalidUUID := uuid.New()
		text := "JESUS is KING"
		messageData := schemas.MessageCreateSchema{ChatID: &invalidUUID, Text: &text}
		// Test for valid response for invalid chat id
		res := ProcessTestBody(t, app, url, "POST", messageData, token)
		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "User has no chat with that ID", body["message"])

		// Test for valid response for valid entry
		messageData.ChatID = &chat.ID
		res = ProcessTestBody(t, app, url, "POST", messageData, token)
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
					"name":     schemas.FullName(sender),
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

func getChatMessages(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	message := CreateMessage(db)
	chat := message.Edges.Chat
	owner := chat.Edges.Owner
	token := AccessToken(db)
	t.Run("Retrieve Chat Messages", func(t *testing.T) {
		invalidChatID := uuid.New()
		url := fmt.Sprintf("%s/%s", baseUrl, invalidChatID)
		req := httptest.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		res, _ := app.Test(req)

		// Verify the request fails with invalid chat ID
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, 404, res.StatusCode)
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "User has no chat with that ID", body["message"])

		// Verify the request succeeds with valid chat ID
		url = fmt.Sprintf("%s/%s", baseUrl, chat.ID)
		req = httptest.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		res, _ = app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		ownerData := map[string]interface{}{
			"name":     schemas.FullName(owner),
			"username": owner.Username,
			"avatar":   nil,
		}
		recipientUser := chat.Edges.Users[0]
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Messages fetched",
			"data": map[string]interface{}{
				"chat": map[string]interface{}{
					"id":          chat.ID,
					"name":        chat.Name,
					"owner":       ownerData,
					"ctype":       chat.Ctype,
					"description": chat.Description,
					"image":       nil,
					"latest_message": map[string]interface{}{
						"sender": ownerData,
						"text":   message.Text,
						"file":   nil,
					},
					"created_at": ConvertDateTime(chat.CreatedAt),
					"updated_at": ConvertDateTime(chat.UpdatedAt),
				},
				"messages": map[string]interface{}{
					"per_page":     400,
					"current_page": 1,
					"last_page":    1,
					"items": []map[string]interface{}{
						{
							"id":         message.ID,
							"chat_id":    chat.ID,
							"sender":     ownerData,
							"text":       message.Text,
							"file":       nil,
							"created_at": ConvertDateTime(message.CreatedAt),
							"updated_at": ConvertDateTime(message.UpdatedAt),
						},
					},
				},
				"users": []map[string]interface{}{
					{
						"name":     schemas.FullName(recipientUser),
						"username": recipientUser.Username,
						"avatar":   nil,
					},
				},
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func updateGroupChat(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	chat := CreateGroupChat(db)
	user := chat.Edges.Users[0]
	token := AccessToken(db)
	t.Run("Update Group Chat", func(t *testing.T) {
		url := fmt.Sprintf("%s/%s", baseUrl, uuid.New())
		name := "Updated Group chat name"
		desc := "Updated group chat description"
		chatData := schemas.GroupChatInputSchema{Name: &name, Description: &desc}

		// Test for valid response for invalid chat id
		res := ProcessTestBody(t, app, url, "PATCH", chatData, token)
		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "User owns no group chat with that ID", body["message"])

		// Test for valid response for valid entry
		url = fmt.Sprintf("%s/%s", baseUrl, chat.ID)
		res = ProcessTestBody(t, app, url, "PATCH", chatData, token)
		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)
		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Chat updated",
			"data": map[string]interface{}{
				"id":          chat.ID,
				"name":        chatData.Name,
				"description": chatData.Description,
				"users": []map[string]interface{}{
					{
						"name":     schemas.FullName(user),
						"username": user.Username,
						"avatar":   nil,
					},
				},
				"file_upload_data": nil,
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))

		// You can test for other error responses yourself

	})
}

func deleteGroupChat(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	chat := CreateGroupChat(db)
	token := AccessToken(db)
	t.Run("Delete Group Chat", func(t *testing.T) {
		url := fmt.Sprintf("%s/%s", baseUrl, uuid.New())
		// Test for valid response for invalid chat id
		req := httptest.NewRequest("DELETE", url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "User owns no group chat with that ID", body["message"])

		// Test for valid response for valid entry
		url = fmt.Sprintf("%s/%s", baseUrl, chat.ID)
		req = httptest.NewRequest("DELETE", url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		res, _ = app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)
		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Group Chat Deleted",
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
		// You can test for other error responses yourself

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
	getChatMessages(t, app, db, BASEURL)
	updateGroupChat(t, app, db, BASEURL)
	deleteGroupChat(t, app, db, BASEURL)

	// Drop Tables and Close Connectiom
	DropData(db)
	CloseTestDatabase(db)
}
