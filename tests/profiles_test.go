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
	"github.com/kayprogrammer/socialnet-v4/utils"
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

func getProfile(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	t.Run("Retrieve Profile", func(t *testing.T) {
		// Test for valid response for non-existent username
		url := fmt.Sprintf("%s/profile/invalid_username", baseUrl)
		req := httptest.NewRequest("GET", url, nil)
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)

		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "No user with that username", body["message"])

		// Test for valid response for valid entry
		user := CreateTestVerifiedUser(db)
		url = fmt.Sprintf("%s/profile/%s", baseUrl, user.Username)
		req = httptest.NewRequest("GET", url, nil)
		res, _ = app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "success", body["status"])
		assert.Equal(t, "User details fetched", body["message"])
	})
}

func updateProfile(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	firstName := "TestUpdated"
	lastName := "VerifiedUpdated"
	bio := "Updated my bio"
	t.Run("Update Profile", func(t *testing.T) {
		url := fmt.Sprintf("%s/profile", baseUrl)
		updateProfileData := schemas.ProfileUpdateSchema{
			FirstName: &firstName,
			LastName:  &lastName,
			Bio:       &bio,
		}

		// Test for valid response for valid entry
		res := ProcessTestBody(t, app, url, "PATCH", updateProfileData, AccessToken(db))
		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "success", body["status"])
		assert.Equal(t, "User updated", body["message"])
	})
}

func deleteProfile(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	t.Run("Delete Profile", func(t *testing.T) {
		url := fmt.Sprintf("%s/profile", baseUrl)
		userData := schemas.DeleteUserSchema{
			Password: "invalid_pass",
		}

		// Test for valid response for invalid entry
		res := ProcessTestBody(t, app, url, "POST", userData, AccessToken(db))
		// Assert Status code
		assert.Equal(t, 422, res.StatusCode)

		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_INVALID_ENTRY, body["code"])
		assert.Equal(t, "Invalid Entry", body["message"])

		// Test for valid response for valid entry
		userData.Password = "testpassword"
		res = ProcessTestBody(t, app, url, "POST", userData, AccessToken(db))
		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)
		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "success", body["status"])
		assert.Equal(t, "User deleted", body["message"])
	})
}

func TestProfiles(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TESTING")
	app := fiber.New()
	db := Setup(t, app)
	BASEURL := "/api/v4/profiles"

	// Run Profiles Endpoint Tests
	getCities(t, app, db, BASEURL)
	getProfile(t, app, db, BASEURL)
	updateProfile(t, app, db, BASEURL)
	deleteProfile(t, app, db, BASEURL)

	// Drop Tables and Close Connectiom
	DropData(db)
	CloseTestDatabase(db)
}