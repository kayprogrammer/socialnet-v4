package ent

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SiteDetail struct {
	BaseModel
}

// Fields of the SiteDetail.
func (SiteDetail) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("SocialNet"),
        field.String("email").
			Default("kayprogrammer1@gmail.com"),
        field.String("phone").
			Default("+2348133831036"),
        field.String("address").
			Default("234, Lagos, Nigeria"),
        field.String("fb").
			Default("https://facebook.com"),
        field.String("tw").
			Default("https://twitter.com"),
        field.String("wh").
			Default("https://wa.me/2348133831036"),
        field.String("ig").
			Default("https://instagram.com"),
	}
}
