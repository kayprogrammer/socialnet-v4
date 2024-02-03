package managers

import (
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/chat"
	"github.com/kayprogrammer/socialnet-v4/ent/message"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
	"github.com/kayprogrammer/socialnet-v4/schemas"
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

func (obj ChatManager) UsernamesToAddAndRemoveValidations(chatObj *ent.Chat, usernamesToAdd *[]string, usernamesToRemove *[]string ) *ent.Chat {
	originalExistingUserIDs := chat.users
}
func (obj ChatManager) UpdateGroup(client *ent.Client, owner *ent.User, chatObj ent.Chat, data schemas.GroupChatInputSchema) *ent.Chat {
	usernamesToAdd := data.UsernamesToAdd
	usernamesToRemove := data.UsernamesToRemove
	// Add users to the chat if they don't exist in it already

	c := chatObj.Update().
		SetNillableName(data.Name).
		SetNillableDescription(data.Description).
		setdes
	chatObjCreationQuery := client.Chat.Create().
		SetCtype(ctype).
		SetOwner(owner)

	if len(recipientsOpts) > 0 {
		chatObjCreationQuery = chatObjCreationQuery.AddUsers(recipientsOpts[0]...)
	}
	chatObj := chatObjCreationQuery.SaveX(Ctx)
	return chatObj
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

// ----------------------------------
// MESSAGE MANAGEMENT
// --------------------------------
type MessageManager struct {
}

func (obj MessageManager) Create(client *ent.Client, sender *ent.User, chat *ent.Chat, text *string, fileType *string) *ent.Message {
	var fileID *uuid.UUID
	var file *ent.File
	if fileType != nil {
		file, _ = FileManager{}.Create(client, fileType)
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
