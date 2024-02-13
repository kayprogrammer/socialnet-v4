package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Friend struct {
	ent.Schema
}

// Fields of the Friend.
func (Friend) Fields() []ent.Field {
	return append(
		CommonFields,
		field.UUID("requester_id", uuid.UUID{}),
		field.UUID("requestee_id", uuid.UUID{}),
		field.Enum("status").Values("PENDING", "ACCEPTED").Default("PENDING"),
	)
}

// Edges of the Friend.
func (Friend) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("requester", User.Type).Ref("requester_friends").Field("requester_id").Unique().Required(),
		edge.From("requestee", User.Type).Ref("requestee_friends").Field("requestee_id").Unique().Required(),
	}
}

// Annotations of the Friend.
func (Friend) Annotations() []schema.Annotation {
	return []schema.Annotation{
		&entsql.Annotation{
			// Check constraint to prevent requester and requestee from being the same user
			Checks: map[string]string{
				"different_users": "requester_id <> requestee_id",
			},
		},
	}
}

// Indexes of the Friend.
func (Friend) Indexes() []ent.Index {
	return []ent.Index{
		// Would have done a bidirectional unique constraint 
		// but it doesn't seem possible with ent at the moment
		// Let's make do with this.
		index.Fields("requester_id", "requestee_id").
			Unique().StorageKey("unique_requester_requestee_combination"),
		index.Fields("requestee_id", "requester_id").
			Unique().StorageKey("unique_requestee_requester_combination"),
	}
}


type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return append(
		CommonFields,
		field.UUID("sender_id", uuid.UUID{}).Nillable().Optional(),
		field.Enum("ntype").Values("REACTION", "COMMENT", "REPLY", "ADMIN"),
		field.UUID("post_id", uuid.UUID{}).Nillable().Optional(),    // For reactions or admin reference to a post
		field.UUID("comment_id", uuid.UUID{}).Nillable().Optional(), // For comments and reactions
		field.UUID("reply_id", uuid.UUID{}).Nillable().Optional(),   // For replies and reactions
		field.String("text").Nillable().Optional(),                  // For admin notifications only
	)
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sender", User.Type).Ref("notifications_from").Field("sender_id").Unique(),
		edge.From("receivers", User.Type).Ref("notifications"),
		edge.From("post", Post.Type).Ref("notifications").Field("post_id").Unique(),
		edge.From("comment", Comment.Type).Ref("notifications").Field("comment_id").Unique(),
		edge.From("reply", Reply.Type).Ref("notifications").Field("reply_id").Unique(),
		edge.From("read_by", User.Type).Ref("notifications_read"),
	}
}
