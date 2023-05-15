// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PaluMacil/dan2/ent/grocerylist"
	"github.com/PaluMacil/dan2/ent/grocerylistitem"
	"github.com/PaluMacil/dan2/ent/grocerylistshare"
	"github.com/PaluMacil/dan2/ent/user"
)

// GroceryListCreate is the builder for creating a GroceryList entity.
type GroceryListCreate struct {
	config
	mutation *GroceryListMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (glc *GroceryListCreate) SetName(s string) *GroceryListCreate {
	glc.mutation.SetName(s)
	return glc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (glc *GroceryListCreate) SetNillableName(s *string) *GroceryListCreate {
	if s != nil {
		glc.SetName(*s)
	}
	return glc
}

// SetNote sets the "note" field.
func (glc *GroceryListCreate) SetNote(s string) *GroceryListCreate {
	glc.mutation.SetNote(s)
	return glc
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (glc *GroceryListCreate) SetNillableNote(s *string) *GroceryListCreate {
	if s != nil {
		glc.SetNote(*s)
	}
	return glc
}

// SetArchived sets the "archived" field.
func (glc *GroceryListCreate) SetArchived(b bool) *GroceryListCreate {
	glc.mutation.SetArchived(b)
	return glc
}

// SetNillableArchived sets the "archived" field if the given value is not nil.
func (glc *GroceryListCreate) SetNillableArchived(b *bool) *GroceryListCreate {
	if b != nil {
		glc.SetArchived(*b)
	}
	return glc
}

// SetCreatedAt sets the "created_at" field.
func (glc *GroceryListCreate) SetCreatedAt(t time.Time) *GroceryListCreate {
	glc.mutation.SetCreatedAt(t)
	return glc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (glc *GroceryListCreate) SetNillableCreatedAt(t *time.Time) *GroceryListCreate {
	if t != nil {
		glc.SetCreatedAt(*t)
	}
	return glc
}

// AddGroceryListItemIDs adds the "grocery_list_items" edge to the GroceryListItem entity by IDs.
func (glc *GroceryListCreate) AddGroceryListItemIDs(ids ...int) *GroceryListCreate {
	glc.mutation.AddGroceryListItemIDs(ids...)
	return glc
}

// AddGroceryListItems adds the "grocery_list_items" edges to the GroceryListItem entity.
func (glc *GroceryListCreate) AddGroceryListItems(g ...*GroceryListItem) *GroceryListCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return glc.AddGroceryListItemIDs(ids...)
}

// AddOwnerIDs adds the "owner" edge to the User entity by IDs.
func (glc *GroceryListCreate) AddOwnerIDs(ids ...int) *GroceryListCreate {
	glc.mutation.AddOwnerIDs(ids...)
	return glc
}

// AddOwner adds the "owner" edges to the User entity.
func (glc *GroceryListCreate) AddOwner(u ...*User) *GroceryListCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return glc.AddOwnerIDs(ids...)
}

// AddGroceryListShareIDs adds the "grocery_list_shares" edge to the GroceryListShare entity by IDs.
func (glc *GroceryListCreate) AddGroceryListShareIDs(ids ...int) *GroceryListCreate {
	glc.mutation.AddGroceryListShareIDs(ids...)
	return glc
}

// AddGroceryListShares adds the "grocery_list_shares" edges to the GroceryListShare entity.
func (glc *GroceryListCreate) AddGroceryListShares(g ...*GroceryListShare) *GroceryListCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return glc.AddGroceryListShareIDs(ids...)
}

// Mutation returns the GroceryListMutation object of the builder.
func (glc *GroceryListCreate) Mutation() *GroceryListMutation {
	return glc.mutation
}

// Save creates the GroceryList in the database.
func (glc *GroceryListCreate) Save(ctx context.Context) (*GroceryList, error) {
	glc.defaults()
	return withHooks(ctx, glc.sqlSave, glc.mutation, glc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (glc *GroceryListCreate) SaveX(ctx context.Context) *GroceryList {
	v, err := glc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (glc *GroceryListCreate) Exec(ctx context.Context) error {
	_, err := glc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (glc *GroceryListCreate) ExecX(ctx context.Context) {
	if err := glc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (glc *GroceryListCreate) defaults() {
	if _, ok := glc.mutation.Name(); !ok {
		v := grocerylist.DefaultName
		glc.mutation.SetName(v)
	}
	if _, ok := glc.mutation.Note(); !ok {
		v := grocerylist.DefaultNote
		glc.mutation.SetNote(v)
	}
	if _, ok := glc.mutation.Archived(); !ok {
		v := grocerylist.DefaultArchived
		glc.mutation.SetArchived(v)
	}
	if _, ok := glc.mutation.CreatedAt(); !ok {
		v := grocerylist.DefaultCreatedAt()
		glc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (glc *GroceryListCreate) check() error {
	if _, ok := glc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "GroceryList.name"`)}
	}
	if _, ok := glc.mutation.Note(); !ok {
		return &ValidationError{Name: "note", err: errors.New(`ent: missing required field "GroceryList.note"`)}
	}
	if _, ok := glc.mutation.Archived(); !ok {
		return &ValidationError{Name: "archived", err: errors.New(`ent: missing required field "GroceryList.archived"`)}
	}
	if _, ok := glc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GroceryList.created_at"`)}
	}
	return nil
}

func (glc *GroceryListCreate) sqlSave(ctx context.Context) (*GroceryList, error) {
	if err := glc.check(); err != nil {
		return nil, err
	}
	_node, _spec := glc.createSpec()
	if err := sqlgraph.CreateNode(ctx, glc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	glc.mutation.id = &_node.ID
	glc.mutation.done = true
	return _node, nil
}

func (glc *GroceryListCreate) createSpec() (*GroceryList, *sqlgraph.CreateSpec) {
	var (
		_node = &GroceryList{config: glc.config}
		_spec = sqlgraph.NewCreateSpec(grocerylist.Table, sqlgraph.NewFieldSpec(grocerylist.FieldID, field.TypeInt))
	)
	if value, ok := glc.mutation.Name(); ok {
		_spec.SetField(grocerylist.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := glc.mutation.Note(); ok {
		_spec.SetField(grocerylist.FieldNote, field.TypeString, value)
		_node.Note = value
	}
	if value, ok := glc.mutation.Archived(); ok {
		_spec.SetField(grocerylist.FieldArchived, field.TypeBool, value)
		_node.Archived = value
	}
	if value, ok := glc.mutation.CreatedAt(); ok {
		_spec.SetField(grocerylist.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := glc.mutation.GroceryListItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   grocerylist.GroceryListItemsTable,
			Columns: grocerylist.GroceryListItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grocerylistitem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := glc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   grocerylist.OwnerTable,
			Columns: grocerylist.OwnerPrimaryKey,
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
	if nodes := glc.mutation.GroceryListSharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   grocerylist.GroceryListSharesTable,
			Columns: grocerylist.GroceryListSharesPrimaryKey,
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
	return _node, _spec
}

// GroceryListCreateBulk is the builder for creating many GroceryList entities in bulk.
type GroceryListCreateBulk struct {
	config
	builders []*GroceryListCreate
}

// Save creates the GroceryList entities in the database.
func (glcb *GroceryListCreateBulk) Save(ctx context.Context) ([]*GroceryList, error) {
	specs := make([]*sqlgraph.CreateSpec, len(glcb.builders))
	nodes := make([]*GroceryList, len(glcb.builders))
	mutators := make([]Mutator, len(glcb.builders))
	for i := range glcb.builders {
		func(i int, root context.Context) {
			builder := glcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroceryListMutation)
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
					_, err = mutators[i+1].Mutate(root, glcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, glcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, glcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (glcb *GroceryListCreateBulk) SaveX(ctx context.Context) []*GroceryList {
	v, err := glcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (glcb *GroceryListCreateBulk) Exec(ctx context.Context) error {
	_, err := glcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (glcb *GroceryListCreateBulk) ExecX(ctx context.Context) {
	if err := glcb.Exec(ctx); err != nil {
		panic(err)
	}
}
