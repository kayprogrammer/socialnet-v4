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

func getCities(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	t.Run("Retrieve cities", func(t *testing.T) {
		// Test for valid response for non-existent city name query
		url := fmt.Sprintf("%s/cities?name=non_existent", baseUrl)
		req := httptest.NewRequest("GET", url, nil)
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "success", body["status"])
		assert.Equal(t, "No match found", body["message"])
		assert.Equal(t, []interface{}{}, body["data"])

		// Test for valid response for existent city name query
		city := CreateCity(db)
		url = fmt.Sprintf("%s/cities?name=%s", baseUrl, city.Name)
		req = httptest.NewRequest("GET", url, nil)
		res, _ = app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "success", body["status"])
		assert.Equal(t, "Cities Fetched", body["message"])
		data, _ := json.Marshal(body["data"])
		assert.Equal(t, true, (len(data) > 0))
	})
}

func TestProfiles(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TESTING")
	app := fiber.New()
	db := Setup(t, app)
	BASEURL := "/api/v4/profiles"

	// Run Profiles Endpoint Tests
	getCities(t, app, db, BASEURL)

	// Drop Tables and Close Connectiom
	DropData(db)
	CloseTestDatabase(db)
}
