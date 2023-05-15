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
	"github.com/PaluMacil/dan2/ent/movielist"
	"github.com/PaluMacil/dan2/ent/movielistshare"
	"github.com/PaluMacil/dan2/ent/predicate"
	"github.com/PaluMacil/dan2/ent/user"
)

// MovieListUpdate is the builder for updating MovieList entities.
type MovieListUpdate struct {
	config
	hooks    []Hook
	mutation *MovieListMutation
}

// Where appends a list predicates to the MovieListUpdate builder.
func (mlu *MovieListUpdate) Where(ps ...predicate.MovieList) *MovieListUpdate {
	mlu.mutation.Where(ps...)
	return mlu
}

// SetName sets the "name" field.
func (mlu *MovieListUpdate) SetName(s string) *MovieListUpdate {
	mlu.mutation.SetName(s)
	return mlu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mlu *MovieListUpdate) SetNillableName(s *string) *MovieListUpdate {
	if s != nil {
		mlu.SetName(*s)
	}
	return mlu
}

// SetNote sets the "note" field.
func (mlu *MovieListUpdate) SetNote(s string) *MovieListUpdate {
	mlu.mutation.SetNote(s)
	return mlu
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (mlu *MovieListUpdate) SetNillableNote(s *string) *MovieListUpdate {
	if s != nil {
		mlu.SetNote(*s)
	}
	return mlu
}

// SetShowWatched sets the "show_watched" field.
func (mlu *MovieListUpdate) SetShowWatched(b bool) *MovieListUpdate {
	mlu.mutation.SetShowWatched(b)
	return mlu
}

// SetNillableShowWatched sets the "show_watched" field if the given value is not nil.
func (mlu *MovieListUpdate) SetNillableShowWatched(b *bool) *MovieListUpdate {
	if b != nil {
		mlu.SetShowWatched(*b)
	}
	return mlu
}

// SetCreatedAt sets the "created_at" field.
func (mlu *MovieListUpdate) SetCreatedAt(t time.Time) *MovieListUpdate {
	mlu.mutation.SetCreatedAt(t)
	return mlu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mlu *MovieListUpdate) SetNillableCreatedAt(t *time.Time) *MovieListUpdate {
	if t != nil {
		mlu.SetCreatedAt(*t)
	}
	return mlu
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (mlu *MovieListUpdate) AddMovieIDs(ids ...int) *MovieListUpdate {
	mlu.mutation.AddMovieIDs(ids...)
	return mlu
}

// AddMovies adds the "movies" edges to the Movie entity.
func (mlu *MovieListUpdate) AddMovies(m ...*Movie) *MovieListUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mlu.AddMovieIDs(ids...)
}

// AddOwnerIDs adds the "owner" edge to the User entity by IDs.
func (mlu *MovieListUpdate) AddOwnerIDs(ids ...int) *MovieListUpdate {
	mlu.mutation.AddOwnerIDs(ids...)
	return mlu
}

// AddOwner adds the "owner" edges to the User entity.
func (mlu *MovieListUpdate) AddOwner(u ...*User) *MovieListUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mlu.AddOwnerIDs(ids...)
}

// AddMovieListShareIDs adds the "movie_list_shares" edge to the MovieListShare entity by IDs.
func (mlu *MovieListUpdate) AddMovieListShareIDs(ids ...int) *MovieListUpdate {
	mlu.mutation.AddMovieListShareIDs(ids...)
	return mlu
}

// AddMovieListShares adds the "movie_list_shares" edges to the MovieListShare entity.
func (mlu *MovieListUpdate) AddMovieListShares(m ...*MovieListShare) *MovieListUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mlu.AddMovieListShareIDs(ids...)
}

// Mutation returns the MovieListMutation object of the builder.
func (mlu *MovieListUpdate) Mutation() *MovieListMutation {
	return mlu.mutation
}

// ClearMovies clears all "movies" edges to the Movie entity.
func (mlu *MovieListUpdate) ClearMovies() *MovieListUpdate {
	mlu.mutation.ClearMovies()
	return mlu
}

// RemoveMovieIDs removes the "movies" edge to Movie entities by IDs.
func (mlu *MovieListUpdate) RemoveMovieIDs(ids ...int) *MovieListUpdate {
	mlu.mutation.RemoveMovieIDs(ids...)
	return mlu
}

