// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PaluMacil/dan2/ent/amazonlist"
	"github.com/PaluMacil/dan2/ent/amazonshare"
	"github.com/PaluMacil/dan2/ent/drink"
	"github.com/PaluMacil/dan2/ent/grocerylist"
	"github.com/PaluMacil/dan2/ent/grocerylistshare"
	"github.com/PaluMacil/dan2/ent/movielist"
	"github.com/PaluMacil/dan2/ent/movielistshare"
	"github.com/PaluMacil/dan2/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetDisplayName sets the "display_name" field.
func (uc *UserCreate) SetDisplayName(s string) *UserCreate {
	uc.mutation.SetDisplayName(s)
	return uc
}

// SetNameChanges sets the "name_changes" field.
func (uc *UserCreate) SetNameChanges(i int8) *UserCreate {
	uc.mutation.SetNameChanges(i)
	return uc
}

// SetNillableNameChanges sets the "name_changes" field if the given value is not nil.
func (uc *UserCreate) SetNillableNameChanges(i *int8) *UserCreate {
	if i != nil {
		uc.SetNameChanges(*i)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetVerified sets the "verified" field.
func (uc *UserCreate) SetVerified(b bool) *UserCreate {
	uc.mutation.SetVerified(b)
	return uc
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (uc *UserCreate) SetNillableVerified(b *bool) *UserCreate {
	if b != nil {
		uc.SetVerified(*b)
	}
	return uc
}

// SetLocked sets the "locked" field.
func (uc *UserCreate) SetLocked(b bool) *UserCreate {
	uc.mutation.SetLocked(b)
	return uc
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (uc *UserCreate) SetNillableLocked(b *bool) *UserCreate {
	if b != nil {
		uc.SetLocked(*b)
	}
	return uc
}

// SetLastLogin sets the "last_login" field.
func (uc *UserCreate) SetLastLogin(t time.Time) *UserCreate {
	uc.mutation.SetLastLogin(t)
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// AddAmazonShareIDs adds the "amazon_shares" edge to the AmazonShare entity by IDs.
func (uc *UserCreate) AddAmazonShareIDs(ids ...int) *UserCreate {
	uc.mutation.AddAmazonShareIDs(ids...)
	return uc
}

// AddAmazonShares adds the "amazon_shares" edges to the AmazonShare entity.
func (uc *UserCreate) AddAmazonShares(a ...*AmazonShare) *UserCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddAmazonShareIDs(ids...)
}

// AddAmazonListIDs adds the "amazon_lists" edge to the AmazonList entity by IDs.
func (uc *UserCreate) AddAmazonListIDs(ids ...int) *UserCreate {
	uc.mutation.AddAmazonListIDs(ids...)
	return uc
}

// AddAmazonLists adds the "amazon_lists" edges to the AmazonList entity.
func (uc *UserCreate) AddAmazonLists(a ...*AmazonList) *UserCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddAmazonListIDs(ids...)
}

// AddDrinkIDs adds the "drinks" edge to the Drink entity by IDs.
func (uc *UserCreate) AddDrinkIDs(ids ...int) *UserCreate {
	uc.mutation.AddDrinkIDs(ids...)
	return uc
}

// AddDrinks adds the "drinks" edges to the Drink entity.
func (uc *UserCreate) AddDrinks(d ...*Drink) *UserCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddDrinkIDs(ids...)
}

// AddGroceryListIDs adds the "grocery_lists" edge to the GroceryList entity by IDs.
func (uc *UserCreate) AddGroceryListIDs(ids ...int) *UserCreate {
	uc.mutation.AddGroceryListIDs(ids...)
	return uc
}

// AddGroceryLists adds the "grocery_lists" edges to the GroceryList entity.
func (uc *UserCreate) AddGroceryLists(g ...*GroceryList) *UserCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uc.AddGroceryListIDs(ids...)
}

// AddGroceryListShareIDs adds the "grocery_list_shares" edge to the GroceryListShare entity by IDs.
func (uc *UserCreate) AddGroceryListShareIDs(ids ...int) *UserCreate {
	uc.mutation.AddGroceryListShareIDs(ids...)
	return uc
}

// AddGroceryListShares adds the "grocery_list_shares" edges to the GroceryListShare entity.
func (uc *UserCreate) AddGroceryListShares(g ...*GroceryListShare) *UserCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uc.AddGroceryListShareIDs(ids...)
}

