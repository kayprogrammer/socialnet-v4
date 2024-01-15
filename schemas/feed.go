package schemas

import (
	"time"

	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/utils"
)

type PostSchema struct {
	Edges        		*ent.PostEdges 		`json:"edges,omitempty" swaggerignore:"true"`
	Author 				UserDataSchema 		`json:"author"`
	Text 				string 				`json:"text" example:"God Is Good"`
	Slug 				string				`json:"slug" example:"john-doe-d10dde64-a242-4ed0-bd75-4c759644b3a6"`
	ReactionsCount 		uint				`json:"reactions_count" example:"200"`
	CommentsCount 		uint				`json:"comments_count" example:"35"`
	Image 				*string				`json:"image" example:"https://img.url"`
	CreatedAt 			time.Time			`json:"created_at" example:"2024-01-14T19:00:02.613124+01:00"`
	UpdatedAt 			time.Time			`json:"updated_at" example:"2024-01-14T19:00:02.613124+01:00"`
}

func (post PostSchema) Init () PostSchema {
	// Set Author Details.
	post.Author = post.Author.Init(post.Edges.Author)

	// Set ImageUrl
	image := post.Edges.Image
	if image != nil {
		url := utils.GenerateFileUrl(image.ID.String(), "posts", image.ResourceType)
		post.Image = &url
	}

	// Set Reactions & Comments Count
	post.ReactionsCount = uint(len(post.Edges.Reactions))
	post.CommentsCount = uint(len(post.Edges.Comments))
	
	post.Edges = nil // Omit edges
	return post
}

type PostInputSchema struct {
	Text				string		`json:"text" example:"God is good"`
	FileType			*string		`json:"file_type" example:"image/jpeg"`
}

// RESPONSE SCHEMAS
type PostsResponseDataSchema struct {
	PaginatedResponseDataSchema
	Items			[]PostSchema		`json:"posts"`
}

func (data PostsResponseDataSchema) Init () PostsResponseDataSchema {
	// Set Initial Data
	items := data.Items
	for i := range items {
		items[i] = items[i].Init()
	}
	data.Items = items
	return data
}

type PostsResponseSchema struct {
	ResponseSchema
	Data			PostsResponseDataSchema		`json:"data"`
}
