package schemas

import (
	"time"

	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

type LatestMessageSchema struct {
	Sender UserDataSchema `json:"sender"`
	Text   *string        `json:"text"`
	File   *string        `json:"file"`
}

type ChatSchema struct {
	Edges         *ent.ChatEdges       `json:"edges,omitempty" swaggerignore:"true"`
	ID            uuid.UUID            `json:"id" example:"d10dde64-a242-4ed0-bd75-4c759644b3a6"`
	Owner         UserDataSchema       `json:"owner"`
	Name          *string              `json:"name" example:"Correct Group"`
	Ctype         string               `json:"ctype" example:"DM"`
	Description   *string              `json:"description" example:"A nice group for tech enthusiasts"`
	LatestMessage *LatestMessageSchema `json:"latest_message"`
	Image         *string              `json:"image" example:"https://img.url"`
	CreatedAt     time.Time            `json:"created_at" example:"2024-01-14T19:00:02.613124+01:00"`
	UpdatedAt     time.Time            `json:"updated_at" example:"2024-01-14T19:00:02.613124+01:00"`
}

func (chat ChatSchema) Init() ChatSchema {
	// Set Owner Details.
	chat.Owner = chat.Owner.Init(chat.Edges.Owner)

	// Set ImageUrl
	image := chat.Edges.Image
	if image != nil {
		url := utils.GenerateFileUrl(image.ID.String(), "chats", image.ResourceType)
		chat.Image = &url
	}

	// Set Reactions & Comments Count
	latestMessages := chat.Edges.Messages
	if len(latestMessages) > 0 {
		latestMessage := latestMessages[0]
		file := latestMessage.Edges.File
		var fileUrl *string
		if file != nil {
			url := utils.GenerateFileUrl(file.ID.String(), "messages", file.ResourceType)
			fileUrl = &url
		}
		lm := LatestMessageSchema{
			Text: latestMessage.Text,
			File: fileUrl,
		}
		lm.Sender = lm.Sender.Init(latestMessage.Edges.Sender)
		chat.LatestMessage = &lm
	}
	chat.Edges = nil // Omit edges
	return chat
}

type MessageSchema struct {
	Edges     *ent.MessageEdges `json:"edges,omitempty" swaggerignore:"true"`
	ID        uuid.UUID         `json:"id"`
	ChatID    uuid.UUID         `json:"chat_id"`
	Sender    UserDataSchema    `json:"sender"`
	Text      *string           `json:"text" example:"Jesus is Lord"`
	File      *string           `json:"file" example:"https://img.url"`
	CreatedAt time.Time         `json:"created_at" example:"2024-01-14T19:00:02.613124+01:00"`
	UpdatedAt time.Time         `json:"updated_at" example:"2024-01-14T19:00:02.613124+01:00"`
}

func (message MessageSchema) Init() MessageSchema {
	// Set Author Details.
	message.Sender = message.Sender.Init(message.Edges.Sender)

	// Set FileUrl
	file := message.Edges.File
	if file != nil {
		url := utils.GenerateFileUrl(file.ID.String(), "messages", file.ResourceType)
		message.File = &url
	}

	message.Edges = nil // Omit edges
	return message
}

type GroupChatSchema struct {
	Edges       *ent.ChatEdges   `json:"edges,omitempty" swaggerignore:"true"`
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name" example:"Correct Group"`
	Description *string          `json:"description" example:"Jesus is Lord"`
	Image       *string          `json:"image" example:"https://img.url"`
	Users       []UserDataSchema `json:"users"`
}

func (chat GroupChatSchema) Init() GroupChatSchema {
	// Set Users Details.
	users := []UserDataSchema{}
	for _, user := range chat.Edges.Users {
		userData := UserDataSchema{}.Init(user)
		users = append(users, userData)
	}
	chat.Users = users

	// Set ImageUrl
	image := chat.Edges.Image
	if image != nil {
		url := utils.GenerateFileUrl(image.ID.String(), "groups", image.ResourceType)
		chat.Image = &url
	}

	chat.Edges = nil // Omit edges
	return chat
}

type MessageCreateSchema struct {
	ChatID   *uuid.UUID `json:"chat_id" validate:"omitempty" example:"d10dde64-a242-4ed0-bd75-4c759644b3a6"`
	Username *string    `json:"username,omitempty" validate:"required_without=ChatID" example:"john-doe"`
	Text     *string    `json:"text" validate:"required_without=FileType" example:"I am not in danger skyler, I am the danger"`
	FileType *string    `json:"file_type" validate:"omitempty,file_type_validator" example:"image/jpeg"`
}

type MessageUpdateSchema struct {
	Text     *string `json:"text" validate:"required_without=FileType" example:"The Earth is the Lord's and the fullness thereof"`
	FileType *string `json:"file_type" validate:"omitempty,file_type_validator" example:"image/jpeg"`
}

type MessagesResponseDataSchema struct {
	PaginatedResponseDataSchema
	Items []MessageSchema `json:"items"`
}

func (data MessagesResponseDataSchema) Init() MessagesResponseDataSchema {
	// Set Initial Data
	items := data.Items
	for i := range items {
		items[i] = items[i].Init()
	}
	data.Items = items
	return data
}

type MessagesSchema struct {
	Chat     ChatSchema                 `json:"chat"`
	Messages MessagesResponseDataSchema `json:"messages"`
	Users    []*UserDataSchema          `json:"users"`
}

func (data MessagesSchema) Init() MessagesSchema {
	// Set Initial Data
	// Set Users
	data.Users = ConvertUsers(data.Chat.Edges.Users)

	// Set Chat
	data.Chat = data.Chat.Init()
	return data
}

type GroupChatInputSchema struct {
	Name              *string   `json:"name" validate:"omitempty,max=100" example:"Dopest Group"`
	Description       *string   `json:"description" validate:"omitempty,max=1000" example:"This is a group for bosses."`
	UsernamesToAdd    *[]string `json:"usernames_to_add" validate:"omitempty,min=1,max=99" example:"['john-doe']"`
	UsernamesToRemove *[]string `json:"usernames_to_remove" validate:"omitempty,min=1,max=99" example:"['john-doe']"`
	FileType          *string   `json:"file_type" validate:"omitempty,file_type_validator" example:"image/jpeg"`
}

type GroupChatCreateSchema struct {
	GroupChatInputSchema
	Name              string    `json:"name" validate:"max=100" example:"Dopest Group"`
	UsernamesToAdd    []string  `json:"usernames_to_add" validate:"omitempty,min=1,max=99" example:"['john-doe']"`
	UsernamesToRemove *[]string `json:"-"`
}

// RESPONSE SCHEMAS
// CHATS
type ChatsResponseDataSchema struct {
	PaginatedResponseDataSchema
	Items []ChatSchema `json:"chats"`
}

func (data ChatsResponseDataSchema) Init() ChatsResponseDataSchema {
	// Set Initial Data
	items := data.Items
	for i := range items {
		items[i] = items[i].Init()
	}
	data.Items = items
	return data
}

type ChatsResponseSchema struct {
	ResponseSchema
	Data ChatsResponseDataSchema `json:"data"`
}

type MessageCreateResponseDataSchema struct {
	MessageSchema
	File           *string                `json:"file,omitempty" swaggerignore:"true"` // Remove image during create & update
	FileUploadData *utils.SignatureFormat `json:"file_upload_data"`
}

func (messageData MessageCreateResponseDataSchema) Init(fileType *string) MessageCreateResponseDataSchema {
	file := messageData.MessageSchema.Edges.File
	if fileType != nil && file != nil { // Generate data when file is being uploaded
		fuData := utils.GenerateFileSignature(file.ID.String(), "messages")
		messageData.FileUploadData = &fuData
	}
	messageData.MessageSchema = messageData.MessageSchema.Init()
	return messageData
}

type MessageCreateResponseSchema struct {
	ResponseSchema
	Data MessageCreateResponseDataSchema `json:"data"`
}

type ChatResponseSchema struct {
	ResponseSchema
	Data MessagesSchema `json:"data"`
}

type GroupChatInputResponseDataSchema struct {
	GroupChatSchema
	Image          *string                `json:"image,omitempty" swaggerignore:"true"` // Remove image during create & update
	FileUploadData *utils.SignatureFormat `json:"file_upload_data"`
}

func (groupChatData GroupChatInputResponseDataSchema) Init(fileType *string) GroupChatInputResponseDataSchema {
	image := groupChatData.GroupChatSchema.Edges.Image
	if fileType != nil && image != nil { // Generate data when file is being uploaded
		fuData := utils.GenerateFileSignature(image.ID.String(), "groups")
		groupChatData.FileUploadData = &fuData
	}
	groupChatData.GroupChatSchema = groupChatData.GroupChatSchema.Init()
	return groupChatData
}

type GroupChatInputResponseSchema struct {
	ResponseSchema
	Data 			GroupChatInputResponseDataSchema		`json:"data"`
}