// RemoveMovies removes "movies" edges to Movie entities.
func (mlu *MovieListUpdate) RemoveMovies(m ...*Movie) *MovieListUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mlu.RemoveMovieIDs(ids...)
}

// ClearOwner clears all "owner" edges to the User entity.
func (mlu *MovieListUpdate) ClearOwner() *MovieListUpdate {
	mlu.mutation.ClearOwner()
	return mlu
}

// RemoveOwnerIDs removes the "owner" edge to User entities by IDs.
func (mlu *MovieListUpdate) RemoveOwnerIDs(ids ...int) *MovieListUpdate {
	mlu.mutation.RemoveOwnerIDs(ids...)
	return mlu
}

// RemoveOwner removes "owner" edges to User entities.
func (mlu *MovieListUpdate) RemoveOwner(u ...*User) *MovieListUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mlu.RemoveOwnerIDs(ids...)
}

// ClearMovieListShares clears all "movie_list_shares" edges to the MovieListShare entity.
func (mlu *MovieListUpdate) ClearMovieListShares() *MovieListUpdate {
	mlu.mutation.ClearMovieListShares()
	return mlu
}

// RemoveMovieListShareIDs removes the "movie_list_shares" edge to MovieListShare entities by IDs.
func (mlu *MovieListUpdate) RemoveMovieListShareIDs(ids ...int) *MovieListUpdate {
	mlu.mutation.RemoveMovieListShareIDs(ids...)
	return mlu
}

// RemoveMovieListShares removes "movie_list_shares" edges to MovieListShare entities.
func (mlu *MovieListUpdate) RemoveMovieListShares(m ...*MovieListShare) *MovieListUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mlu.RemoveMovieListShareIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mlu *MovieListUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mlu.sqlSave, mlu.mutation, mlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mlu *MovieListUpdate) SaveX(ctx context.Context) int {
	affected, err := mlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mlu *MovieListUpdate) Exec(ctx context.Context) error {
	_, err := mlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlu *MovieListUpdate) ExecX(ctx context.Context) {
	if err := mlu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mlu *MovieListUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(movielist.Table, movielist.Columns, sqlgraph.NewFieldSpec(movielist.FieldID, field.TypeInt))
	if ps := mlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mlu.mutation.Name(); ok {
		_spec.SetField(movielist.FieldName, field.TypeString, value)
	}
	if value, ok := mlu.mutation.Note(); ok {
		_spec.SetField(movielist.FieldNote, field.TypeString, value)
	}
	if value, ok := mlu.mutation.ShowWatched(); ok {
		_spec.SetField(movielist.FieldShowWatched, field.TypeBool, value)
	}
	if value, ok := mlu.mutation.CreatedAt(); ok {
		_spec.SetField(movielist.FieldCreatedAt, field.TypeTime, value)
	}
	if mlu.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   movielist.MoviesTable,
			Columns: movielist.MoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mlu.mutation.RemovedMoviesIDs(); len(nodes) > 0 && !mlu.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   movielist.MoviesTable,
			Columns: movielist.MoviesPrimaryKey,
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
	if nodes := mlu.mutation.MoviesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   movielist.MoviesTable,
			Columns: movielist.MoviesPrimaryKey,
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
	if mlu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   movielist.OwnerTable,
			Columns: movielist.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mlu.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !mlu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   movielist.OwnerTable,
			Columns: movielist.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mlu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   movielist.OwnerTable,
			Columns: movielist.OwnerPrimaryKey,
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
	if mlu.mutation.MovieListSharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   movielist.MovieListSharesTable,
			Columns: movielist.MovieListSharesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movielistshare.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mlu.mutation.RemovedMovieListSharesIDs(); len(nodes) > 0 && !mlu.mutation.MovieListSharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   movielist.MovieListSharesTable,
			Columns: movielist.MovieListSharesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movielistshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mlu.mutation.MovieListSharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   movielist.MovieListSharesTable,
			Columns: movielist.MovieListSharesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movielistshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{movielist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mlu.mutation.done = true
	return n, nil
}

// MovieListUpdateOne is the builder for updating a single MovieList entity.
type MovieListUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MovieListMutation
}

