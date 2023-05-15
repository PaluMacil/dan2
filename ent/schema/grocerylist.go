package schema

import (
	"entgo.io/ent"
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
			Default("My List"),
		field.String("note").
			Default(""),
		field.Bool("archived").
			Default(false),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the GroceryList.
func (GroceryList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("grocery_list_items", GroceryListItem.Type),
		edge.From("owner", User.Type).
			Ref("grocery_lists"),
		edge.To("grocery_list_shares", GroceryListShare.Type),
	}
}
