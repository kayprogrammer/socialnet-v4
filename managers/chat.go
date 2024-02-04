package managers

import (
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/chat"
	"github.com/kayprogrammer/socialnet-v4/ent/message"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

// ----------------------------------
// CHAT MANAGEMENT
// --------------------------------
type ChatManager struct {
}

func (obj ChatManager) GetUserChats(client *ent.Client, userObj *ent.User) []*ent.Chat {
	chats := client.Chat.Query().
		Where(
			chat.Or(
				chat.OwnerIDEQ(userObj.ID),
				chat.HasUsersWith(user.ID(userObj.ID)),
			),
		).
		WithOwner(func(uq *ent.UserQuery) { uq.WithAvatar() }).
		WithImage().
		WithMessages(
			func(mq *ent.MessageQuery) {
				mq.WithSender(func(uq *ent.UserQuery) { uq.WithAvatar() }).WithFile().Order(ent.Desc(message.FieldCreatedAt))
			}).
		Order(ent.Desc(chat.FieldUpdatedAt)).
		AllX(Ctx)
	return chats
}

func (obj ChatManager) GetByID(client *ent.Client, id uuid.UUID) *ent.Chat {
	chatObj, _ := client.Chat.Query().
		Where(
			chat.IDEQ(id),
		).
		Only(Ctx)
	return chatObj
}

func (obj ChatManager) GetDMChat(client *ent.Client, userObj *ent.User, recipientUser *ent.User) *ent.Chat {
	chatObj, _ := client.Chat.Query().
		Where(
			chat.CtypeEQ("DM"),
			chat.Or(
				chat.And(
					chat.OwnerIDEQ(userObj.ID),
					chat.HasUsersWith(user.ID(recipientUser.ID)),
				),
				chat.And(
					chat.OwnerIDEQ(recipientUser.ID),
					chat.HasUsersWith(user.ID(userObj.ID)),
				),
			),
		).
		Only(Ctx)
	return chatObj
}

func (obj ChatManager) Create(client *ent.Client, owner *ent.User, ctype chat.Ctype, recipientsOpts ...[]*ent.User) *ent.Chat {
	chatObjCreationQuery := client.Chat.Create().
		SetCtype(ctype).
		SetOwner(owner)

	if len(recipientsOpts) > 0 {
		chatObjCreationQuery = chatObjCreationQuery.AddUsers(recipientsOpts[0]...)
	}
	chatObj := chatObjCreationQuery.SaveX(Ctx)
	return chatObj
}

func (obj ChatManager) CreateGroup(client *ent.Client, owner *ent.User, usersToAdd []*ent.User, data schemas.GroupChatCreateSchema) *ent.Chat {
	var imageId *uuid.UUID
	var image *ent.File
	if data.FileType != nil {
		image = FileManager{}.Create(client, *data.FileType)
		imageId = &image.ID
	}

	chat := client.Chat.Create().
		SetOwner(owner).
		SetName(data.Name).
		SetNillableDescription(data.Description).
		SetCtype("GROUP").
		SetNillableImageID(imageId).
		AddUsers(usersToAdd...).
		SaveX(Ctx)

	// Set related data
	chat.Edges.Users = usersToAdd
	chat.Edges.Image = image
	return chat
}

func (obj ChatManager) UsernamesToAddAndRemoveValidations(client *ent.Client, chatObj *ent.Chat, chatUpdateQuery *ent.ChatUpdateOne, usernamesToAdd *[]string, usernamesToRemove *[]string) (*ent.ChatUpdateOne, *utils.ErrorResponse) {
	originalExistingUserIDs := []uuid.UUID{}
	for _, user := range chatObj.Edges.Users {
		originalExistingUserIDs = append(originalExistingUserIDs, user.ID)
	}
	expectedUserTotal := len(originalExistingUserIDs)
	usersToAdd := []*ent.User{}
	if usernamesToAdd != nil {
		usersToAdd = client.User.Query().
			Where(
				user.UsernameIn(*usernamesToAdd...),
				user.Or(
					user.Not(user.IDIn(originalExistingUserIDs...)),
					user.IDNEQ(chatObj.OwnerID),
				),
			).AllX(Ctx)
		expectedUserTotal += len(usersToAdd)
		chatUpdateQuery = chatUpdateQuery.AddUsers(usersToAdd...)
	}
	usersToRemove := []*ent.User{}
	if usernamesToRemove != nil {
		if len(originalExistingUserIDs) < 1 {
			data := map[string]string{
				"usernames_to_remove": "No users to remove",
			}
			errData := utils.RequestErr(utils.ERR_INVALID_ENTRY, "Invalid Entry", data)
			return nil, &errData
		}
		usersToRemove = client.User.Query().
			Where(
				user.UsernameIn(*usernamesToRemove...),
				user.IDIn(originalExistingUserIDs...),
				user.IDNEQ(chatObj.OwnerID),
			).AllX(Ctx)
		expectedUserTotal -= len(usersToRemove)
		chatUpdateQuery = chatUpdateQuery.RemoveUsers(usersToRemove...)
	}
	if expectedUserTotal > 99 {
		data := map[string]string{
			"usernames_to_add": "99 users limit reached",
		}
		errData := utils.RequestErr(utils.ERR_INVALID_ENTRY, "Invalid Entry", data)
		return nil, &errData
	}
	return chatUpdateQuery, nil
}

func (obj ChatManager) UpdateGroup(client *ent.Client, chatObj *ent.Chat, data schemas.GroupChatInputSchema) (*ent.Chat, *utils.ErrorResponse) {
	chatUpdateQuery := chatObj.Update().
		SetNillableName(data.Name).
		SetNillableDescription(data.Description)

	// Handle users upload or remove
	var errData *utils.ErrorResponse
	chatUpdateQuery, errData = obj.UsernamesToAddAndRemoveValidations(client, chatObj, chatUpdateQuery, data.UsernamesToAdd, data.UsernamesToRemove)

	// Handle file upload
	var imageId *uuid.UUID
	image := chatObj.Edges.Image
	if data.FileType != nil {
		// Create or Update Image Object
		image = FileManager{}.UpdateOrCreate(client, image, *data.FileType)
		imageId = &image.ID
	}
	chatUpdateQuery = chatUpdateQuery.SetNillableImageID(imageId)
	updatedChat := chatUpdateQuery.SaveX(Ctx)

	// Set related data
	updatedChat.Edges.Users = chatObj.Edges.Users
	updatedChat.Edges.Image = chatObj.Edges.Image

	return updatedChat, errData
}

func (obj ChatManager) GetSingleUserChat(client *ent.Client, userObj *ent.User, id uuid.UUID) *ent.Chat {
	chat, _ := client.Chat.Query().
		Where(
			chat.IDEQ(id),
			chat.Or(
				chat.OwnerIDEQ(userObj.ID),
				chat.HasUsersWith(user.ID(userObj.ID)),
			),
		).
		Only(Ctx)
	return chat
}

func (obj ChatManager) GetSingleUserChatFullDetails(client *ent.Client, userObj *ent.User, id uuid.UUID) *ent.Chat {
	chat, _ := client.Chat.Query().
		Where(
			chat.IDEQ(id),
			chat.Or(
				chat.OwnerIDEQ(userObj.ID),
				chat.HasUsersWith(user.ID(userObj.ID)),
			),
		).
		WithOwner(func(uq *ent.UserQuery) { uq.WithAvatar() }).
		WithImage().
		WithMessages(
			func(mq *ent.MessageQuery) {
				mq.WithSender(func(uq *ent.UserQuery) { uq.WithAvatar() }).WithFile().Order(ent.Desc(message.FieldCreatedAt))
			},
		).
		WithUsers(func(uq *ent.UserQuery) { uq.WithAvatar() }).
		Only(Ctx)
	return chat
}

func (obj ChatManager) GetUserGroup(client *ent.Client, userObj *ent.User, id uuid.UUID, detailedOpts ...bool) *ent.Chat {
	chatQ := client.Chat.Query().
		Where(
			chat.CtypeEQ("GROUP"),
			chat.IDEQ(id),
			chat.OwnerIDEQ(userObj.ID),
		)
	if len(detailedOpts) > 0 {
		// Extra details
		chatQ = chatQ.
			WithOwner(func(uq *ent.UserQuery) { uq.WithAvatar() }).
			WithImage().
			WithUsers(func(uq *ent.UserQuery) { uq.WithAvatar() })
	}
	chatObj, _ := chatQ.Only(Ctx)
	return chatObj
}

func (obj ChatManager) GetMessagesCount(client *ent.Client, chatID uuid.UUID) int {
	messagesCount := client.Message.Query().
		Where(
			message.ChatIDEQ(chatID),
		).CountX(Ctx)

	return messagesCount
}

// ----------------------------------
// MESSAGE MANAGEMENT
// --------------------------------
type MessageManager struct {
}

func (obj MessageManager) Create(client *ent.Client, sender *ent.User, chat *ent.Chat, text *string, fileType *string) *ent.Message {
	var fileID *uuid.UUID
	var file *ent.File
	if fileType != nil {
		file = FileManager{}.Create(client, *fileType)
		fileID = &file.ID
	}

	messageObj := client.Message.Create().
		SetChat(chat).
		SetSender(sender).
		SetNillableText(text).
		SetNillableFileID(fileID).
		SaveX(Ctx)

	// Set related values
	messageObj.Edges.Sender = sender
	if fileID != nil {
		messageObj.Edges.File = file
	}

	// Update Chat
	client.Chat.UpdateOneID(messageObj.ChatID).Save(Ctx)
	return messageObj
}

func (obj MessageManager) GetUserMessage(client *ent.Client, userObj *ent.User, id uuid.UUID, detailedOpts ...bool) *ent.Message {
	messageQ := client.Message.Query().
		Where(
			message.IDEQ(id),
			message.SenderIDEQ(userObj.ID),
		)
	if len(detailedOpts) > 0 {
		// Extra details
		messageQ = messageQ.
			WithSender(func(uq *ent.UserQuery) { uq.WithAvatar() }).
			WithChat().
			WithFile()
	}
	messageObj, _ := messageQ.Only(Ctx)
	return messageObj
}

func (obj MessageManager) Update(client *ent.Client, message *ent.Message, text *string, fileType *string) *ent.Message {
	var fileId *uuid.UUID
	file := message.Edges.File
	if fileType != nil {
		// Create or Update Image Object
		file = FileManager{}.UpdateOrCreate(client, file, *fileType)
		fileId = &file.ID
	}

	messageObj := message.Update().
		SetNillableText(text).
		SetNillableFileID(fileId).
		SaveX(Ctx)

	// Set related values
	messageObj.Edges.Sender = message.Edges.Sender
	if fileId != nil {
		messageObj.Edges.File = file
	}
	return messageObj
}