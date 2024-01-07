package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"entgo.io/ent/schema/edge"
)

var CommonFields = []ent.Field{
	field.UUID("id", uuid.UUID{}).
		Default(uuid.New).
		Unique().StructTag("json:'-'"),
	field.Time("created_at").
		Default(time.Now).StructTag("json:'-'"),
	field.Time("updated_at").
		Default(time.Now).StructTag("json:'-'").
		UpdateDefault(time.Now),
}

type File struct {
	ent.Schema
}

func (File) Fields() []ent.Field {
	return append(
		CommonFields, 
		field.String("resource_type").
			NotEmpty(),
	)
}

// Edges of the File.
func (File) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("users", User.Type),
    }
}