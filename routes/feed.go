package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

var postManager = managers.PostManager{}
var validator = utils.Validator()

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
		return c.Status(400).JSON(err)
	}
	convertedPosts := utils.ConvertStructData(paginatedPosts, []schemas.PostSchema{}).(*[]schemas.PostSchema)
	response := schemas.PostsResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Posts fetched"}.Init(),
		Data: schemas.PostsResponseDataSchema{
			PaginatedResponseDataSchema: *paginatedData,
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

	// Retrieve, Convert type and return Post
	post, errCode, errData := postManager.GetBySlug(db, slug, true)
	if errCode != nil {
		return c.Status(*errCode).JSON(errData)
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

	postData := schemas.PostInputSchema{}

	// Validate request
	if errCode, errData := DecodeJSONBody(c, &postData); errData != nil {
		return c.Status(errCode).JSON(errData)
	}
	if err := validator.Validate(postData); err != nil {
		return c.Status(422).JSON(err)
	}

	// Retrieve & Validate Post Existence
	post, errCode, errData := postManager.GetBySlug(db, slug, true)
	if errCode != nil {
		return c.Status(*errCode).JSON(errData)
	}

	// Validate Post ownership
	if post.AuthorID != user.ID {
		return c.Status(400).JSON(utils.RequestErr(utils.ERR_INVALID_OWNER, "This Post isn't yours"))
	}

	// Update, Convert type and return Post
	post = postManager.Update(db, post, postData)
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

	// Retrieve & Validate Post Existence
	post, errCode, errData := postManager.GetBySlug(db, slug)
	if errCode != nil {
		return c.Status(*errCode).JSON(errData)
	}

	// Validate Post ownership
	if post.AuthorID != user.ID {
		return c.Status(400).JSON(utils.RequestErr(utils.ERR_INVALID_OWNER, "This Post isn't yours"))
	}

	// Delete and return response
	db.Post.DeleteOne(post).Exec(managers.Ctx)
	response := schemas.ResponseSchema{Message: "Post Deleted"}.Init()
	return c.Status(200).JSON(response)
}

var reactionManager = managers.ReactionManager{}

// @Summary Retrieve Latest Reactions of a Post, Comment, or Reply
// @Description This endpoint retrieves paginated responses of reactions of post, comment, reply
// @Tags Feed
// @Param focus path string true "Specify the usage. Use any of the three: POST, COMMENT, REPLY"
// @Param slug path string true "Enter the slug of the post or comment or reply"
// @Param page query int false "Current Page" default(1)
// @Param reaction_type query string false "Reaction Type. Must be any of these: LIKE, LOVE, HAHA, WOW, SAD, ANGRY"
// @Success 200 {object} schemas.ReactionsResponseSchema
// @Router /feed/reactions/{focus}/{slug} [get]
func RetrieveReactions(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	focus := c.Params("focus")
	slug := c.Params("slug")

	// Validate Focus
	err := ValidateReactionFocus(focus)
	if err != nil {
		return c.Status(404).JSON(err)
	}

	// Paginate, Convert type and return Posts
	reactions, errCode, errData := reactionManager.GetReactionsQueryset(db, c, focus, slug)
	if errCode != nil {
		return c.Status(*errCode).JSON(errData)
	}
	// Paginate, Convert type and return Reactions
	paginatedData, paginatedReactions, err := PaginateQueryset(reactions, c)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	convertedReactions := utils.ConvertStructData(paginatedReactions, []schemas.ReactionSchema{}).(*[]schemas.ReactionSchema)
	response := schemas.ReactionsResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Reactions fetched"}.Init(),
		Data: schemas.ReactionsResponseDataSchema{
			PaginatedResponseDataSchema: *paginatedData,
			Items:                       *convertedReactions,
		}.Init(),
	}
	return c.Status(200).JSON(response)
}

