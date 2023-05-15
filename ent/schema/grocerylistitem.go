package schema

import (
	"entgo.io/ent"
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
		field.String("name"),
		field.Int("quantity").
			Default(0),
		field.String("note").
			Default(""),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the GroceryListItem.
func (GroceryListItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("grocery_list", GroceryList.Type).
			Ref("grocery_list_items"),
	}
}
