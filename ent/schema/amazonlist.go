package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// AmazonList holds the schema definition for the AmazonList entity.
type AmazonList struct {
	ent.Schema
}

// Fields of the AmazonList.
func (AmazonList) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entproto.Field(2),
			),
	}
}

// Edges of the AmazonList.
func (AmazonList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("amazon_orders", AmazonOrder.Type).
			Annotations(entproto.Field(3)),
		edge.From("owner", User.Type).
			Ref("amazon_lists").
			Annotations(entproto.Field(4)),
		edge.To("amazon_shares", AmazonShare.Type).
			Annotations(entproto.Field(5)),
	}
}

func (AmazonList) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
