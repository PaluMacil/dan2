// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PaluMacil/dan2/ent/drink"
	"github.com/PaluMacil/dan2/ent/user"
)

// DrinkCreate is the builder for creating a Drink entity.
type DrinkCreate struct {
	config
	mutation *DrinkMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (dc *DrinkCreate) SetType(d drink.Type) *DrinkCreate {
	dc.mutation.SetType(d)
	return dc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (dc *DrinkCreate) SetNillableType(d *drink.Type) *DrinkCreate {
	if d != nil {
		dc.SetType(*d)
	}
	return dc
}

// SetAbv sets the "abv" field.
func (dc *DrinkCreate) SetAbv(i int8) *DrinkCreate {
	dc.mutation.SetAbv(i)
	return dc
}

// SetNillableAbv sets the "abv" field if the given value is not nil.
func (dc *DrinkCreate) SetNillableAbv(i *int8) *DrinkCreate {
	if i != nil {
		dc.SetAbv(*i)
	}
	return dc
}

// SetOunces sets the "ounces" field.
func (dc *DrinkCreate) SetOunces(i int8) *DrinkCreate {
	dc.mutation.SetOunces(i)
	return dc
}

// SetNillableOunces sets the "ounces" field if the given value is not nil.
func (dc *DrinkCreate) SetNillableOunces(i *int8) *DrinkCreate {
	if i != nil {
		dc.SetOunces(*i)
	}
	return dc
}

// SetYear sets the "year" field.
func (dc *DrinkCreate) SetYear(i int) *DrinkCreate {
	dc.mutation.SetYear(i)
	return dc
}

// SetNillableYear sets the "year" field if the given value is not nil.
func (dc *DrinkCreate) SetNillableYear(i *int) *DrinkCreate {
	if i != nil {
		dc.SetYear(*i)
	}
	return dc
}

// SetMonth sets the "month" field.
func (dc *DrinkCreate) SetMonth(i int) *DrinkCreate {
	dc.mutation.SetMonth(i)
	return dc
}

// SetNillableMonth sets the "month" field if the given value is not nil.
func (dc *DrinkCreate) SetNillableMonth(i *int) *DrinkCreate {
	if i != nil {
		dc.SetMonth(*i)
	}
	return dc
}

// SetDay sets the "day" field.
func (dc *DrinkCreate) SetDay(i int) *DrinkCreate {
	dc.mutation.SetDay(i)
	return dc
}

// SetNillableDay sets the "day" field if the given value is not nil.
func (dc *DrinkCreate) SetNillableDay(i *int) *DrinkCreate {
	if i != nil {
		dc.SetDay(*i)
	}
	return dc
}

// SetNote sets the "note" field.
func (dc *DrinkCreate) SetNote(s string) *DrinkCreate {
	dc.mutation.SetNote(s)
	return dc
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (dc *DrinkCreate) SetNillableNote(s *string) *DrinkCreate {
	if s != nil {
		dc.SetNote(*s)
	}
	return dc
}

// SetCreatedAt sets the "created_at" field.
func (dc *DrinkCreate) SetCreatedAt(t time.Time) *DrinkCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DrinkCreate) SetNillableCreatedAt(t *time.Time) *DrinkCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (dc *DrinkCreate) SetOwnerID(id int) *DrinkCreate {
	dc.mutation.SetOwnerID(id)
	return dc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (dc *DrinkCreate) SetNillableOwnerID(id *int) *DrinkCreate {
	if id != nil {
		dc = dc.SetOwnerID(*id)
	}
	return dc
}

// SetOwner sets the "owner" edge to the User entity.
func (dc *DrinkCreate) SetOwner(u *User) *DrinkCreate {
	return dc.SetOwnerID(u.ID)
}

// Mutation returns the DrinkMutation object of the builder.
func (dc *DrinkCreate) Mutation() *DrinkMutation {
	return dc.mutation
}

// Save creates the Drink in the database.
func (dc *DrinkCreate) Save(ctx context.Context) (*Drink, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DrinkCreate) SaveX(ctx context.Context) *Drink {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DrinkCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DrinkCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DrinkCreate) defaults() {
	if _, ok := dc.mutation.GetType(); !ok {
		v := drink.DefaultType
		dc.mutation.SetType(v)
	}
	if _, ok := dc.mutation.Abv(); !ok {
		v := drink.DefaultAbv
		dc.mutation.SetAbv(v)
	}
	if _, ok := dc.mutation.Ounces(); !ok {
		v := drink.DefaultOunces
		dc.mutation.SetOunces(v)
	}
	if _, ok := dc.mutation.Year(); !ok {
		v := drink.DefaultYear()
		dc.mutation.SetYear(v)
	}
	if _, ok := dc.mutation.Month(); !ok {
		v := drink.DefaultMonth()
		dc.mutation.SetMonth(v)
	}
	if _, ok := dc.mutation.Day(); !ok {
		v := drink.DefaultDay()
		dc.mutation.SetDay(v)
	}
	if _, ok := dc.mutation.Note(); !ok {
		v := drink.DefaultNote
		dc.mutation.SetNote(v)
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := drink.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DrinkCreate) check() error {
	if _, ok := dc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Drink.type"`)}
	}
	if v, ok := dc.mutation.GetType(); ok {
		if err := drink.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Drink.type": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Abv(); !ok {
		return &ValidationError{Name: "abv", err: errors.New(`ent: missing required field "Drink.abv"`)}
	}
	if _, ok := dc.mutation.Ounces(); !ok {
		return &ValidationError{Name: "ounces", err: errors.New(`ent: missing required field "Drink.ounces"`)}
	}
	if _, ok := dc.mutation.Year(); !ok {
		return &ValidationError{Name: "year", err: errors.New(`ent: missing required field "Drink.year"`)}
	}
	if _, ok := dc.mutation.Month(); !ok {
		return &ValidationError{Name: "month", err: errors.New(`ent: missing required field "Drink.month"`)}
	}
	if _, ok := dc.mutation.Day(); !ok {
		return &ValidationError{Name: "day", err: errors.New(`ent: missing required field "Drink.day"`)}
	}
	if _, ok := dc.mutation.Note(); !ok {
		return &ValidationError{Name: "note", err: errors.New(`ent: missing required field "Drink.note"`)}
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Drink.created_at"`)}
	}
	return nil
}

func (dc *DrinkCreate) sqlSave(ctx context.Context) (*Drink, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DrinkCreate) createSpec() (*Drink, *sqlgraph.CreateSpec) {
	var (
		_node = &Drink{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(drink.Table, sqlgraph.NewFieldSpec(drink.FieldID, field.TypeInt))
	)
	if value, ok := dc.mutation.GetType(); ok {
		_spec.SetField(drink.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := dc.mutation.Abv(); ok {
		_spec.SetField(drink.FieldAbv, field.TypeInt8, value)
		_node.Abv = value
	}
	if value, ok := dc.mutation.Ounces(); ok {
		_spec.SetField(drink.FieldOunces, field.TypeInt8, value)
		_node.Ounces = value
	}
	if value, ok := dc.mutation.Year(); ok {
		_spec.SetField(drink.FieldYear, field.TypeInt, value)
		_node.Year = value
	}
	if value, ok := dc.mutation.Month(); ok {
		_spec.SetField(drink.FieldMonth, field.TypeInt, value)
		_node.Month = value
	}
	if value, ok := dc.mutation.Day(); ok {
		_spec.SetField(drink.FieldDay, field.TypeInt, value)
		_node.Day = value
	}
	if value, ok := dc.mutation.Note(); ok {
		_spec.SetField(drink.FieldNote, field.TypeString, value)
		_node.Note = value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(drink.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := dc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drink.OwnerTable,
			Columns: []string{drink.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_drinks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DrinkCreateBulk is the builder for creating many Drink entities in bulk.
type DrinkCreateBulk struct {
	config
	builders []*DrinkCreate
}

// Save creates the Drink entities in the database.
func (dcb *DrinkCreateBulk) Save(ctx context.Context) ([]*Drink, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Drink, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DrinkMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DrinkCreateBulk) SaveX(ctx context.Context) []*Drink {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DrinkCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DrinkCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
