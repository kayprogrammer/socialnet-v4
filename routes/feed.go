package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

var postManager = managers.PostManager{}

// @Summary Retrieve Latest Posts
// @Description This endpoint retrieves paginated responses of latest posts
// @Tags Feed
// @Param page query int false "Current Page" default(1)
// @Success 200 {object} schemas.PostsResponseSchema
// @Router /feed/posts [get]
func RetrievePosts(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	posts := postManager.All(db)

	// Paginate, Convert type and return Posts
	paginatedData, paginatedPosts, err := PaginateQueryset(posts, c)
	if err != nil {
		return c.Status(400).JSON(utils.ErrorResponse{Code: utils.ERR_INVALID_PAGE, Message: *err}.Init())
	}
	convertedPosts := utils.ConvertStructData(paginatedPosts, []schemas.PostSchema{}).(*[]schemas.PostSchema)
	response := schemas.PostsResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Posts fetched"}.Init(),
		Data: schemas.PostsResponseDataSchema{
			PaginatedResponseDataSchema: paginatedData,
			Items:                       *convertedPosts,
		}.Init(),
	}
	return c.Status(200).JSON(response)
}

// @Summary Create Post
// @Description This endpoint creates a new post
// @Tags Feed
// @Param post body schemas.PostInputSchema true "Post object"
// @Success 201 {object} schemas.PostInputResponseSchema
// @Router /feed/posts [post]
// @Security BearerAuth
func CreatePost(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	user := c.Locals("user").(*ent.User)

	validator := utils.Validator()

	postData := schemas.PostInputSchema{}

	// Validate request
	if errCode, errData := DecodeJSONBody(c, &postData); errData != nil {
		return c.Status(errCode).JSON(errData)
	}
	if err := validator.Validate(postData); err != nil {
		return c.Status(422).JSON(err)
	}

	post, _ := postManager.Create(db, user, postData)

	// Convert type and return Post
	convertedPost := utils.ConvertStructData(post, schemas.PostInputResponseDataSchema{}).(*schemas.PostInputResponseDataSchema)
	response := schemas.PostInputResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Post created"}.Init(),
		Data:           convertedPost.Init(postData.FileType),
	}
	return c.Status(201).JSON(response)
}

// @Summary Retrieve Single Post
// @Description This endpoint retrieves a single post
// @Tags Feed
// @Param slug path string true "Post slug"
// @Success 200 {object} schemas.PostResponseSchema
// @Router /feed/posts/{slug} [get]
func RetrievePost(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	slug := c.Params("slug")

	post := postManager.GetBySlug(db, slug)

	// Convert type and return Post
	if post == nil {
		return c.Status(404).JSON(utils.ErrorResponse{Code: utils.ERR_NON_EXISTENT, Message: "Post does not exist"}.Init())
	}
	convertedPost := utils.ConvertStructData(post, schemas.PostSchema{}).(*schemas.PostSchema)
	response := schemas.PostResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Post Detail fetched"}.Init(),
		Data:           convertedPost.Init(),
	}
	return c.Status(200).JSON(response)
}

// @Summary Update Post
// @Description This endpoint updates a post
// @Tags Feed
// @Param slug path string true "Post slug"
// @Param post body schemas.PostInputSchema true "Post object"
// @Success 200 {object} schemas.PostInputResponseSchema
// @Router /feed/posts/{slug} [put]
// @Security BearerAuth
func UpdatePost(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	user := c.Locals("user").(*ent.User)
	slug := c.Params("slug")

	validator := utils.Validator()

	postData := schemas.PostInputSchema{}

	// Validate request
	if errCode, errData := DecodeJSONBody(c, &postData); errData != nil {
		return c.Status(errCode).JSON(errData)
	}
	if err := validator.Validate(postData); err != nil {
		return c.Status(422).JSON(err)
	}
	post := postManager.GetBySlug(db, slug)

	// Validate Post existence and ownership
	if post == nil {
		return c.Status(404).JSON(utils.ErrorResponse{Code: utils.ERR_NON_EXISTENT, Message: "Post does not exist"}.Init())
	}
	if post.AuthorID != user.ID {
		return c.Status(400).JSON(utils.ErrorResponse{Code: utils.ERR_INVALID_OWNER, Message: "This Post isn't yours"}.Init())
	}

	post = postManager.Update(db, post, postData)
	// Convert type and return Post
	convertedPost := utils.ConvertStructData(post, schemas.PostInputResponseDataSchema{}).(*schemas.PostInputResponseDataSchema)
	response := schemas.PostInputResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Post updated"}.Init(),
		Data:           convertedPost.Init(postData.FileType),
	}
	return c.Status(200).JSON(response)
}

// @Summary Delete a Post
// @Description This endpoint deletes a post
// @Tags Feed
// @Param slug path string true "Post slug"
// @Success 200 {object} schemas.ResponseSchema
// @Router /feed/posts/{slug} [delete]
// @Security BearerAuth
func DeletePost(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	slug := c.Params("slug")
	user := c.Locals("user").(*ent.User)

	post := postManager.GetBySlug(db, slug)

	// Validate post existence and ownership
	if post == nil {
		return c.Status(404).JSON(utils.ErrorResponse{Code: utils.ERR_NON_EXISTENT, Message: "Post does not exist"}.Init())
	}
	if post.AuthorID != user.ID {
		return c.Status(400).JSON(utils.ErrorResponse{Code: utils.ERR_INVALID_OWNER, Message: "This Post isn't yours"}.Init())
	}

	db.Post.DeleteOne(post).Exec(managers.Ctx)
	response := schemas.ResponseSchema{Message: "Post Deleted"}.Init()
	return c.Status(200).JSON(response)
}
