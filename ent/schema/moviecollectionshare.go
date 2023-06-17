package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// MovieCollectionShare holds the schema definition for the MovieCollectionShare entity.
type MovieCollectionShare struct {
	ent.Schema
}

// Fields of the MovieCollectionShare.
func (MovieCollectionShare) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("can_edit").
			Default(false).
			Annotations(
				entproto.Field(2),
			),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entproto.Field(3),
			),
	}
}

// Edges of the MovieCollectionShare.
func (MovieCollectionShare) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("movie_collection_shares").
			Unique().
			Annotations(entproto.Field(4)),
		edge.From("movie_collection", MovieCollection.Type).
			Ref("movie_collection_shares").
			Unique().
			Annotations(entproto.Field(5)),
	}
}

func (MovieCollectionShare) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
