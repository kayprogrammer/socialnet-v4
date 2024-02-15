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
)

func getPosts(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	post := CreatePost(db)
	user := post.Edges.Author
	t.Run("Retrieve Posts", func(t *testing.T) {
		url := fmt.Sprintf("%s/posts", baseUrl)
		req := httptest.NewRequest("GET", url, nil)
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
				"per_page":     50,
				"current_page": 1,
				"last_page":    1,
				"posts": []map[string]interface{}{
					{
						"author":          GetUserMap(user),
						"text":            post.Text,
						"slug":            post.Slug,
						"reactions_count": 0,
						"comments_count":  0,
						"image":           nil,
						"created_at":      ConvertDateTime(post.CreatedAt),
						"updated_at":      ConvertDateTime(post.UpdatedAt),
					},
				},
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func createPost(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	sender := CreateTestVerifiedUser(db)
	token := AccessToken(db)
	t.Run("Create Post", func(t *testing.T) {
		url := fmt.Sprintf("%s/posts", baseUrl)
		postData := schemas.PostInputSchema{Text: "My new Post"}

		res := ProcessTestBody(t, app, url, "POST", postData, token)
		// Assert Status code
		assert.Equal(t, 201, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		dataRep := body["data"].(map[string]interface{})
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Post created",
			"data": map[string]interface{}{
				"author":           GetUserMap(sender),
				"text":             postData.Text,
				"slug":             dataRep["slug"],
				"reactions_count":  0,
				"comments_count":   0,
				"created_at":       dataRep["created_at"],
				"updated_at":       dataRep["updated_at"],
				"file_upload_data": nil,
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func getPost(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	post := CreatePost(db)
	user := post.Edges.Author
	t.Run("Retrieve Post", func(t *testing.T) {
		// Test for post with invalid slug
		url := fmt.Sprintf("%s/posts/invalid_slug", baseUrl)
		req := httptest.NewRequest("GET", url, nil)
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "Post does not exist", body["message"])

		// Test for post with valid slug
		url = fmt.Sprintf("%s/posts/%s", baseUrl, post.Slug)
		req = httptest.NewRequest("GET", url, nil)
		res, _ = app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Post Detail fetched",
			"data": map[string]interface{}{
				"author":          GetUserMap(user),
				"text":            post.Text,
				"slug":            post.Slug,
				"reactions_count": 0,
				"comments_count":  0,
				"image":           nil,
				"created_at":      ConvertDateTime(post.CreatedAt),
				"updated_at":      ConvertDateTime(post.UpdatedAt),
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func updatePost(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	post := CreatePost(db)
	user := post.Edges.Author
	token := AccessToken(db)
	t.Run("Update Post", func(t *testing.T) {
		postData := schemas.PostInputSchema{Text: "Post Text Updated"}

		// Check if endpoint fails for invalid post
		url := fmt.Sprintf("%s/posts/invalid_slug", baseUrl)

		res := ProcessTestBody(t, app, url, "PUT", postData, token)
		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "Post does not exist", body["message"])

		// Check if endpoint fails for invalid owner
		url = fmt.Sprintf("%s/posts/%s", baseUrl, post.Slug)
		res = ProcessTestBody(t, app, url, "PUT", postData, AnotherAccessToken(db))
		// Assert Status code
		assert.Equal(t, 400, res.StatusCode)
		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_INVALID_OWNER, body["code"])
		assert.Equal(t, "This Post isn't yours", body["message"])

		// Check if endpoint succeeds if all requirements are met
		res = ProcessTestBody(t, app, url, "PUT", postData, token)
		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)
		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		dataRep := body["data"].(map[string]interface{})
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Post updated",
			"data": map[string]interface{}{
				"author":           GetUserMap(user),
				"text":             postData.Text,
				"slug":             dataRep["slug"],
				"reactions_count":  0,
				"comments_count":   0,
				"created_at":       dataRep["created_at"],
				"updated_at":       dataRep["updated_at"],
				"file_upload_data": nil,
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func deletePost(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	post := CreatePost(db)
	token := AccessToken(db)
	t.Run("Delete A Post", func(t *testing.T) {
		url := fmt.Sprintf("%s/posts/invalid_slug", baseUrl)
		// Test for valid response for invalid post id
		req := httptest.NewRequest("DELETE", url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "Post does not exist", body["message"])

		// Test for valid response for valid entry
		url = fmt.Sprintf("%s/posts/%s", baseUrl, post.Slug)
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
			"message": "Post Deleted",
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func getReactions(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	reaction := CreateReaction(db)
	user := reaction.Edges.User
	post := reaction.Edges.Post
	t.Run("Retrieve Reactions", func(t *testing.T) {
		// Test for invalid focus_value
		url := fmt.Sprintf("%s/reactions/invalid_focus/%s", baseUrl, post.Slug)
		req := httptest.NewRequest("GET", url, nil)
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_INVALID_VALUE, body["code"])
		assert.Equal(t, "Invalid 'focus' value", body["message"])

		// Test for invalid slug
		url = fmt.Sprintf("%s/reactions/POST/invalid_slug", baseUrl)
		req = httptest.NewRequest("GET", url, nil)
		res, _ = app.Test(req)

		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "Post does not exist", body["message"])

		// Test for valid values
		url = fmt.Sprintf("%s/reactions/POST/%s", baseUrl, post.Slug)
		req = httptest.NewRequest("GET", url, nil)
		res, _ = app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Reactions fetched",
			"data": map[string]interface{}{
				"per_page":     50,
				"current_page": 1,
				"last_page":    1,
				"reactions": []map[string]interface{}{
					{
						"id":    reaction.ID,
						"user":  GetUserMap(user),
						"rtype": reaction.Rtype,
					},
				},
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func createReaction(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	post := CreatePost(db)
	token := AccessToken(db)
	t.Run("Create Reaction", func(t *testing.T) {
		url := fmt.Sprintf("%s/reactions/POST/%s", baseUrl, post.Slug)
		reactionData := schemas.ReactionInputSchema{Rtype: "LIKE"}

		res := ProcessTestBody(t, app, url, "POST", reactionData, token)
		// Assert Status code
		assert.Equal(t, 201, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Reaction created",
			"data": map[string]interface{}{
				"id":    body["data"].(map[string]interface{})["id"],
				"user":  GetUserMap(post.Edges.Author),
				"rtype": reactionData.Rtype,
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func deleteReaction(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	reaction := CreateReaction(db)
	token := AccessToken(db)
	t.Run("Delete A Reaction", func(t *testing.T) {
		url := fmt.Sprintf("%s/reactions/%s", baseUrl, uuid.New())
		// Test for valid response for invalid reaction id
		req := httptest.NewRequest("DELETE", url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "Reaction does not exist", body["message"])

		// Test for valid response for valid entry
		url = fmt.Sprintf("%s/reactions/%s", baseUrl, reaction.ID)
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
			"message": "Reaction Deleted",
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func getComments(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	comment := CreateComment(db)
	user := comment.Edges.Author
	post := comment.Edges.Post
	t.Run("Retrieve Comments", func(t *testing.T) {
		// Test for invalid slug
		url := fmt.Sprintf("%s/posts/invalid_slug/comments", baseUrl)
		req := httptest.NewRequest("GET", url, nil)
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "Post does not exist", body["message"])

		// Test for valid values
		url = fmt.Sprintf("%s/posts/%s/comments", baseUrl, post.Slug)
		req = httptest.NewRequest("GET", url, nil)
		res, _ = app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Comments fetched",
			"data": map[string]interface{}{
				"per_page":     50,
				"current_page": 1,
				"last_page":    1,
				"comments": []map[string]interface{}{
					{
						"author":          GetUserMap(user),
						"slug":            comment.Slug,
						"text":            comment.Text,
						"reactions_count": 0,
						"replies_count":   0,
					},
				},
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func createComment(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	post := CreatePost(db)
	user := post.Edges.Author
	token := AccessToken(db)
	t.Run("Create Comment", func(t *testing.T) {
		url := fmt.Sprintf("%s/posts/%s/comments", baseUrl, post.Slug)
		commentData := schemas.CommentInputSchema{Text: "My new comment"}

		res := ProcessTestBody(t, app, url, "POST", commentData, token)
		// Assert Status code
		assert.Equal(t, 201, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Comment created",
			"data": map[string]interface{}{
				"author":          GetUserMap(user),
				"slug":            body["data"].(map[string]interface{})["slug"],
				"text":            commentData.Text,
				"reactions_count": 0,
				"replies_count":   0,
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func getCommentWithReplies(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	reply := CreateReply(db)
	comment := reply.Edges.Comment
	user := GetUserMap(reply.Edges.Author)
	t.Run("Retrieve Comment With Replies", func(t *testing.T) {
		// Test for comment slug
		url := fmt.Sprintf("%s/comments/invalid_slug", baseUrl)
		req := httptest.NewRequest("GET", url, nil)
		res, _ := app.Test(req)

		// Assert Status code
		assert.Equal(t, 404, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		assert.Equal(t, "failure", body["status"])
		assert.Equal(t, utils.ERR_NON_EXISTENT, body["code"])
		assert.Equal(t, "Comment does not exist", body["message"])

		// Test for valid values
		url = fmt.Sprintf("%s/comments/%s", baseUrl, comment.Slug)
		req = httptest.NewRequest("GET", url, nil)
		res, _ = app.Test(req)

		// Assert Status code
		assert.Equal(t, 200, res.StatusCode)

		// Parse and assert body
		body = ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Comment with replies fetched",
			"data": map[string]interface{}{
				"comment": map[string]interface{}{
					"author":          user,
					"slug":            comment.Slug,
					"text":            comment.Text,
					"reactions_count": 0,
					"replies_count":   1,
				},
				"replies": map[string]interface{}{
					"per_page":     50,
					"current_page": 1,
					"last_page":    1,
					"items": []map[string]interface{}{
						{
							"author":          user,
							"slug":            reply.Slug,
							"text":            reply.Text,
							"reactions_count": 0,
						},
					},
				},
			},
		}
		expectedDataJson, _ := json.Marshal(expectedData)
		assert.JSONEq(t, string(expectedDataJson), string(data))
	})
}

func createReply(t *testing.T, app *fiber.App, db *ent.Client, baseUrl string) {
	comment := CreateComment(db)
	user := comment.Edges.Author
	token := AccessToken(db)
	t.Run("Create Reply", func(t *testing.T) {
		url := fmt.Sprintf("%s/comments/%s", baseUrl, comment.Slug)
		replyData := schemas.CommentInputSchema{Text: "New Cool reply"}

		res := ProcessTestBody(t, app, url, "POST", replyData, token)
		// Assert Status code
		assert.Equal(t, 201, res.StatusCode)
		// Parse and assert body
		body := ParseResponseBody(t, res.Body).(map[string]interface{})
		data, _ := json.Marshal(body)
		expectedData := map[string]interface{}{
			"status":  "success",
			"message": "Reply created",
			"data": map[string]interface{}{
				"author":          GetUserMap(user),
				"slug":            body["data"].(map[string]interface{})["slug"],
				"text":            replyData.Text,
				"reactions_count": 0,
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
	createPost(t, app, db, BASEURL)
	getPost(t, app, db, BASEURL)
	updatePost(t, app, db, BASEURL)
	deletePost(t, app, db, BASEURL)
	getReactions(t, app, db, BASEURL)
	createReaction(t, app, db, BASEURL)
	deleteReaction(t, app, db, BASEURL)
	getComments(t, app, db, BASEURL)
	createComment(t, app, db, BASEURL)
	getCommentWithReplies(t, app, db, BASEURL)
	createReply(t, app, db, BASEURL)

	// Drop Tables and Close Connectiom
	DropData(db)
	CloseTestDatabase(db)
}
