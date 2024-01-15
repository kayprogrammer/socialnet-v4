package managers

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/post"
	"github.com/kayprogrammer/socialnet-v4/schemas"
)

// ----------------------------------
// POST MANAGEMENT
// --------------------------------
type PostManager struct {

}

func (obj PostManager) All(client *ent.Client) []*ent.Post {
	posts, _ := client.Post.Query().
		WithAuthor(func(uq *ent.UserQuery) {uq.WithAvatar()}).
		WithImage().
		WithReactions().
		WithComments().
		Order(ent.Desc(post.FieldCreatedAt)).
		All(Ctx)
	return posts
}

func (obj PostManager) Create(client *ent.Client, author * ent.User, postData schemas.PostInputSchema) (*ent.Post, error) {
	id := uuid.New()
	slug := slug.Make(author.FirstName + "-" + author.LastName + "-" + id.String())

	var imageId *uuid.UUID
	var image *ent.File
	if postData.FileType != nil {
		image, _ = FileManager{}.Create(client, postData.FileType)
		imageId = &image.ID
	}
	p, err := client.Post.
		Create().
		SetID(id).
		SetAuthor(author).
		SetSlug(slug). 
		SetText(postData.Text). 
		SetNillableImageID(imageId).
		Save(Ctx)
	
	// Set related values
	p.Edges.Author = author
	if imageId != nil {
		p.Edges.Image = image
	}

	if err != nil {
		fmt.Printf("failed creating post: %v\n", err)
		return nil, nil
	}
	return p, nil
}