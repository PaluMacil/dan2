// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PaluMacil/dan2/ent/movie"
	"github.com/PaluMacil/dan2/ent/moviecollection"
	"github.com/PaluMacil/dan2/ent/moviecollectionshare"
	"github.com/PaluMacil/dan2/ent/predicate"
	"github.com/PaluMacil/dan2/ent/user"
)

// MovieCollectionUpdate is the builder for updating MovieCollection entities.
type MovieCollectionUpdate struct {
	config
	hooks    []Hook
	mutation *MovieCollectionMutation
}

// Where appends a list predicates to the MovieCollectionUpdate builder.
func (mcu *MovieCollectionUpdate) Where(ps ...predicate.MovieCollection) *MovieCollectionUpdate {
	mcu.mutation.Where(ps...)
	return mcu
}

// SetName sets the "name" field.
func (mcu *MovieCollectionUpdate) SetName(s string) *MovieCollectionUpdate {
	mcu.mutation.SetName(s)
	return mcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mcu *MovieCollectionUpdate) SetNillableName(s *string) *MovieCollectionUpdate {
	if s != nil {
		mcu.SetName(*s)
	}
	return mcu
}

// SetNote sets the "note" field.
func (mcu *MovieCollectionUpdate) SetNote(s string) *MovieCollectionUpdate {
	mcu.mutation.SetNote(s)
	return mcu
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (mcu *MovieCollectionUpdate) SetNillableNote(s *string) *MovieCollectionUpdate {
	if s != nil {
		mcu.SetNote(*s)
	}
	return mcu
}

// SetShowWatched sets the "show_watched" field.
func (mcu *MovieCollectionUpdate) SetShowWatched(b bool) *MovieCollectionUpdate {
	mcu.mutation.SetShowWatched(b)
	return mcu
}

// SetNillableShowWatched sets the "show_watched" field if the given value is not nil.
func (mcu *MovieCollectionUpdate) SetNillableShowWatched(b *bool) *MovieCollectionUpdate {
	if b != nil {
		mcu.SetShowWatched(*b)
	}
	return mcu
}

// SetCreatedAt sets the "created_at" field.
func (mcu *MovieCollectionUpdate) SetCreatedAt(t time.Time) *MovieCollectionUpdate {
	mcu.mutation.SetCreatedAt(t)
	return mcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcu *MovieCollectionUpdate) SetNillableCreatedAt(t *time.Time) *MovieCollectionUpdate {
	if t != nil {
		mcu.SetCreatedAt(*t)
	}
	return mcu
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (mcu *MovieCollectionUpdate) AddMovieIDs(ids ...int) *MovieCollectionUpdate {
	mcu.mutation.AddMovieIDs(ids...)
	return mcu
}

// AddMovies adds the "movies" edges to the Movie entity.
func (mcu *MovieCollectionUpdate) AddMovies(m ...*Movie) *MovieCollectionUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcu.AddMovieIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (mcu *MovieCollectionUpdate) SetOwnerID(id int) *MovieCollectionUpdate {
	mcu.mutation.SetOwnerID(id)
	return mcu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (mcu *MovieCollectionUpdate) SetNillableOwnerID(id *int) *MovieCollectionUpdate {
	if id != nil {
		mcu = mcu.SetOwnerID(*id)
	}
	return mcu
}

// SetOwner sets the "owner" edge to the User entity.
func (mcu *MovieCollectionUpdate) SetOwner(u *User) *MovieCollectionUpdate {
	return mcu.SetOwnerID(u.ID)
}

// AddMovieCollectionShareIDs adds the "movie_collection_shares" edge to the MovieCollectionShare entity by IDs.
func (mcu *MovieCollectionUpdate) AddMovieCollectionShareIDs(ids ...int) *MovieCollectionUpdate {
	mcu.mutation.AddMovieCollectionShareIDs(ids...)
	return mcu
}

// AddMovieCollectionShares adds the "movie_collection_shares" edges to the MovieCollectionShare entity.
func (mcu *MovieCollectionUpdate) AddMovieCollectionShares(m ...*MovieCollectionShare) *MovieCollectionUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcu.AddMovieCollectionShareIDs(ids...)
}

// Mutation returns the MovieCollectionMutation object of the builder.
func (mcu *MovieCollectionUpdate) Mutation() *MovieCollectionMutation {
	return mcu.mutation
}

// ClearMovies clears all "movies" edges to the Movie entity.
func (mcu *MovieCollectionUpdate) ClearMovies() *MovieCollectionUpdate {
	mcu.mutation.ClearMovies()
	return mcu
}

// RemoveMovieIDs removes the "movies" edge to Movie entities by IDs.
func (mcu *MovieCollectionUpdate) RemoveMovieIDs(ids ...int) *MovieCollectionUpdate {
	mcu.mutation.RemoveMovieIDs(ids...)
	return mcu
}

// RemoveMovies removes "movies" edges to Movie entities.
func (mcu *MovieCollectionUpdate) RemoveMovies(m ...*Movie) *MovieCollectionUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcu.RemoveMovieIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (mcu *MovieCollectionUpdate) ClearOwner() *MovieCollectionUpdate {
	mcu.mutation.ClearOwner()
	return mcu
}

// ClearMovieCollectionShares clears all "movie_collection_shares" edges to the MovieCollectionShare entity.
func (mcu *MovieCollectionUpdate) ClearMovieCollectionShares() *MovieCollectionUpdate {
	mcu.mutation.ClearMovieCollectionShares()
	return mcu
}

// RemoveMovieCollectionShareIDs removes the "movie_collection_shares" edge to MovieCollectionShare entities by IDs.
func (mcu *MovieCollectionUpdate) RemoveMovieCollectionShareIDs(ids ...int) *MovieCollectionUpdate {
	mcu.mutation.RemoveMovieCollectionShareIDs(ids...)
	return mcu
}

// RemoveMovieCollectionShares removes "movie_collection_shares" edges to MovieCollectionShare entities.
func (mcu *MovieCollectionUpdate) RemoveMovieCollectionShares(m ...*MovieCollectionShare) *MovieCollectionUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcu.RemoveMovieCollectionShareIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mcu *MovieCollectionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mcu.sqlSave, mcu.mutation, mcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcu *MovieCollectionUpdate) SaveX(ctx context.Context) int {
	affected, err := mcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mcu *MovieCollectionUpdate) Exec(ctx context.Context) error {
	_, err := mcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcu *MovieCollectionUpdate) ExecX(ctx context.Context) {
	if err := mcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mcu *MovieCollectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(moviecollection.Table, moviecollection.Columns, sqlgraph.NewFieldSpec(moviecollection.FieldID, field.TypeInt))
	if ps := mcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mcu.mutation.Name(); ok {
		_spec.SetField(moviecollection.FieldName, field.TypeString, value)
	}
	if value, ok := mcu.mutation.Note(); ok {
		_spec.SetField(moviecollection.FieldNote, field.TypeString, value)
	}
	if value, ok := mcu.mutation.ShowWatched(); ok {
		_spec.SetField(moviecollection.FieldShowWatched, field.TypeBool, value)
	}
	if value, ok := mcu.mutation.CreatedAt(); ok {
		_spec.SetField(moviecollection.FieldCreatedAt, field.TypeTime, value)
	}
	if mcu.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moviecollection.MoviesTable,
			Columns: []string{moviecollection.MoviesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.RemovedMoviesIDs(); len(nodes) > 0 && !mcu.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moviecollection.MoviesTable,
			Columns: []string{moviecollection.MoviesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.MoviesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moviecollection.MoviesTable,
			Columns: []string{moviecollection.MoviesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   moviecollection.OwnerTable,
			Columns: []string{moviecollection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   moviecollection.OwnerTable,
			Columns: []string{moviecollection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcu.mutation.MovieCollectionSharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moviecollection.MovieCollectionSharesTable,
			Columns: []string{moviecollection.MovieCollectionSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(moviecollectionshare.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.RemovedMovieCollectionSharesIDs(); len(nodes) > 0 && !mcu.mutation.MovieCollectionSharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moviecollection.MovieCollectionSharesTable,
			Columns: []string{moviecollection.MovieCollectionSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(moviecollectionshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.MovieCollectionSharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moviecollection.MovieCollectionSharesTable,
			Columns: []string{moviecollection.MovieCollectionSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(moviecollectionshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{moviecollection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mcu.mutation.done = true
	return n, nil
}

// MovieCollectionUpdateOne is the builder for updating a single MovieCollection entity.
type MovieCollectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MovieCollectionMutation
}

// SetName sets the "name" field.
func (mcuo *MovieCollectionUpdateOne) SetName(s string) *MovieCollectionUpdateOne {
	mcuo.mutation.SetName(s)
	return mcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mcuo *MovieCollectionUpdateOne) SetNillableName(s *string) *MovieCollectionUpdateOne {
	if s != nil {
		mcuo.SetName(*s)
	}
	return mcuo
}

// SetNote sets the "note" field.
func (mcuo *MovieCollectionUpdateOne) SetNote(s string) *MovieCollectionUpdateOne {
	mcuo.mutation.SetNote(s)
	return mcuo
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (mcuo *MovieCollectionUpdateOne) SetNillableNote(s *string) *MovieCollectionUpdateOne {
	if s != nil {
		mcuo.SetNote(*s)
	}
	return mcuo
}

// SetShowWatched sets the "show_watched" field.
func (mcuo *MovieCollectionUpdateOne) SetShowWatched(b bool) *MovieCollectionUpdateOne {
	mcuo.mutation.SetShowWatched(b)
	return mcuo
}

// SetNillableShowWatched sets the "show_watched" field if the given value is not nil.
func (mcuo *MovieCollectionUpdateOne) SetNillableShowWatched(b *bool) *MovieCollectionUpdateOne {
	if b != nil {
		mcuo.SetShowWatched(*b)
	}
	return mcuo
}

// SetCreatedAt sets the "created_at" field.
func (mcuo *MovieCollectionUpdateOne) SetCreatedAt(t time.Time) *MovieCollectionUpdateOne {
	mcuo.mutation.SetCreatedAt(t)
	return mcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcuo *MovieCollectionUpdateOne) SetNillableCreatedAt(t *time.Time) *MovieCollectionUpdateOne {
	if t != nil {
		mcuo.SetCreatedAt(*t)
	}
	return mcuo
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (mcuo *MovieCollectionUpdateOne) AddMovieIDs(ids ...int) *MovieCollectionUpdateOne {
	mcuo.mutation.AddMovieIDs(ids...)
	return mcuo
}

// AddMovies adds the "movies" edges to the Movie entity.
func (mcuo *MovieCollectionUpdateOne) AddMovies(m ...*Movie) *MovieCollectionUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcuo.AddMovieIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (mcuo *MovieCollectionUpdateOne) SetOwnerID(id int) *MovieCollectionUpdateOne {
	mcuo.mutation.SetOwnerID(id)
	return mcuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (mcuo *MovieCollectionUpdateOne) SetNillableOwnerID(id *int) *MovieCollectionUpdateOne {
	if id != nil {
		mcuo = mcuo.SetOwnerID(*id)
	}
	return mcuo
}

// SetOwner sets the "owner" edge to the User entity.
func (mcuo *MovieCollectionUpdateOne) SetOwner(u *User) *MovieCollectionUpdateOne {
	return mcuo.SetOwnerID(u.ID)
}

// AddMovieCollectionShareIDs adds the "movie_collection_shares" edge to the MovieCollectionShare entity by IDs.
func (mcuo *MovieCollectionUpdateOne) AddMovieCollectionShareIDs(ids ...int) *MovieCollectionUpdateOne {
	mcuo.mutation.AddMovieCollectionShareIDs(ids...)
	return mcuo
}

// AddMovieCollectionShares adds the "movie_collection_shares" edges to the MovieCollectionShare entity.
func (mcuo *MovieCollectionUpdateOne) AddMovieCollectionShares(m ...*MovieCollectionShare) *MovieCollectionUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcuo.AddMovieCollectionShareIDs(ids...)
}

// Mutation returns the MovieCollectionMutation object of the builder.
func (mcuo *MovieCollectionUpdateOne) Mutation() *MovieCollectionMutation {
	return mcuo.mutation
}

// ClearMovies clears all "movies" edges to the Movie entity.
func (mcuo *MovieCollectionUpdateOne) ClearMovies() *MovieCollectionUpdateOne {
	mcuo.mutation.ClearMovies()
	return mcuo
}

// RemoveMovieIDs removes the "movies" edge to Movie entities by IDs.
func (mcuo *MovieCollectionUpdateOne) RemoveMovieIDs(ids ...int) *MovieCollectionUpdateOne {
	mcuo.mutation.RemoveMovieIDs(ids...)
	return mcuo
}

// RemoveMovies removes "movies" edges to Movie entities.
func (mcuo *MovieCollectionUpdateOne) RemoveMovies(m ...*Movie) *MovieCollectionUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcuo.RemoveMovieIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (mcuo *MovieCollectionUpdateOne) ClearOwner() *MovieCollectionUpdateOne {
	mcuo.mutation.ClearOwner()
	return mcuo
}

// ClearMovieCollectionShares clears all "movie_collection_shares" edges to the MovieCollectionShare entity.
func (mcuo *MovieCollectionUpdateOne) ClearMovieCollectionShares() *MovieCollectionUpdateOne {
	mcuo.mutation.ClearMovieCollectionShares()
	return mcuo
}

// RemoveMovieCollectionShareIDs removes the "movie_collection_shares" edge to MovieCollectionShare entities by IDs.
func (mcuo *MovieCollectionUpdateOne) RemoveMovieCollectionShareIDs(ids ...int) *MovieCollectionUpdateOne {
	mcuo.mutation.RemoveMovieCollectionShareIDs(ids...)
	return mcuo
}

// RemoveMovieCollectionShares removes "movie_collection_shares" edges to MovieCollectionShare entities.
func (mcuo *MovieCollectionUpdateOne) RemoveMovieCollectionShares(m ...*MovieCollectionShare) *MovieCollectionUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcuo.RemoveMovieCollectionShareIDs(ids...)
}

// Where appends a list predicates to the MovieCollectionUpdate builder.
func (mcuo *MovieCollectionUpdateOne) Where(ps ...predicate.MovieCollection) *MovieCollectionUpdateOne {
	mcuo.mutation.Where(ps...)
	return mcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mcuo *MovieCollectionUpdateOne) Select(field string, fields ...string) *MovieCollectionUpdateOne {
	mcuo.fields = append([]string{field}, fields...)
	return mcuo
}

// Save executes the query and returns the updated MovieCollection entity.
func (mcuo *MovieCollectionUpdateOne) Save(ctx context.Context) (*MovieCollection, error) {
	return withHooks(ctx, mcuo.sqlSave, mcuo.mutation, mcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcuo *MovieCollectionUpdateOne) SaveX(ctx context.Context) *MovieCollection {
	node, err := mcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mcuo *MovieCollectionUpdateOne) Exec(ctx context.Context) error {
	_, err := mcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcuo *MovieCollectionUpdateOne) ExecX(ctx context.Context) {
	if err := mcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mcuo *MovieCollectionUpdateOne) sqlSave(ctx context.Context) (_node *MovieCollection, err error) {
	_spec := sqlgraph.NewUpdateSpec(moviecollection.Table, moviecollection.Columns, sqlgraph.NewFieldSpec(moviecollection.FieldID, field.TypeInt))
	id, ok := mcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MovieCollection.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, moviecollection.FieldID)
		for _, f := range fields {
			if !moviecollection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != moviecollection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mcuo.mutation.Name(); ok {
		_spec.SetField(moviecollection.FieldName, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.Note(); ok {
		_spec.SetField(moviecollection.FieldNote, field.TypeString, value)
	}
	if value, ok := mcuo.mutation.ShowWatched(); ok {
		_spec.SetField(moviecollection.FieldShowWatched, field.TypeBool, value)
	}
	if value, ok := mcuo.mutation.CreatedAt(); ok {
		_spec.SetField(moviecollection.FieldCreatedAt, field.TypeTime, value)
	}
	if mcuo.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moviecollection.MoviesTable,
			Columns: []string{moviecollection.MoviesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.RemovedMoviesIDs(); len(nodes) > 0 && !mcuo.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moviecollection.MoviesTable,
			Columns: []string{moviecollection.MoviesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.MoviesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moviecollection.MoviesTable,
			Columns: []string{moviecollection.MoviesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   moviecollection.OwnerTable,
			Columns: []string{moviecollection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   moviecollection.OwnerTable,
			Columns: []string{moviecollection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcuo.mutation.MovieCollectionSharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moviecollection.MovieCollectionSharesTable,
			Columns: []string{moviecollection.MovieCollectionSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(moviecollectionshare.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.RemovedMovieCollectionSharesIDs(); len(nodes) > 0 && !mcuo.mutation.MovieCollectionSharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moviecollection.MovieCollectionSharesTable,
			Columns: []string{moviecollection.MovieCollectionSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(moviecollectionshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.MovieCollectionSharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moviecollection.MovieCollectionSharesTable,
			Columns: []string{moviecollection.MovieCollectionSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(moviecollectionshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MovieCollection{config: mcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{moviecollection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mcuo.mutation.done = true
	return _node, nil
}
