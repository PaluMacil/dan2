package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
			Unique().
			Annotations(
				entproto.Field(2),
			),
		field.Int8("name_changes").
			Default(0).
			Annotations(
				entproto.Field(3),
			),
		field.String("email").
			Unique().
			Annotations(
				entproto.Field(4),
			),
		field.Bool("verified").
			Default(false).
			Annotations(
				entproto.Field(5),
			),
		field.Bool("locked").
			Default(false).
			Annotations(
				entproto.Field(6),
			),
		field.Time("last_login").
			Nillable().
			Annotations(
				entproto.Field(7),
			),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entproto.Field(8),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("amazon_shares", AmazonShare.Type).
			Annotations(entproto.Field(9)),
		edge.To("amazon_lists", AmazonList.Type).
			Annotations(entproto.Field(10)),
		edge.To("drinks", Drink.Type).
			Annotations(entproto.Field(11)),
		edge.To("grocery_lists", GroceryList.Type).
			Annotations(entproto.Field(12)),
		edge.To("grocery_list_shares", GroceryListShare.Type).
			Annotations(entproto.Field(13)),
		edge.To("movie_collections", MovieCollection.Type).
			Annotations(entproto.Field(14)),
		edge.To("movie_collection_shares", MovieCollectionShare.Type).
			Annotations(entproto.Field(15)),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
