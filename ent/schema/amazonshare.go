package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AmazonShare holds the schema definition for the AmazonShare entity.
type AmazonShare struct {
	ent.Schema
}

// Fields of the AmazonShare.
func (AmazonShare) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("can_edit").
			Default(false).
			Annotations(
				entproto.Field(2),
			),
		field.Time("created_at").
			Annotations(
				entproto.Field(3),
			),
	}
}

// Edges of the AmazonShare.
func (AmazonShare) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("amazon_shares").
			Unique().
			Annotations(entproto.Field(4)),
		edge.From("amazon_list", AmazonList.Type).
			Ref("amazon_shares").
			Unique().
			Annotations(entproto.Field(5)),
	}
}

func (AmazonShare) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
