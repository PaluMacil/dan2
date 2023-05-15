package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// AmazonOrder holds the schema definition for the AmazonOrder entity.
type AmazonOrder struct {
	ent.Schema
}

// Fields of the AmazonOrder.
func (AmazonOrder) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("category"),
		field.String("brand"),
		field.String("seller"),
		field.String("address1"),
		field.String("address2"),
		field.String("city"),
		field.String("state"),
		field.String("zip"),
		field.Float32("price"),
		field.Float32("tax"),
		field.Bool("refund"),
		field.Time("ordered_at"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the AmazonOrder.
func (AmazonOrder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("amazon_list", AmazonList.Type).
			Ref("amazon_orders"),
	}
}
