package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SiteDetail struct {
	ent.Schema
}

// Fields of the SiteDetail.
func (SiteDetail) Fields() []ent.Field {
	return append(
		CommonFields, 
		field.String("name").
			Default("SocialNet").
			StructTag(`example:"SocialNet"`),
        field.String("email").
			Default("kayprogrammer1@gmail.com").
			StructTag(`example:"kayprogrammer1@gmail.com"`),
        field.String("phone").
			Default("+2348133831036").
			StructTag(`example:"+2348133831036"`),
        field.String("address").
			Default("234, Lagos, Nigeria").
			StructTag(`example:"234, Lagos, Nigeria"`),
        field.String("fb").
			Default("https://facebook.com").
			StructTag(`example:"https://facebook.com"`),
        field.String("tw").
			Default("https://twitter.com").
			StructTag(`example:"https://twitter.com"`),
        field.String("wh").
			Default("https://wa.me/2348133831036").
			StructTag(`example:"https://wa.me/2348133831036"`),
        field.String("ig").
			Default("https://instagram.com").
			StructTag(`example:"https://instagram.com"`),
	)
}
