package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Drink holds the schema definition for the Drink entity.
type Drink struct {
	ent.Schema
}

// Fields of the Drink.
func (Drink) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values("unknown", "lightbeer", "craftbeer", "wine", "liquor", "highball", "cocktail").
			Default("unknown"),
		field.Int8("abv").
			Default(0),
		field.Int8("ounces").
			Default(0),
		field.Int("year").
			DefaultFunc(NowYear),
		field.Int("month").
			DefaultFunc(NowMonth),
		field.Int("day").
			DefaultFunc(NowDay),
		field.String("note").
			Default(""),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Drink.
func (Drink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("drinks"),
	}
}
