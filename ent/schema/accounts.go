package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Country struct {
	ent.Schema
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return append(
		CommonFields,
		field.String("name").
			NotEmpty(),
		field.String("code").
			NotEmpty(),
	)
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cities", City.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("regions", Region.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

type Region struct {
	ent.Schema
}

// Fields of the Region.
func (Region) Fields() []ent.Field {
	return append(
		CommonFields,
		field.String("name").
			NotEmpty(),
		field.UUID("country_id", uuid.UUID{}),
	)
}

// Edges of the Region.
func (Region) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("country", Country.Type).Ref("regions").Field("country_id").Unique().Required(),
		edge.To("cities", City.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
	}
}

type City struct {
	ent.Schema
}

// Fields of the City.
func (City) Fields() []ent.Field {
	return append(
		CommonFields,
		field.String("name").
			NotEmpty(),
		field.UUID("region_id", uuid.UUID{}).Nillable().Optional(),
		field.UUID("country_id", uuid.UUID{}),
	)
}

// Edges of the City.
func (City) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("region", Region.Type).Ref("cities").Field("region_id").Unique(),
		edge.From("country", Country.Type).Ref("cities").Field("country_id").Unique().Required(),
		edge.To("users", User.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
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
		edge.From("city", City.Type).Ref("users").Field("city_id").Unique().Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.From("avatar", File.Type).Ref("users").Field("avatar_id").Unique().Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.To("otp", Otp.Type).Unique().Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("posts", Post.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("reactions", Reaction.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("comments", Comment.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("replies", Reply.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("requester_friends", Friend.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("requestee_friends", Friend.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("notifications_from", Notification.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("notifications", Notification.Type),
		edge.To("notifications_read", Notification.Type),
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
