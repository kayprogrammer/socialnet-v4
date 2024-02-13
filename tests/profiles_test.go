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
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
	"github.com/stretchr/testify/assert"
)

var (
	notificationManager = managers.NotificationManager{}
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

func getFriends(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	t.Run("Retrieve Friends", func(t *testing.T) {
		friend := CreateFriend(db, "ACCEPTED")
		requestee := friend.Edges.Requestee

		// Test for valid response
		url := fmt.Sprintf("%s/friends", baseUrl)
		req := httptest.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken(db)))
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "success", body["status"])
		assert.Equal(t, "Friends fetched", body["message"])

		data, _ := json.Marshal(body["data"])
		expectedData := map[string]interface{}{
			"per_page":     20,
			"current_page": 1,
			"last_page":    1,
			"users": []map[string]interface{}{
				{
					"first_name": requestee.FirstName,
					"last_name":  requestee.LastName,
					"username":   requestee.Username,
					"email":      requestee.Email,
					"bio":        requestee.Bio,
					"avatar":     nil,
					"dob":        requestee.Dob,
					"city":       nil,
					"created_at": ConvertDateTime(requestee.CreatedAt),
					"updated_at": ConvertDateTime(requestee.UpdatedAt),
				},
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.Equal(t, expectedDataJson, data)
	})
}

func sendFriendRequest(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	// Drop Friends data
	friendManager.DropData(db)
	user := CreateAnotherTestVerifiedUser(db)
	t.Run("Send Friend Request", func(t *testing.T) {
		url := fmt.Sprintf("%s/friends/requests", baseUrl)
		userData := schemas.SendFriendRequestSchema{
			Username: "invalid_username",
		}
		// Test for valid response for non-existent user name
		res := ProcessTestBody(t, app, url, "POST", userData, AccessToken(db))
		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "User does not exist!", body["message"])

		// Test for valid response for valid entry
		userData.Username = user.Username
		res = ProcessTestBody(t, app, url, "POST", userData, AccessToken(db))
		// Assert Status code
		assert.Equal(t, 201, res.StatusCode)
		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "success", body["status"])
		assert.Equal(t, "Friend Request sent", body["message"])
	})
}

func acceptOrRejectFriendRequest(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	// Drop & Create Friends data
	friendManager.DropData(db)
	friend := CreateFriend(db, "PENDING")
	t.Run("Accept Or Reject Friend Request", func(t *testing.T) {
		url := fmt.Sprintf("%s/friends/requests", baseUrl)
		userData := schemas.AcceptFriendRequestSchema{
			Username: "invalid_username",
			Accepted: true,
		}
		// Test for valid response for non-existent user name
		res := ProcessTestBody(t, app, url, "PUT", userData, AnotherAccessToken(db))
		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "User does not exist!", body["message"])

		// Test for valid response for valid entry
		userData.Username = friend.Edges.Requester.Username
		res = ProcessTestBody(t, app, url, "PUT", userData, AnotherAccessToken(db))
		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)
		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "success", body["status"])
		assert.Equal(t, "Friend Request Accepted", body["message"])
	})
}

func getNotifications(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	user := CreateTestVerifiedUser(db)
	text := "A new update is coming!"
	notification := notificationManager.Create(db, nil, "ADMIN", []uuid.UUID{user.ID}, nil, nil, nil, &text)
	t.Run("Retrieve Notifications", func(t *testing.T) {
		// Test for valid response
		url := fmt.Sprintf("%s/notifications", baseUrl)
		req := httptest.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken(db)))
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Notifications fetched",
			"data": map[string]interface{}{
				"per_page":     50,
				"current_page": 1,
				"last_page":    1,
				"notifications": []map[string]interface{}{
					{
						"id":           notification.ID,
						"sender":       nil,
						"ntype":        notification.Ntype,
						"message":      notification.Text,
						"post_slug":    nil,
						"comment_slug": nil,
						"reply_slug":   nil,
						"is_read":      false,
					},
				},
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.Equal(t, expectedDataJson, data)
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
	getFriends(t, app, db, BASEURL)
	sendFriendRequest(t, app, db, BASEURL)
	acceptOrRejectFriendRequest(t, app, db, BASEURL)
	getNotifications(t, app, db, BASEURL)

	// Drop Tables and Close Connectiom
	DropData(db)
	CloseTestDatabase(db)
}
