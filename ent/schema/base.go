package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

var CommonFields = []ent.Field{
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