// SetName sets the "name" field.
func (mluo *MovieListUpdateOne) SetName(s string) *MovieListUpdateOne {
	mluo.mutation.SetName(s)
	return mluo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mluo *MovieListUpdateOne) SetNillableName(s *string) *MovieListUpdateOne {
	if s != nil {
		mluo.SetName(*s)
	}
	return mluo
}

// SetNote sets the "note" field.
func (mluo *MovieListUpdateOne) SetNote(s string) *MovieListUpdateOne {
	mluo.mutation.SetNote(s)
	return mluo
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (mluo *MovieListUpdateOne) SetNillableNote(s *string) *MovieListUpdateOne {
	if s != nil {
		mluo.SetNote(*s)
	}
	return mluo
}

// SetShowWatched sets the "show_watched" field.
func (mluo *MovieListUpdateOne) SetShowWatched(b bool) *MovieListUpdateOne {
	mluo.mutation.SetShowWatched(b)
	return mluo
}

// SetNillableShowWatched sets the "show_watched" field if the given value is not nil.
func (mluo *MovieListUpdateOne) SetNillableShowWatched(b *bool) *MovieListUpdateOne {
	if b != nil {
		mluo.SetShowWatched(*b)
	}
	return mluo
}

// SetCreatedAt sets the "created_at" field.
func (mluo *MovieListUpdateOne) SetCreatedAt(t time.Time) *MovieListUpdateOne {
	mluo.mutation.SetCreatedAt(t)
	return mluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mluo *MovieListUpdateOne) SetNillableCreatedAt(t *time.Time) *MovieListUpdateOne {
	if t != nil {
		mluo.SetCreatedAt(*t)
	}
	return mluo
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (mluo *MovieListUpdateOne) AddMovieIDs(ids ...int) *MovieListUpdateOne {
	mluo.mutation.AddMovieIDs(ids...)
	return mluo
}

// AddMovies adds the "movies" edges to the Movie entity.
func (mluo *MovieListUpdateOne) AddMovies(m ...*Movie) *MovieListUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mluo.AddMovieIDs(ids...)
}

// AddOwnerIDs adds the "owner" edge to the User entity by IDs.
func (mluo *MovieListUpdateOne) AddOwnerIDs(ids ...int) *MovieListUpdateOne {
	mluo.mutation.AddOwnerIDs(ids...)
	return mluo
}

// AddOwner adds the "owner" edges to the User entity.
func (mluo *MovieListUpdateOne) AddOwner(u ...*User) *MovieListUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mluo.AddOwnerIDs(ids...)
}

// AddMovieListShareIDs adds the "movie_list_shares" edge to the MovieListShare entity by IDs.
func (mluo *MovieListUpdateOne) AddMovieListShareIDs(ids ...int) *MovieListUpdateOne {
	mluo.mutation.AddMovieListShareIDs(ids...)
	return mluo
}

// AddMovieListShares adds the "movie_list_shares" edges to the MovieListShare entity.
func (mluo *MovieListUpdateOne) AddMovieListShares(m ...*MovieListShare) *MovieListUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mluo.AddMovieListShareIDs(ids...)
}

// Mutation returns the MovieListMutation object of the builder.
func (mluo *MovieListUpdateOne) Mutation() *MovieListMutation {
	return mluo.mutation
}

// ClearMovies clears all "movies" edges to the Movie entity.
func (mluo *MovieListUpdateOne) ClearMovies() *MovieListUpdateOne {
	mluo.mutation.ClearMovies()
	return mluo
}

// RemoveMovieIDs removes the "movies" edge to Movie entities by IDs.
func (mluo *MovieListUpdateOne) RemoveMovieIDs(ids ...int) *MovieListUpdateOne {
	mluo.mutation.RemoveMovieIDs(ids...)
	return mluo
}

// RemoveMovies removes "movies" edges to Movie entities.
func (mluo *MovieListUpdateOne) RemoveMovies(m ...*Movie) *MovieListUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mluo.RemoveMovieIDs(ids...)
}

// ClearOwner clears all "owner" edges to the User entity.
func (mluo *MovieListUpdateOne) ClearOwner() *MovieListUpdateOne {
	mluo.mutation.ClearOwner()
	return mluo
}

// RemoveOwnerIDs removes the "owner" edge to User entities by IDs.
func (mluo *MovieListUpdateOne) RemoveOwnerIDs(ids ...int) *MovieListUpdateOne {
	mluo.mutation.RemoveOwnerIDs(ids...)
	return mluo
}

