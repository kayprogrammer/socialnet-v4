package ent

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type BaseModel struct {
	ent.Schema
}

// Fields of the BaseModel.
func (BaseModel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			StructTag(`json:"oid,omitempty"`),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
            UpdateDefault(time.Now),
	}
}
