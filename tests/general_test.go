package tests

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/utils"
	"github.com/stretchr/testify/assert"
)

func getSiteDetails(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	t.Run("Get Site Details", func(t *testing.T) {
		url := fmt.Sprintf("%s/site-detail", baseUrl)
		req := httptest.NewRequest("GET", url, nil)
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "success", body["status"])
		assert.Equal(t, "Site Details Fetched!", body["message"])
		dataKeys := []string{"address", "email", "fb", "ig", "name", "phone", "tw", "wh"}
		assert.Equal(t, true, utils.KeysExistInMap(dataKeys, body["data"].(map[string]interface{})))
	})
}

func TestGeneral(t *testing.T) {
	app := fiber.New()
	db := Setup(t, app)
	BASEURL := "/api/v4/general"

	// Run General Endpoint Tests
	getSiteDetails(t, app, db, BASEURL)

	// Drop Tables and Close Connectiom
	DropData(db)
	CloseTestDatabase(db)
}