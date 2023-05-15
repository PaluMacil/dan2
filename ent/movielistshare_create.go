// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PaluMacil/dan2/ent/movielist"
	"github.com/PaluMacil/dan2/ent/movielistshare"
	"github.com/PaluMacil/dan2/ent/user"
)

// MovieListShareCreate is the builder for creating a MovieListShare entity.
type MovieListShareCreate struct {
	config
	mutation *MovieListShareMutation
	hooks    []Hook
}

// SetCanEdit sets the "can_edit" field.
func (mlsc *MovieListShareCreate) SetCanEdit(b bool) *MovieListShareCreate {
	mlsc.mutation.SetCanEdit(b)
	return mlsc
}

// SetNillableCanEdit sets the "can_edit" field if the given value is not nil.
func (mlsc *MovieListShareCreate) SetNillableCanEdit(b *bool) *MovieListShareCreate {
	if b != nil {
		mlsc.SetCanEdit(*b)
	}
	return mlsc
}

// SetCreatedAt sets the "created_at" field.
func (mlsc *MovieListShareCreate) SetCreatedAt(t time.Time) *MovieListShareCreate {
	mlsc.mutation.SetCreatedAt(t)
	return mlsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mlsc *MovieListShareCreate) SetNillableCreatedAt(t *time.Time) *MovieListShareCreate {
	if t != nil {
		mlsc.SetCreatedAt(*t)
	}
	return mlsc
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (mlsc *MovieListShareCreate) AddUserIDs(ids ...int) *MovieListShareCreate {
	mlsc.mutation.AddUserIDs(ids...)
	return mlsc
}

// AddUser adds the "user" edges to the User entity.
func (mlsc *MovieListShareCreate) AddUser(u ...*User) *MovieListShareCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mlsc.AddUserIDs(ids...)
}

// AddMovieListIDs adds the "movie_list" edge to the MovieList entity by IDs.
func (mlsc *MovieListShareCreate) AddMovieListIDs(ids ...int) *MovieListShareCreate {
	mlsc.mutation.AddMovieListIDs(ids...)
	return mlsc
}

// AddMovieList adds the "movie_list" edges to the MovieList entity.
func (mlsc *MovieListShareCreate) AddMovieList(m ...*MovieList) *MovieListShareCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mlsc.AddMovieListIDs(ids...)
}

// Mutation returns the MovieListShareMutation object of the builder.
func (mlsc *MovieListShareCreate) Mutation() *MovieListShareMutation {
	return mlsc.mutation
}

// Save creates the MovieListShare in the database.
func (mlsc *MovieListShareCreate) Save(ctx context.Context) (*MovieListShare, error) {
	mlsc.defaults()
	return withHooks(ctx, mlsc.sqlSave, mlsc.mutation, mlsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mlsc *MovieListShareCreate) SaveX(ctx context.Context) *MovieListShare {
	v, err := mlsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mlsc *MovieListShareCreate) Exec(ctx context.Context) error {
	_, err := mlsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlsc *MovieListShareCreate) ExecX(ctx context.Context) {
	if err := mlsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mlsc *MovieListShareCreate) defaults() {
	if _, ok := mlsc.mutation.CanEdit(); !ok {
		v := movielistshare.DefaultCanEdit
		mlsc.mutation.SetCanEdit(v)
	}
	if _, ok := mlsc.mutation.CreatedAt(); !ok {
		v := movielistshare.DefaultCreatedAt()
		mlsc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mlsc *MovieListShareCreate) check() error {
	if _, ok := mlsc.mutation.CanEdit(); !ok {
		return &ValidationError{Name: "can_edit", err: errors.New(`ent: missing required field "MovieListShare.can_edit"`)}
	}
	if _, ok := mlsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MovieListShare.created_at"`)}
	}
	return nil
}

func (mlsc *MovieListShareCreate) sqlSave(ctx context.Context) (*MovieListShare, error) {
	if err := mlsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mlsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mlsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mlsc.mutation.id = &_node.ID
	mlsc.mutation.done = true
	return _node, nil
}

func (mlsc *MovieListShareCreate) createSpec() (*MovieListShare, *sqlgraph.CreateSpec) {
	var (
		_node = &MovieListShare{config: mlsc.config}
		_spec = sqlgraph.NewCreateSpec(movielistshare.Table, sqlgraph.NewFieldSpec(movielistshare.FieldID, field.TypeInt))
	)
	if value, ok := mlsc.mutation.CanEdit(); ok {
		_spec.SetField(movielistshare.FieldCanEdit, field.TypeBool, value)
		_node.CanEdit = value
	}
	if value, ok := mlsc.mutation.CreatedAt(); ok {
		_spec.SetField(movielistshare.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := mlsc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   movielistshare.UserTable,
			Columns: movielistshare.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mlsc.mutation.MovieListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   movielistshare.MovieListTable,
			Columns: movielistshare.MovieListPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movielist.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MovieListShareCreateBulk is the builder for creating many MovieListShare entities in bulk.
type MovieListShareCreateBulk struct {
	config
	builders []*MovieListShareCreate
}

// Save creates the MovieListShare entities in the database.
func (mlscb *MovieListShareCreateBulk) Save(ctx context.Context) ([]*MovieListShare, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mlscb.builders))
	nodes := make([]*MovieListShare, len(mlscb.builders))
	mutators := make([]Mutator, len(mlscb.builders))
	for i := range mlscb.builders {
		func(i int, root context.Context) {
			builder := mlscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MovieListShareMutation)
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
					_, err = mutators[i+1].Mutate(root, mlscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mlscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mlscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mlscb *MovieListShareCreateBulk) SaveX(ctx context.Context) []*MovieListShare {
	v, err := mlscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mlscb *MovieListShareCreateBulk) Exec(ctx context.Context) error {
	_, err := mlscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlscb *MovieListShareCreateBulk) ExecX(ctx context.Context) {
	if err := mlscb.Exec(ctx); err != nil {
		panic(err)
	}
}