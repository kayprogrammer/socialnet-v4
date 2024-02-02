package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/managers"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

var (
	chatManager    = managers.ChatManager{}
	messageManager = managers.MessageManager{}
)

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

// @Summary Send a message
// @Description `This endpoint sends a message`
// @Description
// @Description `You must either send a text or a file or both.`
// @Description
// @Description `If there's no chat_id, then its a new chat and you must set username and leave chat_id`
// @Description
// @Description `If chat_id is available, then ignore username and set the correct chat_id`
// @Description
// @Description `The file_upload_data in the response is what is used for uploading the file to cloudinary from client`
// @Tags Chat
// @Param post body schemas.MessageCreateSchema true "Message object"
// @Success 201 {object} schemas.MessageCreateResponseSchema
// @Router /chats [post]
// @Security BearerAuth
func SendMessage(c *fiber.Ctx) error {
	db := c.Locals("db").(*ent.Client)
	user := c.Locals("user").(*ent.User)

	messageData := schemas.MessageCreateSchema{}

	// Validate request
	if errCode, errData := DecodeJSONBody(c, &messageData); errData != nil {
		return c.Status(errCode).JSON(errData)
	}
	if err := validator.Validate(messageData); err != nil {
		return c.Status(422).JSON(err)
	}

	chatID := messageData.ChatID
	username := messageData.Username

	var chat *ent.Chat
	if chatID == nil {
		// Create a new chat dm with current user and recipient user
		recipientUser := userManager.GetByUsername(db, *username)
		if recipientUser == nil {
			data := map[string]string{
				"username": "No user with that username",
			}
			return c.Status(422).JSON(utils.RequestErr(utils.ERR_INVALID_ENTRY, "Invalid entry", data))
		}
		chat = chatManager.GetDMChat(db, user, recipientUser)
		// Check if a chat already exists between both users
		if chat != nil {
			data := map[string]string{
				"username": "A chat already exist between you and the recipient",
			}
			return c.Status(422).JSON(utils.RequestErr(utils.ERR_INVALID_ENTRY, "Invalid entry", data))
		}
		chat = chatManager.Create(db, user, "DM", []*ent.User{recipientUser})
	} else {
		// Get the chat with chat id and check if the current user is the owner or the recipient
		chat = chatManager.GetSingleUserChat(db, user, *chatID)
		if chat == nil {
			return c.Status(404).JSON(utils.RequestErr(utils.ERR_NON_EXISTENT, "User has no chat with that ID"))
		}
	}

	//Create Message
	message := messageManager.Create(db, user, chat, messageData.Text, messageData.FileType)

	// Convert type and return Message
	convertedMessage := utils.ConvertStructData(message, schemas.MessageCreateResponseDataSchema{}).(*schemas.MessageCreateResponseDataSchema)
	response := schemas.MessageCreateResponseSchema{
		ResponseSchema: schemas.ResponseSchema{Message: "Message sent"}.Init(),
		Data:           convertedMessage.Init(messageData.FileType),
	}
	return c.Status(201).JSON(response)
}
