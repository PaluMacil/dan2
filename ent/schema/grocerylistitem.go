package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// GroceryListItem holds the schema definition for the GroceryListItem entity.
type GroceryListItem struct {
	ent.Schema
}

// Fields of the GroceryListItem.
func (GroceryListItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(
				entproto.Field(2),
			),
		field.Int("quantity").
			Default(0).
			Annotations(
				entproto.Field(3),
			),
		field.String("note").
			Default("").
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

// Edges of the GroceryListItem.
func (GroceryListItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("grocery_list", GroceryList.Type).
			Ref("grocery_list_items").
			Unique().
			Annotations(entproto.Field(6)),
	}
}

func (GroceryListItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
