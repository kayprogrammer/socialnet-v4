package schemas

import (
	"time"

	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

type LatestMessageSchema struct {
	Sender			UserDataSchema			`json:"sender"`
	Text			*string					`json:"text"`
	File			*string					`json:"file"`
}

type ChatSchema struct {
	Edges        		*ent.ChatEdges 		`json:"edges,omitempty" swaggerignore:"true"`
	ID					uuid.UUID			`json:"id"`
	Owner 				UserDataSchema 		`json:"owner"`
	Name 				*string 			`json:"name" example:"Correct Group"`
	Ctype 				string 				`json:"ctype" example:"DM"`
	Description 		*string 			`json:"description" example:"A nice group for tech enthusiasts"`
	LatestMessage 		LatestMessageSchema	`json:"latest_message"`
	Image 				*string				`json:"image" example:"https://img.url"`
	CreatedAt 			time.Time			`json:"created_at" example:"2024-01-14T19:00:02.613124+01:00"`
	UpdatedAt 			time.Time			`json:"updated_at" example:"2024-01-14T19:00:02.613124+01:00"`
}

func (chat ChatSchema) Init () ChatSchema {
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
		chat.LatestMessage = LatestMessageSchema{
			Sender: chat.LatestMessage.Sender.Init(latestMessage.Edges.Sender),
			Text: latestMessage.Text,
			File: fileUrl,
		}
	}
	chat.Edges = nil // Omit edges
	return chat
}

type MessageSchema struct {
	Edges        		*ent.MessageEdges 	`json:"edges,omitempty" swaggerignore:"true"`
	ID					uuid.UUID			`json:"id"`
	ChatID				uuid.UUID			`json:"chat_id"`
	Sender 				UserDataSchema 		`json:"sender"`
	Text 				*string 			`json:"text" example:"Jesus is Lord"`
	File 				*string				`json:"file" example:"https://img.url"`
	CreatedAt 			time.Time			`json:"created_at" example:"2024-01-14T19:00:02.613124+01:00"`
	UpdatedAt 			time.Time			`json:"updated_at" example:"2024-01-14T19:00:02.613124+01:00"`
}

type GroupChatSchema struct {
	Edges        		*ent.ChatEdges	 	`json:"edges,omitempty" swaggerignore:"true"`
	ID					uuid.UUID			`json:"id"`
	Name 				string 				`json:"name" example:"Correct Group"`
	Description 		*string 			`json:"description" example:"Jesus is Lord"`
	Image 				*string				`json:"image" example:"https://img.url"`
	Users 				[]*UserDataSchema	`json:"users"`
}

// RESPONSE SCHEMAS
// CHATS
type ChatsResponseDataSchema struct {
	PaginatedResponseDataSchema
	Items			[]ChatSchema		`json:"chats"`
}

func (data ChatsResponseDataSchema) Init () ChatsResponseDataSchema {
	// Set Initial Data
	items := data.Items
	for i := range items {
		items[i] = items[i].Init()
	}
	data.Items = items
	return data
}