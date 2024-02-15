package tests

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/stretchr/testify/assert"
)

func getPosts(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	post := CreatePost(db)
	user := post.Edges.Author
	token := AccessToken(db)
	t.Run("Retrieve Posts", func(t *testing.T) {
		url := fmt.Sprintf("%s/posts", baseUrl)
		req := httptest.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Posts fetched",
			"data": map[string]interface{}{
				"per_page": 50,
				"current_page": 1,
				"last_page": 1,
				"posts": []map[string]interface{}{
					{
						"author":      map[string]interface{}{
							"name":     schemas.FullName(user),
							"username": user.Username,
							"avatar":   nil,
						},
						"text":        post.Text,
						"slug":       post.Slug,
						"reactions_count": 0,
						"comments_count": 0,
						"image":       nil,
						"created_at": ConvertDateTime(post.CreatedAt),
						"updated_at": ConvertDateTime(post.UpdatedAt),
					},
				},
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}
func TestFeed(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TESTING")
	app := fiber.New()
	db := Setup(t, app)
	BASEURL := "/api/v4/feed"

	// Run Feed Endpoint Tests
	getPosts(t, app, db, BASEURL)

	// Drop Tables and Close Connectiom
	DropData(db)
	CloseTestDatabase(db)
}
