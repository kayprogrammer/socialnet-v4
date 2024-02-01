package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

var chatManager = managers.ChatManager{}

// @Summary Retrieve User Chats
// @Description `This endpoint retrieves a paginated list of the current user chats`
// @Tags Chat
// @Param page query int false "Current Page" default(1)
// @Success 200 {object} schemas.ChatsResponseSchema
// @Router /chats [get]
// @Security BearerAuth
func RetrieveUserChats(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	user := c.Locals("user").(*ent.User)
	chats := chatManager.GetUserChats(db, user)

	// Paginate, Convert type and return chats
	paginatedData, paginatedChats, err := PaginateQueryset(chats, c, 200)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	convertedChats := utils.ConvertStructData(paginatedChats, []schemas.ChatSchema{}).(*[]schemas.ChatSchema)
	response := schemas.ChatsResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Chats fetched"}.Init(),
		Data: schemas.ChatsResponseDataSchema{
			PaginatedResponseDataSchema: *paginatedData,
			Items:                       *convertedChats,
		}.Init(),
	}
	return c.Status(200).JSON(response)
}
