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
		WithAuthor(func(uq *ent.UserQuery) { uq.WithAvatar() }).
		WithImage().
		WithReactions().
		WithComments().
		Order(ent.Desc(post.FieldCreatedAt)).
		All(Ctx)
	return posts
}

func (obj PostManager) Create(client *ent.Client, author *ent.User, postData schemas.PostInputSchema) (*ent.Post, error) {
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
	// p.Edges
	if imageId != nil {
		p.Edges.Image = image
	}

	if err != nil {
		fmt.Printf("failed creating post: %v\n", err)
		return nil, nil
	}
	return p, nil
}

func (obj PostManager) GetBySlug(client *ent.Client, slug string) *ent.Post {
	p, _ := client.Post.Query().
		Where(post.Slug(slug)).
		WithAuthor(func(uq *ent.UserQuery) { uq.WithAvatar() }).
		WithImage().
		WithReactions().
		WithComments().
		Only(Ctx)
	return p
}

func (obj PostManager) Update(client *ent.Client, post *ent.Post, postData schemas.PostInputSchema) *ent.Post {
	var imageId *uuid.UUID
	image := post.Edges.Image
	fileM := FileManager{}
	if postData.FileType != nil {
		// Create or Update Image Object
		if image == nil {
			image, _ = FileManager{}.Create(client, postData.FileType)
		} else {
			image = fileM.Update(client, image, *postData.FileType)

		}
		imageId = &image.ID
	}
	p, _ := post.
		Update().
		SetText(postData.Text).
		SetNillableImageID(imageId).
		Save(Ctx)

	// Set related values
	p.Edges.Author = post.Edges.Author
	p.Edges.Comments = post.Edges.Comments
	p.Edges.Reactions = post.Edges.Reactions
	p.Edges.Image = image
	return p
}