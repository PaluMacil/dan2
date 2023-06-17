package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(
				entproto.Field(2),
			),
		field.String("note").
			Default("").
			Annotations(
				entproto.Field(3),
			),
		field.Bool("watched").
			Default(false).
			Annotations(
				entproto.Field(4),
			),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entproto.Field(5),
			),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("movie_collection", MovieCollection.Type).
			Ref("movies").
			Unique().
			Annotations(entproto.Field(6)),
	}
}

func (Movie) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
