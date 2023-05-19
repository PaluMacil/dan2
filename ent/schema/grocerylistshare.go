package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// GroceryListShare holds the schema definition for the GroceryListShare entity.
type GroceryListShare struct {
	ent.Schema
}

// Fields of the GroceryListShare.
func (GroceryListShare) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("can_edit").
			Default(false),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the GroceryListShare.
func (GroceryListShare) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("grocery_list_shares").
			Unique(),
		edge.From("grocery_list", GroceryList.Type).
			Ref("grocery_list_shares").
			Unique(),
	}
}
