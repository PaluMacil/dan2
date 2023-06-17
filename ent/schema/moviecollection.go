package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// MovieCollection holds the schema definition for the MovieCollection entity.
type MovieCollection struct {
	ent.Schema
}

// Fields of the MovieCollection.
func (MovieCollection) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("My List").
			Annotations(
				entproto.Field(2),
			),
		field.String("note").
			Default("").
			Annotations(
				entproto.Field(3),
			),
		field.Bool("show_watched").
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

// Edges of the MovieCollection.
func (MovieCollection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("movies", Movie.Type).
			Annotations(entproto.Field(6)),
		edge.From("owner", User.Type).
			Ref("movie_collections").
			Unique().
			Annotations(entproto.Field(7)),
		edge.To("movie_collection_shares", MovieCollectionShare.Type).
			Annotations(entproto.Field(8)),
	}
}

func (MovieCollection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
