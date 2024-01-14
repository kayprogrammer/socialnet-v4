package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type City struct {
	ent.Schema
}

// Fields of the City.
func (City) Fields() []ent.Field {
	return append(
		CommonFields,
		field.String("name").
			NotEmpty(),
	)
}

// Edges of the City.
func (City) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}

type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return append(
		CommonFields,
		field.String("first_name").
			NotEmpty(),
		field.String("last_name").
			NotEmpty(),
		field.String("username").
			NotEmpty().Unique(),
		field.String("email").
			NotEmpty(),
		field.String("password").
			NotEmpty(),
		field.Bool("terms_agreement").
			Default(false),
		field.Bool("is_email_verified").
			Default(false),
		field.Bool("is_staff").
			Default(false),
		field.Bool("is_active").
			Default(true),
		field.String("bio").
			Optional().Nillable(),
		field.Time("dob").
			Optional().Nillable(),
		field.String("access").
			Optional().Nillable(),
		field.String("refresh").
			Optional().Nillable(),
		field.UUID("city_id", uuid.UUID{}).Optional().Nillable(),
		field.UUID("avatar_id", uuid.UUID{}).Optional().Nillable(),
	)
}

// Edges of the User. (Relationship)
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("city", City.Type).Ref("users").Field("city_id").Unique(),
		edge.From("avatar", File.Type).Ref("users").Field("avatar_id").Unique(),
		edge.To("otp", Otp.Type).Unique(),
        edge.To("posts", Post.Type),
        edge.To("reactions", Reaction.Type),
        edge.To("comments", Comment.Type),
        edge.To("replies", Reply.Type),
	}
}

type Otp struct {
	ent.Schema
}

// Fields of the City.
func (Otp) Fields() []ent.Field {
	return append(
		CommonFields,
		field.Uint32("code"),
		field.UUID("user_id", uuid.UUID{}),
	)
}

// Edges of the Otp.
func (Otp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("otp").Field("user_id").Unique().Required(),
	}
}