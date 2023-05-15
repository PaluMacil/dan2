package schema

import (
	"entgo.io/ent"
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
		field.String("name"),
		field.String("note").
			Default(""),
		field.Bool("watched").
			Default(false),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("movie_list", MovieList.Type).
			Ref("movies"),
	}
}
