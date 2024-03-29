package managers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/kayprogrammer/socialnet-v4/ent/comment"
	"github.com/kayprogrammer/socialnet-v4/ent/post"
	"github.com/kayprogrammer/socialnet-v4/ent/reaction"
	"github.com/kayprogrammer/socialnet-v4/ent/reply"
	"github.com/kayprogrammer/socialnet-v4/schemas"
	"github.com/kayprogrammer/socialnet-v4/utils"
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

func (obj PostManager) Create(client *ent.Client, author *ent.User, postData schemas.PostInputSchema) *ent.Post {
	id := uuid.New()
	slug := slug.Make(author.FirstName + "-" + author.LastName + "-" + id.String())

	var imageId *uuid.UUID
	var image *ent.File
	if postData.FileType != nil {
		image = FileManager{}.Create(client, *postData.FileType)
		imageId = &image.ID
	}
	p := client.Post.
		Create().
		SetID(id).
		SetAuthor(author).
		SetSlug(slug).
		SetText(postData.Text).
		SetNillableImageID(imageId).
		SaveX(Ctx)

	// Set related values
	p.Edges.Author = author
	if imageId != nil {
		p.Edges.Image = image
	}
	return p
}

func (obj PostManager) GetBySlug(client *ent.Client, slug string, opts ...bool) (*ent.Post, *int, *utils.ErrorResponse) {
	q := client.Post.Query().Where(post.Slug(slug))
	if len(opts) > 0 { // Detailed param provided.
		q = q.WithAuthor(func(uq *ent.UserQuery) { uq.WithAvatar() }).
			WithImage().
			WithReactions().
			WithComments()
	}
	post, _ := q.Only(Ctx)
	if post == nil {
		status_code := 404
		errData := utils.RequestErr(utils.ERR_NON_EXISTENT, "Post does not exist")
		return nil, &status_code, &errData
	}
	return post, nil, nil
}