// @Summary Create Reaction
// @Description This endpoint creates a new reaction.
// @Tags Feed
// @Param focus path string true "Specify the usage. Use any of the three: POST, COMMENT, REPLY"
// @Param slug path string true "Enter the slug of the post or comment or reply"
// @Param post body schemas.ReactionInputSchema true "Reaction object. rtype should be any of these: LIKE, LOVE, HAHA, WOW, SAD, ANGRY"
// @Success 201 {object} schemas.ReactionResponseSchema
// @Router /feed/reactions/{focus}/{slug} [post]
// @Security BearerAuth
func CreateReaction(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	focus := c.Params("focus")
	slug := c.Params("slug")
	user := c.Locals("user").(*ent.User)

	// Validate Focus
	err := ValidateReactionFocus(focus)
	if err != nil {
		return c.Status(404).JSON(err)
	}

	reactionData := schemas.ReactionInputSchema{}

	// Validate request
	if errCode, errData := DecodeJSONBody(c, &reactionData); errData != nil {
		return c.Status(errCode).JSON(errData)
	}
	if err := validator.Validate(reactionData); err != nil {
		return c.Status(422).JSON(err)
	}

	// Update Or Create Reaction
	reaction, errCode, errData := reactionManager.UpdateOrCreate(db, user, focus, slug, reactionData.Rtype)
	if errCode != nil {
		return c.Status(*errCode).JSON(errData)
	}

	// Convert type and return Reactions
	convertedReaction := utils.ConvertStructData(reaction, schemas.ReactionSchema{}).(*schemas.ReactionSchema)
	response := schemas.ReactionResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Reaction created"}.Init(),
		Data:           convertedReaction.Init(),
	}

	// Send Notifications here later
	return c.Status(201).JSON(response)
}

// @Summary Remove Reaction
// @Description This endpoint deletes a reaction
// @Tags Feed
// @Param id path string true "Reaction id (uuid)"
// @Success 200 {object} schemas.ResponseSchema
// @Router /feed/reactions/{id} [delete]
// @Security BearerAuth
func DeleteReaction(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	id := c.Params("id")
	// Parse the UUID parameter
	reactionID, err := utils.ParseUUID(id)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	user := c.Locals("user").(*ent.User)

	// Retrieve & Validate Reaction Existence & Ownership
	reaction, errCode, errData := reactionManager.GetByID(db, *reactionID)
	if errCode != nil {
		return c.Status(*errCode).JSON(errData)
	}

	// Validate Reaction ownership
	if reaction.UserID != user.ID {
		return c.Status(400).JSON(utils.RequestErr(utils.ERR_INVALID_OWNER, "This Reaction isn't yours"))
	}

	// Remove Reaction Notifications here later

	// Delete and return response
	db.Reaction.DeleteOne(reaction).Exec(managers.Ctx)
	response := schemas.ResponseSchema{Message: "Reaction Deleted"}.Init()
	return c.Status(200).JSON(response)
}

var commentManager = managers.CommentManager{}

// @Summary Retrieve Post Comments
// @Description This endpoint retrieves comments of a particular post
// @Tags Feed
// @Param slug path string true "Post Slug"
// @Param page query int false "Current Page" default(1)
// @Success 200 {object} schemas.CommentsResponseSchema
// @Router /feed/posts/{slug}/comments [get]
func RetrieveComments(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	slug := c.Params("slug")

	// Get Post
	post, errCode, errData := postManager.GetBySlug(db, slug)
	if errCode != nil {
		return c.Status(*errCode).JSON(errData)
	}

	// Get Comments
	comments := commentManager.GetByPostID(db, post.ID)

	// Paginate, Convert type and return comments
	paginatedData, paginatedComments, err := PaginateQueryset(comments, c)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	convertedComments := utils.ConvertStructData(paginatedComments, []schemas.CommentSchema{}).(*[]schemas.CommentSchema)
	response := schemas.CommentsResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Comments fetched"}.Init(),
		Data: schemas.CommentsResponseDataSchema{
			PaginatedResponseDataSchema: *paginatedData,
			Items:                       *convertedComments,
		}.Init(),
	}
	return c.Status(200).JSON(response)
}
