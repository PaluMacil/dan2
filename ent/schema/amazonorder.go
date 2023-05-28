package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("name").
			Annotations(
				entproto.Field(2),
			),
		field.String("category").
			Annotations(
				entproto.Field(3),
			),
		field.String("brand").
			Annotations(
				entproto.Field(4),
			),
		field.String("seller").
			Annotations(
				entproto.Field(5),
			),
		field.String("address1").
			Annotations(
				entproto.Field(6),
			),
		field.String("address2").
			Annotations(
				entproto.Field(7),
			),
		field.String("city").
			Annotations(
				entproto.Field(8),
			),
		field.String("state").
			Annotations(
				entproto.Field(9),
			),
		field.String("zip").
			Annotations(
				entproto.Field(10),
			),
		field.Float32("price").
			Annotations(
				entproto.Field(11),
			),
		field.Float32("tax").
			Annotations(
				entproto.Field(12),
			),
		field.Bool("refund").
			Annotations(
				entproto.Field(13),
			),
		field.Time("ordered_at").
			Annotations(
				entproto.Field(14),
			),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entproto.Field(15),
			),
	}
}

// Edges of the AmazonOrder.
func (AmazonOrder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("amazon_list", AmazonList.Type).
			Ref("amazon_orders").
			Unique().
			Annotations(entproto.Field(16)),
	}
}

func (AmazonOrder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
