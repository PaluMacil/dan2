// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/PaluMacil/dan2/ent/moviecollection"
	"github.com/PaluMacil/dan2/ent/user"
)

// MovieCollection is the model entity for the MovieCollection schema.
type MovieCollection struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// ShowWatched holds the value of the "show_watched" field.
	ShowWatched bool `json:"show_watched,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MovieCollectionQuery when eager-loading is set.
	Edges                  MovieCollectionEdges `json:"edges"`
	user_movie_collections *int
	selectValues           sql.SelectValues
}

// MovieCollectionEdges holds the relations/edges for other nodes in the graph.
type MovieCollectionEdges struct {
	// Movies holds the value of the movies edge.
	Movies []*Movie `json:"movies,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// MovieCollectionShares holds the value of the movie_collection_shares edge.
	MovieCollectionShares []*MovieCollectionShare `json:"movie_collection_shares,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// MoviesOrErr returns the Movies value or an error if the edge
// was not loaded in eager-loading.
func (e MovieCollectionEdges) MoviesOrErr() ([]*Movie, error) {
	if e.loadedTypes[0] {
		return e.Movies, nil
	}
	return nil, &NotLoadedError{edge: "movies"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MovieCollectionEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// MovieCollectionSharesOrErr returns the MovieCollectionShares value or an error if the edge
// was not loaded in eager-loading.
func (e MovieCollectionEdges) MovieCollectionSharesOrErr() ([]*MovieCollectionShare, error) {
	if e.loadedTypes[2] {
		return e.MovieCollectionShares, nil
	}
	return nil, &NotLoadedError{edge: "movie_collection_shares"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MovieCollection) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case moviecollection.FieldShowWatched:
			values[i] = new(sql.NullBool)
		case moviecollection.FieldID:
			values[i] = new(sql.NullInt64)
		case moviecollection.FieldName, moviecollection.FieldNote:
			values[i] = new(sql.NullString)
		case moviecollection.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case moviecollection.ForeignKeys[0]: // user_movie_collections
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MovieCollection fields.
func (mc *MovieCollection) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case moviecollection.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mc.ID = int(value.Int64)
		case moviecollection.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mc.Name = value.String
			}
		case moviecollection.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				mc.Note = value.String
			}
		case moviecollection.FieldShowWatched:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field show_watched", values[i])
			} else if value.Valid {
				mc.ShowWatched = value.Bool
			}
		case moviecollection.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mc.CreatedAt = value.Time
			}
		case moviecollection.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_movie_collections", value)
			} else if value.Valid {
				mc.user_movie_collections = new(int)
				*mc.user_movie_collections = int(value.Int64)
			}
		default:
			mc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MovieCollection.
// This includes values selected through modifiers, order, etc.
func (mc *MovieCollection) Value(name string) (ent.Value, error) {
	return mc.selectValues.Get(name)
}

// QueryMovies queries the "movies" edge of the MovieCollection entity.
func (mc *MovieCollection) QueryMovies() *MovieQuery {
	return NewMovieCollectionClient(mc.config).QueryMovies(mc)
}

// QueryOwner queries the "owner" edge of the MovieCollection entity.
func (mc *MovieCollection) QueryOwner() *UserQuery {
	return NewMovieCollectionClient(mc.config).QueryOwner(mc)
}

// QueryMovieCollectionShares queries the "movie_collection_shares" edge of the MovieCollection entity.
func (mc *MovieCollection) QueryMovieCollectionShares() *MovieCollectionShareQuery {
	return NewMovieCollectionClient(mc.config).QueryMovieCollectionShares(mc)
}

// Update returns a builder for updating this MovieCollection.
// Note that you need to call MovieCollection.Unwrap() before calling this method if this MovieCollection
// was returned from a transaction, and the transaction was committed or rolled back.
func (mc *MovieCollection) Update() *MovieCollectionUpdateOne {
	return NewMovieCollectionClient(mc.config).UpdateOne(mc)
}

// Unwrap unwraps the MovieCollection entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mc *MovieCollection) Unwrap() *MovieCollection {
	_tx, ok := mc.config.driver.(*txDriver)
	if !ok {
		panic("ent: MovieCollection is not a transactional entity")
	}
	mc.config.driver = _tx.drv
	return mc
}

// String implements the fmt.Stringer.
func (mc *MovieCollection) String() string {
	var builder strings.Builder
	builder.WriteString("MovieCollection(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mc.ID))
	builder.WriteString("name=")
	builder.WriteString(mc.Name)
	builder.WriteString(", ")
	builder.WriteString("note=")
	builder.WriteString(mc.Note)
	builder.WriteString(", ")
	builder.WriteString("show_watched=")
	builder.WriteString(fmt.Sprintf("%v", mc.ShowWatched))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(mc.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MovieCollections is a parsable slice of MovieCollection.
type MovieCollections []*MovieCollection
