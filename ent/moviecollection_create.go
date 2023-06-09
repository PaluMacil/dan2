// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PaluMacil/dan2/ent/movie"
	"github.com/PaluMacil/dan2/ent/moviecollection"
	"github.com/PaluMacil/dan2/ent/moviecollectionshare"
	"github.com/PaluMacil/dan2/ent/user"
)

// MovieCollectionCreate is the builder for creating a MovieCollection entity.
type MovieCollectionCreate struct {
	config
	mutation *MovieCollectionMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (mcc *MovieCollectionCreate) SetName(s string) *MovieCollectionCreate {
	mcc.mutation.SetName(s)
	return mcc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mcc *MovieCollectionCreate) SetNillableName(s *string) *MovieCollectionCreate {
	if s != nil {
		mcc.SetName(*s)
	}
	return mcc
}

// SetNote sets the "note" field.
func (mcc *MovieCollectionCreate) SetNote(s string) *MovieCollectionCreate {
	mcc.mutation.SetNote(s)
	return mcc
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (mcc *MovieCollectionCreate) SetNillableNote(s *string) *MovieCollectionCreate {
	if s != nil {
		mcc.SetNote(*s)
	}
	return mcc
}

// SetShowWatched sets the "show_watched" field.
func (mcc *MovieCollectionCreate) SetShowWatched(b bool) *MovieCollectionCreate {
	mcc.mutation.SetShowWatched(b)
	return mcc
}

// SetNillableShowWatched sets the "show_watched" field if the given value is not nil.
func (mcc *MovieCollectionCreate) SetNillableShowWatched(b *bool) *MovieCollectionCreate {
	if b != nil {
		mcc.SetShowWatched(*b)
	}
	return mcc
}

// SetCreatedAt sets the "created_at" field.
func (mcc *MovieCollectionCreate) SetCreatedAt(t time.Time) *MovieCollectionCreate {
	mcc.mutation.SetCreatedAt(t)
	return mcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcc *MovieCollectionCreate) SetNillableCreatedAt(t *time.Time) *MovieCollectionCreate {
	if t != nil {
		mcc.SetCreatedAt(*t)
	}
	return mcc
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (mcc *MovieCollectionCreate) AddMovieIDs(ids ...int) *MovieCollectionCreate {
	mcc.mutation.AddMovieIDs(ids...)
	return mcc
}

// AddMovies adds the "movies" edges to the Movie entity.
func (mcc *MovieCollectionCreate) AddMovies(m ...*Movie) *MovieCollectionCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcc.AddMovieIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (mcc *MovieCollectionCreate) SetOwnerID(id int) *MovieCollectionCreate {
	mcc.mutation.SetOwnerID(id)
	return mcc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (mcc *MovieCollectionCreate) SetNillableOwnerID(id *int) *MovieCollectionCreate {
	if id != nil {
		mcc = mcc.SetOwnerID(*id)
	}
	return mcc
}

// SetOwner sets the "owner" edge to the User entity.
func (mcc *MovieCollectionCreate) SetOwner(u *User) *MovieCollectionCreate {
	return mcc.SetOwnerID(u.ID)
}

// AddMovieCollectionShareIDs adds the "movie_collection_shares" edge to the MovieCollectionShare entity by IDs.
func (mcc *MovieCollectionCreate) AddMovieCollectionShareIDs(ids ...int) *MovieCollectionCreate {
	mcc.mutation.AddMovieCollectionShareIDs(ids...)
	return mcc
}

// AddMovieCollectionShares adds the "movie_collection_shares" edges to the MovieCollectionShare entity.
func (mcc *MovieCollectionCreate) AddMovieCollectionShares(m ...*MovieCollectionShare) *MovieCollectionCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcc.AddMovieCollectionShareIDs(ids...)
}

// Mutation returns the MovieCollectionMutation object of the builder.
func (mcc *MovieCollectionCreate) Mutation() *MovieCollectionMutation {
	return mcc.mutation
}

// Save creates the MovieCollection in the database.
func (mcc *MovieCollectionCreate) Save(ctx context.Context) (*MovieCollection, error) {
	mcc.defaults()
	return withHooks(ctx, mcc.sqlSave, mcc.mutation, mcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mcc *MovieCollectionCreate) SaveX(ctx context.Context) *MovieCollection {
	v, err := mcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcc *MovieCollectionCreate) Exec(ctx context.Context) error {
	_, err := mcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcc *MovieCollectionCreate) ExecX(ctx context.Context) {
	if err := mcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcc *MovieCollectionCreate) defaults() {
	if _, ok := mcc.mutation.Name(); !ok {
		v := moviecollection.DefaultName
		mcc.mutation.SetName(v)
	}
	if _, ok := mcc.mutation.Note(); !ok {
		v := moviecollection.DefaultNote
		mcc.mutation.SetNote(v)
	}
	if _, ok := mcc.mutation.ShowWatched(); !ok {
		v := moviecollection.DefaultShowWatched
		mcc.mutation.SetShowWatched(v)
	}
	if _, ok := mcc.mutation.CreatedAt(); !ok {
		v := moviecollection.DefaultCreatedAt()
		mcc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mcc *MovieCollectionCreate) check() error {
	if _, ok := mcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MovieCollection.name"`)}
	}
	if _, ok := mcc.mutation.Note(); !ok {
		return &ValidationError{Name: "note", err: errors.New(`ent: missing required field "MovieCollection.note"`)}
	}
	if _, ok := mcc.mutation.ShowWatched(); !ok {
		return &ValidationError{Name: "show_watched", err: errors.New(`ent: missing required field "MovieCollection.show_watched"`)}
	}
	if _, ok := mcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MovieCollection.created_at"`)}
	}
	return nil
}

func (mcc *MovieCollectionCreate) sqlSave(ctx context.Context) (*MovieCollection, error) {
	if err := mcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mcc.mutation.id = &_node.ID
	mcc.mutation.done = true
	return _node, nil
}

func (mcc *MovieCollectionCreate) createSpec() (*MovieCollection, *sqlgraph.CreateSpec) {
	var (
		_node = &MovieCollection{config: mcc.config}
		_spec = sqlgraph.NewCreateSpec(moviecollection.Table, sqlgraph.NewFieldSpec(moviecollection.FieldID, field.TypeInt))
	)
	if value, ok := mcc.mutation.Name(); ok {
		_spec.SetField(moviecollection.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mcc.mutation.Note(); ok {
		_spec.SetField(moviecollection.FieldNote, field.TypeString, value)
		_node.Note = value
	}
	if value, ok := mcc.mutation.ShowWatched(); ok {
		_spec.SetField(moviecollection.FieldShowWatched, field.TypeBool, value)
		_node.ShowWatched = value
	}
	if value, ok := mcc.mutation.CreatedAt(); ok {
		_spec.SetField(moviecollection.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := mcc.mutation.MoviesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mcc.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_node.user_movie_collections = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mcc.mutation.MovieCollectionSharesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MovieCollectionCreateBulk is the builder for creating many MovieCollection entities in bulk.
type MovieCollectionCreateBulk struct {
	config
	builders []*MovieCollectionCreate
}

// Save creates the MovieCollection entities in the database.
func (mccb *MovieCollectionCreateBulk) Save(ctx context.Context) ([]*MovieCollection, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mccb.builders))
	nodes := make([]*MovieCollection, len(mccb.builders))
	mutators := make([]Mutator, len(mccb.builders))
	for i := range mccb.builders {
		func(i int, root context.Context) {
			builder := mccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MovieCollectionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mccb *MovieCollectionCreateBulk) SaveX(ctx context.Context) []*MovieCollection {
	v, err := mccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mccb *MovieCollectionCreateBulk) Exec(ctx context.Context) error {
	_, err := mccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mccb *MovieCollectionCreateBulk) ExecX(ctx context.Context) {
	if err := mccb.Exec(ctx); err != nil {
		panic(err)
	}
}
