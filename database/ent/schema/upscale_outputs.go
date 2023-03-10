package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UpscaleOutput holds the schema definition for the UpscaleOutput entity.
type UpscaleOutput struct {
	ent.Schema
}

// Fields of the UpscaleOutput.
func (UpscaleOutput) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("image_path"),
		// ! Relationships / many-to-one
		field.UUID("upscale_id", uuid.UUID{}),
		// ! End relationships
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UpscaleOutput.
func (UpscaleOutput) Edges() []ent.Edge {
	return []ent.Edge{
		// M2O with upscales
		edge.From("upscales", Upscale.Type).
			Ref("upscale_outputs").
			Field("upscale_id").
			Required().
			Unique(),
	}
}

// Annotations of the UpscaleOutput.
func (UpscaleOutput) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "upscale_outputs"},
	}
}
