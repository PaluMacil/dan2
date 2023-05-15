package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// MovieList holds the schema definition for the MovieList entity.
type MovieList struct {
	ent.Schema
}

// Fields of the MovieList.
func (MovieList) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("My List"),
		field.String("note").
			Default(""),
		field.Bool("show_watched").
			Default(false),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the MovieList.
func (MovieList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("movies", Movie.Type),
		edge.From("owner", User.Type).
			Ref("movie_lists"),
		edge.To("movie_list_shares", MovieListShare.Type),
	}
}
