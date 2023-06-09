// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/PaluMacil/dan2/ent/movie"
	"github.com/PaluMacil/dan2/ent/moviecollection"
)

// Movie is the model entity for the Movie schema.
type Movie struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// Watched holds the value of the "watched" field.
	Watched bool `json:"watched,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MovieQuery when eager-loading is set.
	Edges                   MovieEdges `json:"edges"`
	movie_collection_movies *int
	selectValues            sql.SelectValues
}

// MovieEdges holds the relations/edges for other nodes in the graph.
type MovieEdges struct {
	// MovieCollection holds the value of the movie_collection edge.
	MovieCollection *MovieCollection `json:"movie_collection,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MovieCollectionOrErr returns the MovieCollection value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MovieEdges) MovieCollectionOrErr() (*MovieCollection, error) {
	if e.loadedTypes[0] {
		if e.MovieCollection == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: moviecollection.Label}
		}
		return e.MovieCollection, nil
	}
	return nil, &NotLoadedError{edge: "movie_collection"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Movie) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case movie.FieldWatched:
			values[i] = new(sql.NullBool)
		case movie.FieldID:
			values[i] = new(sql.NullInt64)
		case movie.FieldName, movie.FieldNote:
			values[i] = new(sql.NullString)
		case movie.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case movie.ForeignKeys[0]: // movie_collection_movies
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Movie fields.
func (m *Movie) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case movie.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case movie.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case movie.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				m.Note = value.String
			}
		case movie.FieldWatched:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field watched", values[i])
			} else if value.Valid {
				m.Watched = value.Bool
			}
		case movie.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case movie.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field movie_collection_movies", value)
			} else if value.Valid {
				m.movie_collection_movies = new(int)
				*m.movie_collection_movies = int(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Movie.
// This includes values selected through modifiers, order, etc.
func (m *Movie) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryMovieCollection queries the "movie_collection" edge of the Movie entity.
func (m *Movie) QueryMovieCollection() *MovieCollectionQuery {
	return NewMovieClient(m.config).QueryMovieCollection(m)
}

// Update returns a builder for updating this Movie.
// Note that you need to call Movie.Unwrap() before calling this method if this Movie
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Movie) Update() *MovieUpdateOne {
	return NewMovieClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Movie entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Movie) Unwrap() *Movie {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Movie is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Movie) String() string {
	var builder strings.Builder
	builder.WriteString("Movie(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("note=")
	builder.WriteString(m.Note)
	builder.WriteString(", ")
	builder.WriteString("watched=")
	builder.WriteString(fmt.Sprintf("%v", m.Watched))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Movies is a parsable slice of Movie.
type Movies []*Movie
