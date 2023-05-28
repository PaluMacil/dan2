package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
			Default(false).
			Annotations(
				entproto.Field(2),
			),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entproto.Field(3),
			),
	}
}

// Edges of the GroceryListShare.
func (GroceryListShare) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("grocery_list_shares").
			Unique().
			Annotations(entproto.Field(4)),
		edge.From("grocery_list", GroceryList.Type).
			Ref("grocery_list_shares").
			Unique().
			Annotations(entproto.Field(5)),
	}
}

func (GroceryListShare) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