// AddMovieListIDs adds the "movie_lists" edge to the MovieList entity by IDs.
func (uc *UserCreate) AddMovieListIDs(ids ...int) *UserCreate {
	uc.mutation.AddMovieListIDs(ids...)
	return uc
}

// AddMovieLists adds the "movie_lists" edges to the MovieList entity.
func (uc *UserCreate) AddMovieLists(m ...*MovieList) *UserCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uc.AddMovieListIDs(ids...)
}

// AddMovieListShareIDs adds the "movie_list_shares" edge to the MovieListShare entity by IDs.
func (uc *UserCreate) AddMovieListShareIDs(ids ...int) *UserCreate {
	uc.mutation.AddMovieListShareIDs(ids...)
	return uc
}

// AddMovieListShares adds the "movie_list_shares" edges to the MovieListShare entity.
func (uc *UserCreate) AddMovieListShares(m ...*MovieListShare) *UserCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uc.AddMovieListShareIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.NameChanges(); !ok {
		v := user.DefaultNameChanges
		uc.mutation.SetNameChanges(v)
	}
	if _, ok := uc.mutation.Verified(); !ok {
		v := user.DefaultVerified
		uc.mutation.SetVerified(v)
	}
	if _, ok := uc.mutation.Locked(); !ok {
		v := user.DefaultLocked
		uc.mutation.SetLocked(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`ent: missing required field "User.display_name"`)}
	}
	if _, ok := uc.mutation.NameChanges(); !ok {
		return &ValidationError{Name: "name_changes", err: errors.New(`ent: missing required field "User.name_changes"`)}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if _, ok := uc.mutation.Verified(); !ok {
		return &ValidationError{Name: "verified", err: errors.New(`ent: missing required field "User.verified"`)}
	}
	if _, ok := uc.mutation.Locked(); !ok {
		return &ValidationError{Name: "locked", err: errors.New(`ent: missing required field "User.locked"`)}
	}
	if _, ok := uc.mutation.LastLogin(); !ok {
		return &ValidationError{Name: "last_login", err: errors.New(`ent: missing required field "User.last_login"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.DisplayName(); ok {
		_spec.SetField(user.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := uc.mutation.NameChanges(); ok {
		_spec.SetField(user.FieldNameChanges, field.TypeInt8, value)
		_node.NameChanges = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Verified(); ok {
		_spec.SetField(user.FieldVerified, field.TypeBool, value)
		_node.Verified = value
	}
	if value, ok := uc.mutation.Locked(); ok {
		_spec.SetField(user.FieldLocked, field.TypeBool, value)
		_node.Locked = value
	}
	if value, ok := uc.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
		_node.LastLogin = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := uc.mutation.AmazonSharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AmazonSharesTable,
			Columns: []string{user.AmazonSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.AmazonListsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AmazonListsTable,
			Columns: user.AmazonListsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonlist.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.DrinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DrinksTable,
			Columns: []string{user.DrinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(drink.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.GroceryListsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GroceryListsTable,
			Columns: []string{user.GroceryListsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grocerylist.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.GroceryListSharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GroceryListSharesTable,
			Columns: []string{user.GroceryListSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grocerylistshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.MovieListsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MovieListsTable,
			Columns: []string{user.MovieListsColumn},
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
	if nodes := uc.mutation.MovieListSharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MovieListSharesTable,
			Columns: []string{user.MovieListSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movielistshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
