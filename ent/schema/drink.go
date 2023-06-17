package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
			Values("unknown", "light_beer", "craft_beer", "wine", "liquor", "highball", "cocktail").
			Default("unknown").
			Annotations(
				entproto.Field(2),
				entproto.Enum(map[string]int32{
					"unknown":    0,
					"light_beer": 1,
					"craft_beer": 2,
					"wine":       3,
					"liquor":     4,
					"highball":   5,
					"cocktail":   6,
				}),
			),
		field.Int8("abv").
			Default(0).
			Annotations(
				entproto.Field(3),
			),
		field.Int8("ounces").
			Default(0).
			Annotations(
				entproto.Field(4),
			),
		field.Int("year").
			DefaultFunc(NowYear).
			Annotations(
				entproto.Field(5),
			),
		field.Int("month").
			DefaultFunc(NowMonth).
			Annotations(
				entproto.Field(6),
			),
		field.Int("day").
			DefaultFunc(NowDay).
			Annotations(
				entproto.Field(7),
			),
		field.String("note").
			Default("").
			Annotations(
				entproto.Field(8),
			),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entproto.Field(9),
			),
	}
}

// Edges of the Drink.
func (Drink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("drinks").
			Unique().
			Annotations(entproto.Field(10)),
	}
}

func (Drink) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
