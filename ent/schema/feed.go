package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

var CommonFeedFields = append(
	CommonFields,
	field.String("text").NotEmpty(),
	field.String("slug").NotEmpty(),
	field.UUID("author_id", uuid.UUID{}),
)

var CommonFeedEdges = []ent.Edge{
	edge.To("reactions", Reaction.Type),
}

// --------------------------------
// POST MODEL
// --------------------------------
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return append(
		CommonFeedFields,
		field.UUID("image_id", uuid.UUID{}).Optional().Nillable(),
	)
}

// Edges of the Post. (Relationship)
func (Post) Edges() []ent.Edge {
	return append(
		CommonFeedEdges,
		edge.From("author", User.Type).Ref("posts").Field("author_id").Required().Unique(),
		edge.From("image", File.Type).Ref("posts").Field("image_id").Unique(),
		edge.To("comments", Comment.Type),
	)
}

// --------------------------------
// COMMENT MODEL
// --------------------------------
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return append(
		CommonFeedFields,
		field.UUID("post_id", uuid.UUID{}),
	)
}

// Edges of the Comment. (Relationship)
func (Comment) Edges() []ent.Edge {
	return append(
		CommonFeedEdges,
		edge.From("author", User.Type).Ref("comments").Field("author_id").Required().Unique(),
		edge.From("post", Post.Type).Ref("comments").Field("post_id").Required().Unique(),
		edge.To("replies", Reply.Type),
	)
}

// --------------------------------
// REPLY MODEL
// --------------------------------
type Reply struct {
	ent.Schema
}

// Fields of the Reply.
func (Reply) Fields() []ent.Field {
	return append(
		CommonFeedFields,
		field.UUID("comment_id", uuid.UUID{}),
	)
}

// Edges of the Reply. (Relationship)
func (Reply) Edges() []ent.Edge {
	return append(
		CommonFeedEdges,
		edge.From("author", User.Type).Ref("replies").Field("author_id").Required().Unique(),
		edge.From("comment", Comment.Type).Ref("replies").Field("comment_id").Required().Unique(),
	)
}

// --------------------------------
// REACTION MODEL
// --------------------------------
type Reaction struct {
	ent.Schema
}

// Fields of the Reaction.
func (Reaction) Fields() []ent.Field {
	return append(
		CommonFields,
		field.UUID("user_id", uuid.UUID{}),
		field.String("rtype").NotEmpty(),
		field.UUID("post_id", uuid.UUID{}).Optional().Nillable(),
		field.UUID("comment_id", uuid.UUID{}).Optional().Nillable(),
		field.UUID("reply_id", uuid.UUID{}).Optional().Nillable(),
	)
}

// Edges of the Reaction. (Relationship)
func (Reaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("reactions").Field("user_id").Required().Unique(),
		edge.From("post", Post.Type).Ref("reactions").Field("post_id").Unique(),
		edge.From("comment", Comment.Type).Ref("reactions").Field("comment_id").Unique(),
		edge.From("reply", Reply.Type).Ref("reactions").Field("reply_id").Unique(),
	}
}

// Indexes of the Reaction.
func (Reaction) Indexes() []ent.Index {
    return []ent.Index{
        // Create a unique constraint based on user_id and post_id.
        index.Fields("user_id", "post_id").
            Unique(),
		// Create a unique constraint based on user_id and comment_id.
        index.Fields("user_id", "comment_id").
            Unique(),
		// Create a unique constraint based on user_id and reply_id.
        index.Fields("user_id", "reply_id").
            Unique(),
    }
}