// RemoveOwner removes "owner" edges to User entities.
func (mluo *MovieListUpdateOne) RemoveOwner(u ...*User) *MovieListUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mluo.RemoveOwnerIDs(ids...)
}

// ClearMovieListShares clears all "movie_list_shares" edges to the MovieListShare entity.
func (mluo *MovieListUpdateOne) ClearMovieListShares() *MovieListUpdateOne {
	mluo.mutation.ClearMovieListShares()
	return mluo
}

// RemoveMovieListShareIDs removes the "movie_list_shares" edge to MovieListShare entities by IDs.
func (mluo *MovieListUpdateOne) RemoveMovieListShareIDs(ids ...int) *MovieListUpdateOne {
	mluo.mutation.RemoveMovieListShareIDs(ids...)
	return mluo
}

// RemoveMovieListShares removes "movie_list_shares" edges to MovieListShare entities.
func (mluo *MovieListUpdateOne) RemoveMovieListShares(m ...*MovieListShare) *MovieListUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mluo.RemoveMovieListShareIDs(ids...)
}

// Where appends a list predicates to the MovieListUpdate builder.
func (mluo *MovieListUpdateOne) Where(ps ...predicate.MovieList) *MovieListUpdateOne {
	mluo.mutation.Where(ps...)
	return mluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mluo *MovieListUpdateOne) Select(field string, fields ...string) *MovieListUpdateOne {
	mluo.fields = append([]string{field}, fields...)
	return mluo
}

// Save executes the query and returns the updated MovieList entity.
func (mluo *MovieListUpdateOne) Save(ctx context.Context) (*MovieList, error) {
	return withHooks(ctx, mluo.sqlSave, mluo.mutation, mluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mluo *MovieListUpdateOne) SaveX(ctx context.Context) *MovieList {
	node, err := mluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mluo *MovieListUpdateOne) Exec(ctx context.Context) error {
	_, err := mluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mluo *MovieListUpdateOne) ExecX(ctx context.Context) {
	if err := mluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mluo *MovieListUpdateOne) sqlSave(ctx context.Context) (_node *MovieList, err error) {
	_spec := sqlgraph.NewUpdateSpec(movielist.Table, movielist.Columns, sqlgraph.NewFieldSpec(movielist.FieldID, field.TypeInt))
	id, ok := mluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MovieList.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, movielist.FieldID)
		for _, f := range fields {
			if !movielist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != movielist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mluo.mutation.Name(); ok {
		_spec.SetField(movielist.FieldName, field.TypeString, value)
	}
	if value, ok := mluo.mutation.Note(); ok {
		_spec.SetField(movielist.FieldNote, field.TypeString, value)
	}
	if value, ok := mluo.mutation.ShowWatched(); ok {
		_spec.SetField(movielist.FieldShowWatched, field.TypeBool, value)
	}
	if value, ok := mluo.mutation.CreatedAt(); ok {
		_spec.SetField(movielist.FieldCreatedAt, field.TypeTime, value)
	}
	if mluo.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   movielist.MoviesTable,
			Columns: movielist.MoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mluo.mutation.RemovedMoviesIDs(); len(nodes) > 0 && !mluo.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   movielist.MoviesTable,
			Columns: movielist.MoviesPrimaryKey,
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
	if nodes := mluo.mutation.MoviesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   movielist.MoviesTable,
			Columns: movielist.MoviesPrimaryKey,
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
	if mluo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   movielist.OwnerTable,
			Columns: movielist.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mluo.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !mluo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   movielist.OwnerTable,
			Columns: movielist.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mluo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   movielist.OwnerTable,
			Columns: movielist.OwnerPrimaryKey,
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
	if mluo.mutation.MovieListSharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   movielist.MovieListSharesTable,
			Columns: movielist.MovieListSharesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movielistshare.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mluo.mutation.RemovedMovieListSharesIDs(); len(nodes) > 0 && !mluo.mutation.MovieListSharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   movielist.MovieListSharesTable,
			Columns: movielist.MovieListSharesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movielistshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mluo.mutation.MovieListSharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   movielist.MovieListSharesTable,
			Columns: movielist.MovieListSharesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movielistshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MovieList{config: mluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{movielist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mluo.mutation.done = true
	return _node, nil
}