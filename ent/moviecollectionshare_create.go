// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PaluMacil/dan2/ent/moviecollection"
	"github.com/PaluMacil/dan2/ent/moviecollectionshare"
	"github.com/PaluMacil/dan2/ent/user"
)

// MovieCollectionShareCreate is the builder for creating a MovieCollectionShare entity.
type MovieCollectionShareCreate struct {
	config
	mutation *MovieCollectionShareMutation
	hooks    []Hook
}

// SetCanEdit sets the "can_edit" field.
func (mcsc *MovieCollectionShareCreate) SetCanEdit(b bool) *MovieCollectionShareCreate {
	mcsc.mutation.SetCanEdit(b)
	return mcsc
}

// SetNillableCanEdit sets the "can_edit" field if the given value is not nil.
func (mcsc *MovieCollectionShareCreate) SetNillableCanEdit(b *bool) *MovieCollectionShareCreate {
	if b != nil {
		mcsc.SetCanEdit(*b)
	}
	return mcsc
}

// SetCreatedAt sets the "created_at" field.
func (mcsc *MovieCollectionShareCreate) SetCreatedAt(t time.Time) *MovieCollectionShareCreate {
	mcsc.mutation.SetCreatedAt(t)
	return mcsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcsc *MovieCollectionShareCreate) SetNillableCreatedAt(t *time.Time) *MovieCollectionShareCreate {
	if t != nil {
		mcsc.SetCreatedAt(*t)
	}
	return mcsc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mcsc *MovieCollectionShareCreate) SetUserID(id int) *MovieCollectionShareCreate {
	mcsc.mutation.SetUserID(id)
	return mcsc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (mcsc *MovieCollectionShareCreate) SetNillableUserID(id *int) *MovieCollectionShareCreate {
	if id != nil {
		mcsc = mcsc.SetUserID(*id)
	}
	return mcsc
}

// SetUser sets the "user" edge to the User entity.
func (mcsc *MovieCollectionShareCreate) SetUser(u *User) *MovieCollectionShareCreate {
	return mcsc.SetUserID(u.ID)
}

// SetMovieCollectionID sets the "movie_collection" edge to the MovieCollection entity by ID.
func (mcsc *MovieCollectionShareCreate) SetMovieCollectionID(id int) *MovieCollectionShareCreate {
	mcsc.mutation.SetMovieCollectionID(id)
	return mcsc
}

// SetNillableMovieCollectionID sets the "movie_collection" edge to the MovieCollection entity by ID if the given value is not nil.
func (mcsc *MovieCollectionShareCreate) SetNillableMovieCollectionID(id *int) *MovieCollectionShareCreate {
	if id != nil {
		mcsc = mcsc.SetMovieCollectionID(*id)
	}
	return mcsc
}

// SetMovieCollection sets the "movie_collection" edge to the MovieCollection entity.
func (mcsc *MovieCollectionShareCreate) SetMovieCollection(m *MovieCollection) *MovieCollectionShareCreate {
	return mcsc.SetMovieCollectionID(m.ID)
}

// Mutation returns the MovieCollectionShareMutation object of the builder.
func (mcsc *MovieCollectionShareCreate) Mutation() *MovieCollectionShareMutation {
	return mcsc.mutation
}

// Save creates the MovieCollectionShare in the database.
func (mcsc *MovieCollectionShareCreate) Save(ctx context.Context) (*MovieCollectionShare, error) {
	mcsc.defaults()
	return withHooks(ctx, mcsc.sqlSave, mcsc.mutation, mcsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mcsc *MovieCollectionShareCreate) SaveX(ctx context.Context) *MovieCollectionShare {
	v, err := mcsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcsc *MovieCollectionShareCreate) Exec(ctx context.Context) error {
	_, err := mcsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcsc *MovieCollectionShareCreate) ExecX(ctx context.Context) {
	if err := mcsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcsc *MovieCollectionShareCreate) defaults() {
	if _, ok := mcsc.mutation.CanEdit(); !ok {
		v := moviecollectionshare.DefaultCanEdit
		mcsc.mutation.SetCanEdit(v)
	}
	if _, ok := mcsc.mutation.CreatedAt(); !ok {
		v := moviecollectionshare.DefaultCreatedAt()
		mcsc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mcsc *MovieCollectionShareCreate) check() error {
	if _, ok := mcsc.mutation.CanEdit(); !ok {
		return &ValidationError{Name: "can_edit", err: errors.New(`ent: missing required field "MovieCollectionShare.can_edit"`)}
	}
	if _, ok := mcsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MovieCollectionShare.created_at"`)}
	}
	return nil
}

func (mcsc *MovieCollectionShareCreate) sqlSave(ctx context.Context) (*MovieCollectionShare, error) {
	if err := mcsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mcsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mcsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mcsc.mutation.id = &_node.ID
	mcsc.mutation.done = true
	return _node, nil
}

func (mcsc *MovieCollectionShareCreate) createSpec() (*MovieCollectionShare, *sqlgraph.CreateSpec) {
	var (
		_node = &MovieCollectionShare{config: mcsc.config}
		_spec = sqlgraph.NewCreateSpec(moviecollectionshare.Table, sqlgraph.NewFieldSpec(moviecollectionshare.FieldID, field.TypeInt))
	)
	if value, ok := mcsc.mutation.CanEdit(); ok {
		_spec.SetField(moviecollectionshare.FieldCanEdit, field.TypeBool, value)
		_node.CanEdit = value
	}
	if value, ok := mcsc.mutation.CreatedAt(); ok {
		_spec.SetField(moviecollectionshare.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := mcsc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   moviecollectionshare.UserTable,
			Columns: []string{moviecollectionshare.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_movie_collection_shares = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mcsc.mutation.MovieCollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   moviecollectionshare.MovieCollectionTable,
			Columns: []string{moviecollectionshare.MovieCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(moviecollection.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.movie_collection_movie_collection_shares = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MovieCollectionShareCreateBulk is the builder for creating many MovieCollectionShare entities in bulk.
type MovieCollectionShareCreateBulk struct {
	config
	builders []*MovieCollectionShareCreate
}

// Save creates the MovieCollectionShare entities in the database.
func (mcscb *MovieCollectionShareCreateBulk) Save(ctx context.Context) ([]*MovieCollectionShare, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcscb.builders))
	nodes := make([]*MovieCollectionShare, len(mcscb.builders))
	mutators := make([]Mutator, len(mcscb.builders))
	for i := range mcscb.builders {
		func(i int, root context.Context) {
			builder := mcscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MovieCollectionShareMutation)
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
					_, err = mutators[i+1].Mutate(root, mcscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcscb *MovieCollectionShareCreateBulk) SaveX(ctx context.Context) []*MovieCollectionShare {
	v, err := mcscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcscb *MovieCollectionShareCreateBulk) Exec(ctx context.Context) error {
	_, err := mcscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcscb *MovieCollectionShareCreateBulk) ExecX(ctx context.Context) {
	if err := mcscb.Exec(ctx); err != nil {
		panic(err)
	}
}
