package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// GroceryList holds the schema definition for the GroceryList entity.
type GroceryList struct {
	ent.Schema
}

// Fields of the GroceryList.
func (GroceryList) Fields() []ent.Field {
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
		field.Bool("archived").
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

// Edges of the GroceryList.
func (GroceryList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("grocery_list_items", GroceryListItem.Type).
			Annotations(entproto.Field(6)),
		edge.From("owner", User.Type).
			Ref("grocery_lists").
			Unique().
			Annotations(entproto.Field(7)),
		edge.To("grocery_list_shares", GroceryListShare.Type).
			Annotations(entproto.Field(8)),
	}
}

func (GroceryList) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
