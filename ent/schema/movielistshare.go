package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// MovieListShare holds the schema definition for the MovieListShare entity.
type MovieListShare struct {
	ent.Schema
}

// Fields of the MovieListShare.
func (MovieListShare) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("can_edit").
			Default(false),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the MovieListShare.
func (MovieListShare) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("movie_list_shares").
			Unique(),
		edge.From("movie_list", MovieList.Type).
			Ref("movie_list_shares").
			Unique(),
	}
}