func (obj PostManager) Update(client *ent.Client, post *ent.Post, postData schemas.PostInputSchema) *ent.Post {
	var imageId *uuid.UUID
	image := post.Edges.Image
	if postData.FileType != nil {
		// Create or Update Image Object
		image = FileManager{}.UpdateOrCreate(client, image, *postData.FileType)
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
func (obj PostManager) DropData(client *ent.Client) {
	client.Post.Delete().ExecX(Ctx)
}

// ----------------------------------
// COMMENT MANAGEMENT
// --------------------------------
type CommentManager struct {
}

func (obj CommentManager) GetBySlug(client *ent.Client, slug string, opts ...bool) (*ent.Comment, *int, *utils.ErrorResponse) {
	q := client.Comment.Query().Where(comment.Slug(slug))
	if len(opts) > 0 { // Detailed param provided.
		q = q.WithAuthor(func(uq *ent.UserQuery) { uq.WithAvatar() }).
			WithReactions().
			WithReplies(func(rq *ent.ReplyQuery) { rq.WithAuthor(func(uq *ent.UserQuery) { uq.WithAvatar() }) })
	}
	comment, _ := q.Only(Ctx)
	if comment == nil {
		status_code := 404
		errData := utils.RequestErr(utils.ERR_NON_EXISTENT, "Comment does not exist")
		return nil, &status_code, &errData
	}
	return comment, nil, nil
}

func (obj CommentManager) GetByPostID(client *ent.Client, postID uuid.UUID) []*ent.Comment {
	comments, _ := client.Comment.Query().
		Where(comment.PostID(postID)).
		WithAuthor(func(uq *ent.UserQuery) { uq.WithAvatar() }).
		WithReactions().
		WithReplies().
		All(Ctx)
	return comments
}

func (obj CommentManager) Create(client *ent.Client, author *ent.User, postID uuid.UUID, text string) *ent.Comment {
	id := uuid.New()
	slug := slug.Make(author.FirstName + "-" + author.LastName + "-" + id.String())
	comment, _ := client.Comment.Create().
		SetID(id).
		SetSlug(slug).
		SetAuthorID(author.ID).
		SetPostID(postID).
		SetText(text).
		Save(Ctx)

	// Set important edges
	comment.Edges.Author = author
	return comment
}

func (obj CommentManager) Update(comment *ent.Comment, author *ent.User, text string) *ent.Comment {
	c, _ := comment.Update().
		SetText(text).
		Save(Ctx)

	// Set important edges
	c.Edges.Author = comment.Edges.Author
	c.Edges.Reactions = comment.Edges.Reactions
	c.Edges.Replies = comment.Edges.Replies
	return c
}
func (obj CommentManager) DropData(client *ent.Client) {
	client.Comment.Delete().ExecX(Ctx)
}

// ----------------------------------
// REPLY MANAGEMENT
// --------------------------------
type ReplyManager struct {
}

func (obj ReplyManager) GetBySlug(client *ent.Client, slug string, opts ...bool) (*ent.Reply, *int, *utils.ErrorResponse) {
	q := client.Reply.Query().Where(reply.Slug(slug))
	if len(opts) > 0 { // Detailed param provided.
		q = q.WithAuthor(func(uq *ent.UserQuery) { uq.WithAvatar() }).
			WithReactions()
	}
	reply, _ := q.Only(Ctx)
	if reply == nil {
		status_code := 404
		errData := utils.RequestErr(utils.ERR_NON_EXISTENT, "Reply does not exist")
		return nil, &status_code, &errData
	}
	return reply, nil, nil
}

func (obj ReplyManager) Create(client *ent.Client, author *ent.User, commentID uuid.UUID, text string) *ent.Reply {
	id := uuid.New()
	slug := slug.Make(author.FirstName + "-" + author.LastName + "-" + id.String())
	reply, _ := client.Reply.Create().
		SetID(id).
		SetSlug(slug).
		SetAuthorID(author.ID).
		SetCommentID(commentID).
		SetText(text).
		Save(Ctx)

	// Set important edges
	reply.Edges.Author = author
	return reply
}

func (obj ReplyManager) Update(reply *ent.Reply, author *ent.User, text string) *ent.Reply {
	r, _ := reply.Update().
		SetText(text).
		Save(Ctx)

	// Set important edges
	r.Edges.Author = reply.Edges.Author
	r.Edges.Reactions = reply.Edges.Reactions
	return r
}

func (obj ReplyManager) DropData(client *ent.Client) {
	client.Reply.Delete().ExecX(Ctx)
}
// ----------------------------------
// REACTIONS MANAGEMENT
// --------------------------------
type ReactionManager struct {
}

func (obj ReactionManager) GetReactionsQueryset(client *ent.Client, fiberCtx *fiber.Ctx, focus string, slug string) ([]*ent.Reaction, *int, *utils.ErrorResponse) {
	q := client.Reaction.Query()
	if focus == "POST" {
		// Get Post Object and Query reactions for the post
		post, errCode, errData := PostManager{}.GetBySlug(client, slug)
		if errCode != nil {
			return nil, errCode, errData
		}
		q = q.Where(reaction.PostID(post.ID))
	} else if focus == "COMMENT" {
		// Get Comment Object and Query reactions for the comment
		comment, errCode, errData := CommentManager{}.GetBySlug(client, slug)
		if errCode != nil {
			return nil, errCode, errData
		}
		q = q.Where(reaction.CommentID(comment.ID))
	} else {
		// Get Reply Object and Query reactions for the reply
		reply, errCode, errData := ReplyManager{}.GetBySlug(client, slug)
		if errCode != nil {
			return nil, errCode, errData
		}
		q = q.Where(reaction.ReplyID(reply.ID))
	}

	// Filter by Reaction type if provided (e.g LIKE, LOVE)
	rtype := reaction.Rtype(fiberCtx.Query("reaction_type"))
	if len(rtype) > 0 {
		q = q.Where(reaction.RtypeEQ(rtype))
	}
	reactions, _ := q.WithUser(func(uq *ent.UserQuery) { uq.WithAvatar() }).All(Ctx)
	return reactions, nil, nil
}

func (obj ReactionManager) Update(reaction *ent.Reaction, focus string, id uuid.UUID, rtype reaction.Rtype) *ent.Reaction {
	r := reaction.Update().SetRtype(rtype)
	if focus == "POST" {
		r = r.SetPostID(id)
	} else if focus == "COMMENT" {
		r = r.SetCommentID(id)
	} else {
		r = r.SetReplyID(id)
	}
	reaction, _ = r.Save(Ctx)
	return reaction
}

func (obj ReactionManager) Create(client *ent.Client, userID uuid.UUID, focus string, focusID uuid.UUID, rtype reaction.Rtype) *ent.Reaction {
	r := client.Reaction.Create().SetUserID(userID).SetRtype(rtype)
	if focus == "POST" {
		r = r.SetPostID(focusID)
	} else if focus == "COMMENT" {
		r = r.SetCommentID(focusID)
	} else {
		r = r.SetReplyID(focusID)
	}
	reaction, _ := r.Save(Ctx)
	return reaction
}

func (obj ReactionManager) UpdateOrCreate(client *ent.Client, user *ent.User, focus string, slug string, rtype reaction.Rtype) (*ent.Reaction, *ent.User, *int, *utils.ErrorResponse) {
	q := client.Reaction.Query()
	var focusID *uuid.UUID
	var targetedObjAuthor *ent.User

	var postObj *ent.Post 
	var commentObj *ent.Comment 
	var replyObj *ent.Reply 
	if focus == "POST" {
		// Get Post Object and Query reactions for the post
		post, errCode, errData := PostManager{}.GetBySlug(client, slug, true)
		if errCode != nil {
			return nil, nil, errCode, errData
		}
		focusID = &post.ID
		q = q.Where(reaction.PostID(*focusID))
		targetedObjAuthor = post.Edges.Author
		postObj = post
	} else if focus == "COMMENT" {
		// Get Comment Object and Query reactions for the comment
		comment, errCode, errData := CommentManager{}.GetBySlug(client, slug, true)
		if errCode != nil {
			return nil, nil, errCode, errData
		}
		focusID = &comment.ID
		q = q.Where(reaction.CommentID(*focusID))
		targetedObjAuthor = comment.Edges.Author
		commentObj = comment
	} else {
		// Get Reply Object and Query reactions for the reply
		reply, errCode, errData := ReplyManager{}.GetBySlug(client, slug, true)
		if errCode != nil {
			return nil, nil, errCode, errData
		}
		focusID = &reply.ID
		q = q.Where(reaction.ReplyID(*focusID))
		targetedObjAuthor = reply.Edges.Author
		replyObj = reply
	}

	reaction, _ := q.WithUser(func(uq *ent.UserQuery) { uq.WithAvatar() }).Only(Ctx)
	if reaction == nil {
		// Create reaction
		reaction = obj.Create(client, user.ID, focus, *focusID, rtype)
	} else {
		// Update
		reaction = obj.Update(reaction, focus, *focusID, rtype)
	}

	// Set Related Data
	reaction.Edges.User = user
	reaction.Edges.Post = postObj
	reaction.Edges.Comment = commentObj
	reaction.Edges.Reply = replyObj
	return reaction, targetedObjAuthor, nil, nil
}

func (obj ReactionManager) GetByID(client *ent.Client, id uuid.UUID) (*ent.Reaction, *int, *utils.ErrorResponse) {
	r, _ := client.Reaction.Query().Where(reaction.ID(id)).
		WithPost().
		WithComment().
		WithReply().
		Only(Ctx)
	if r == nil {
		statusCode := 404
		errData := utils.RequestErr(utils.ERR_NON_EXISTENT, "Reaction does not exist")
		return nil, &statusCode, &errData
	}
	return r, nil, nil
}

func (obj ReactionManager) DropData(client *ent.Client) {
	client.Reaction.Delete().ExecX(Ctx)
}