package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

var CommonFields = []ent.Field{
	field.UUID("id", uuid.UUID{}).
		Default(uuid.New),
	field.Time("created_at").
		Default(time.Now),
	field.Time("updated_at").
		Default(time.Now).
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
        edge.To("users", User.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
        edge.To("posts", Post.Type).Annotations(entsql.OnDelete(entsql.SetNull)),
    }
}
