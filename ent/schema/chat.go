package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Chat struct {
	ent.Schema
}

// Fields of the Chat.
func (Chat) Fields() []ent.Field {
	return append(
		CommonFields,
		field.String("name").
			Nillable().Optional(),
		field.Enum("ctype").Values("DM", "GROUP"),
		field.String("description").
			Nillable().Optional(),
		field.UUID("owner_id", uuid.UUID{}),
		field.UUID("image_id", uuid.UUID{}).Optional().Nillable(),
	)
}

// Edges of the Chat.
func (Chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("owned_chats").Field("owner_id").Unique().Required(),
		edge.From("image", File.Type).Ref("chats").Field("image_id").Unique(),
		edge.From("users", User.Type).Ref("member_chats"),
		edge.To("messages", Message.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

// Annotations of the Friend.
func (Chat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		&entsql.Annotation{
			// Check constraint to prevent requester and requestee from being the same user
			Checks: map[string]string{
				"dm_chat_constraints":    "(ctype = 'DM' AND name IS NULL AND description IS NULL AND image_id IS NULL) OR (ctype = 'GROUP')", // DMs cannot have name, image and description
				"group_chat_constraints": "(ctype = 'GROUP' AND name IS NOT NULL) OR (ctype = 'DM')",                                          // Enter name for group chat
			},
		},
	}
}

type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return append(
		CommonFields,
		field.String("text").
			Nillable().Optional(),
		field.UUID("sender_id", uuid.UUID{}),
		field.UUID("chat_id", uuid.UUID{}),
		field.UUID("file_id", uuid.UUID{}).Optional().Nillable(),
	)
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sender", User.Type).Ref("messages").Field("sender_id").Unique().Required(),
		edge.From("chat", Chat.Type).Ref("messages").Field("chat_id").Unique().Required(),
		edge.From("file", File.Type).Ref("messages").Field("file_id").Unique(),
	}
}

// Hooks of the User.
// func (User) Hooks() []ent.Hook {
// 	return []ent.Hook{
// 		hook.On(
// 			hook.OpCreate,
// 			hook.HandlerFunc(func(ctx context.Context, tx *ent.Tx, user *User) error {
// 				fmt.Println("Executing custom logic before create operation")
// 				return nil
// 			}),
// 		),
// 	}
// }
