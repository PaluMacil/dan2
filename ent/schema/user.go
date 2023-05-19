package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("display_name").
			Unique(),
		field.Int8("name_changes").
			Default(0),
		field.String("email").
			Unique(),
		field.Bool("verified").
			Default(false),
		field.Bool("locked").
			Default(false),
		//TODO: consider defaulting to zero time but check how js sees that
		field.Time("last_login"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("amazon_shares", AmazonShare.Type),
		edge.To("amazon_lists", AmazonList.Type),
		edge.To("drinks", Drink.Type),
		edge.To("grocery_lists", GroceryList.Type),
		edge.To("grocery_list_shares", GroceryListShare.Type),
		edge.To("movie_lists", MovieList.Type),
		edge.To("movie_list_shares", MovieListShare.Type),
	}
}
