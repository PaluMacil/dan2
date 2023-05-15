package schema

import (
	"entgo.io/ent"
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
			Default(time.Now),
	}
}

// Edges of the AmazonList.
func (AmazonList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("amazon_orders", AmazonOrder.Type),
		edge.From("owner", User.Type).
			Ref("amazon_lists"),
		edge.To("amazon_shares", AmazonShare.Type),
	}
}
