// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/PaluMacil/dan2/ent/drink"
)

// Drink is the model entity for the Drink schema.
type Drink struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type drink.Type `json:"type,omitempty"`
	// Abv holds the value of the "abv" field.
	Abv int8 `json:"abv,omitempty"`
	// Ounces holds the value of the "ounces" field.
	Ounces int8 `json:"ounces,omitempty"`
	// Year holds the value of the "year" field.
	Year int `json:"year,omitempty"`
	// Month holds the value of the "month" field.
	Month int `json:"month,omitempty"`
	// Day holds the value of the "day" field.
	Day int `json:"day,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DrinkQuery when eager-loading is set.
	Edges        DrinkEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DrinkEdges holds the relations/edges for other nodes in the graph.
type DrinkEdges struct {
	// Owner holds the value of the owner edge.
	Owner []*User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading.
func (e DrinkEdges) OwnerOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Drink) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case drink.FieldID, drink.FieldAbv, drink.FieldOunces, drink.FieldYear, drink.FieldMonth, drink.FieldDay:
			values[i] = new(sql.NullInt64)
		case drink.FieldType, drink.FieldNote:
			values[i] = new(sql.NullString)
		case drink.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Drink fields.
func (d *Drink) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case drink.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case drink.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				d.Type = drink.Type(value.String)
			}
		case drink.FieldAbv:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field abv", values[i])
			} else if value.Valid {
				d.Abv = int8(value.Int64)
			}
		case drink.FieldOunces:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ounces", values[i])
			} else if value.Valid {
				d.Ounces = int8(value.Int64)
			}
		case drink.FieldYear:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field year", values[i])
			} else if value.Valid {
				d.Year = int(value.Int64)
			}
		case drink.FieldMonth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field month", values[i])
			} else if value.Valid {
				d.Month = int(value.Int64)
			}
		case drink.FieldDay:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field day", values[i])
			} else if value.Valid {
				d.Day = int(value.Int64)
			}
		case drink.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				d.Note = value.String
			}
		case drink.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Drink.
// This includes values selected through modifiers, order, etc.
func (d *Drink) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Drink entity.
func (d *Drink) QueryOwner() *UserQuery {
	return NewDrinkClient(d.config).QueryOwner(d)
}

// Update returns a builder for updating this Drink.
// Note that you need to call Drink.Unwrap() before calling this method if this Drink
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Drink) Update() *DrinkUpdateOne {
	return NewDrinkClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Drink entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Drink) Unwrap() *Drink {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Drink is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Drink) String() string {
	var builder strings.Builder
	builder.WriteString("Drink(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", d.Type))
	builder.WriteString(", ")
	builder.WriteString("abv=")
	builder.WriteString(fmt.Sprintf("%v", d.Abv))
	builder.WriteString(", ")
	builder.WriteString("ounces=")
	builder.WriteString(fmt.Sprintf("%v", d.Ounces))
	builder.WriteString(", ")
	builder.WriteString("year=")
	builder.WriteString(fmt.Sprintf("%v", d.Year))
	builder.WriteString(", ")
	builder.WriteString("month=")
	builder.WriteString(fmt.Sprintf("%v", d.Month))
	builder.WriteString(", ")
	builder.WriteString("day=")
	builder.WriteString(fmt.Sprintf("%v", d.Day))
	builder.WriteString(", ")
	builder.WriteString("note=")
	builder.WriteString(d.Note)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Drinks is a parsable slice of Drink.
type Drinks []*Drink