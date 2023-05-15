package schema

import (
	"entgo.io/ent"
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
			Default(false),
		field.Time("created_at"),
	}
}

// Edges of the AmazonShare.
func (AmazonShare) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("amazon_shares"),
		edge.From("amazon_list", AmazonList.Type).
			Ref("amazon_shares"),
	}
}
