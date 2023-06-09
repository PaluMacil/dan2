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
	"github.com/PaluMacil/dan2/ent/amazonorder"
)

// AmazonOrderCreate is the builder for creating a AmazonOrder entity.
type AmazonOrderCreate struct {
	config
	mutation *AmazonOrderMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (aoc *AmazonOrderCreate) SetName(s string) *AmazonOrderCreate {
	aoc.mutation.SetName(s)
	return aoc
}

// SetCategory sets the "category" field.
func (aoc *AmazonOrderCreate) SetCategory(s string) *AmazonOrderCreate {
	aoc.mutation.SetCategory(s)
	return aoc
}

// SetBrand sets the "brand" field.
func (aoc *AmazonOrderCreate) SetBrand(s string) *AmazonOrderCreate {
	aoc.mutation.SetBrand(s)
	return aoc
}

// SetSeller sets the "seller" field.
func (aoc *AmazonOrderCreate) SetSeller(s string) *AmazonOrderCreate {
	aoc.mutation.SetSeller(s)
	return aoc
}

// SetAddress1 sets the "address1" field.
func (aoc *AmazonOrderCreate) SetAddress1(s string) *AmazonOrderCreate {
	aoc.mutation.SetAddress1(s)
	return aoc
}

// SetAddress2 sets the "address2" field.
func (aoc *AmazonOrderCreate) SetAddress2(s string) *AmazonOrderCreate {
	aoc.mutation.SetAddress2(s)
	return aoc
}

// SetCity sets the "city" field.
func (aoc *AmazonOrderCreate) SetCity(s string) *AmazonOrderCreate {
	aoc.mutation.SetCity(s)
	return aoc
}

// SetState sets the "state" field.
func (aoc *AmazonOrderCreate) SetState(s string) *AmazonOrderCreate {
	aoc.mutation.SetState(s)
	return aoc
}

// SetZip sets the "zip" field.
func (aoc *AmazonOrderCreate) SetZip(s string) *AmazonOrderCreate {
	aoc.mutation.SetZip(s)
	return aoc
}

// SetPrice sets the "price" field.
func (aoc *AmazonOrderCreate) SetPrice(f float32) *AmazonOrderCreate {
	aoc.mutation.SetPrice(f)
	return aoc
}

// SetTax sets the "tax" field.
func (aoc *AmazonOrderCreate) SetTax(f float32) *AmazonOrderCreate {
	aoc.mutation.SetTax(f)
	return aoc
}

// SetRefund sets the "refund" field.
func (aoc *AmazonOrderCreate) SetRefund(b bool) *AmazonOrderCreate {
	aoc.mutation.SetRefund(b)
	return aoc
}

// SetOrderedAt sets the "ordered_at" field.
func (aoc *AmazonOrderCreate) SetOrderedAt(t time.Time) *AmazonOrderCreate {
	aoc.mutation.SetOrderedAt(t)
	return aoc
}

// SetCreatedAt sets the "created_at" field.
func (aoc *AmazonOrderCreate) SetCreatedAt(t time.Time) *AmazonOrderCreate {
	aoc.mutation.SetCreatedAt(t)
	return aoc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aoc *AmazonOrderCreate) SetNillableCreatedAt(t *time.Time) *AmazonOrderCreate {
	if t != nil {
		aoc.SetCreatedAt(*t)
	}
	return aoc
}

// SetAmazonListID sets the "amazon_list" edge to the AmazonList entity by ID.
func (aoc *AmazonOrderCreate) SetAmazonListID(id int) *AmazonOrderCreate {
	aoc.mutation.SetAmazonListID(id)
	return aoc
}

// SetNillableAmazonListID sets the "amazon_list" edge to the AmazonList entity by ID if the given value is not nil.
func (aoc *AmazonOrderCreate) SetNillableAmazonListID(id *int) *AmazonOrderCreate {
	if id != nil {
		aoc = aoc.SetAmazonListID(*id)
	}
	return aoc
}

// SetAmazonList sets the "amazon_list" edge to the AmazonList entity.
func (aoc *AmazonOrderCreate) SetAmazonList(a *AmazonList) *AmazonOrderCreate {
	return aoc.SetAmazonListID(a.ID)
}

// Mutation returns the AmazonOrderMutation object of the builder.
func (aoc *AmazonOrderCreate) Mutation() *AmazonOrderMutation {
	return aoc.mutation
}

// Save creates the AmazonOrder in the database.
func (aoc *AmazonOrderCreate) Save(ctx context.Context) (*AmazonOrder, error) {
	aoc.defaults()
	return withHooks(ctx, aoc.sqlSave, aoc.mutation, aoc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (aoc *AmazonOrderCreate) SaveX(ctx context.Context) *AmazonOrder {
	v, err := aoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aoc *AmazonOrderCreate) Exec(ctx context.Context) error {
	_, err := aoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aoc *AmazonOrderCreate) ExecX(ctx context.Context) {
	if err := aoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aoc *AmazonOrderCreate) defaults() {
	if _, ok := aoc.mutation.CreatedAt(); !ok {
		v := amazonorder.DefaultCreatedAt()
		aoc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aoc *AmazonOrderCreate) check() error {
	if _, ok := aoc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AmazonOrder.name"`)}
	}
	if _, ok := aoc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "AmazonOrder.category"`)}
	}
	if _, ok := aoc.mutation.Brand(); !ok {
		return &ValidationError{Name: "brand", err: errors.New(`ent: missing required field "AmazonOrder.brand"`)}
	}
	if _, ok := aoc.mutation.Seller(); !ok {
		return &ValidationError{Name: "seller", err: errors.New(`ent: missing required field "AmazonOrder.seller"`)}
	}
	if _, ok := aoc.mutation.Address1(); !ok {
		return &ValidationError{Name: "address1", err: errors.New(`ent: missing required field "AmazonOrder.address1"`)}
	}
	if _, ok := aoc.mutation.Address2(); !ok {
		return &ValidationError{Name: "address2", err: errors.New(`ent: missing required field "AmazonOrder.address2"`)}
	}
	if _, ok := aoc.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New(`ent: missing required field "AmazonOrder.city"`)}
	}
	if _, ok := aoc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "AmazonOrder.state"`)}
	}
	if _, ok := aoc.mutation.Zip(); !ok {
		return &ValidationError{Name: "zip", err: errors.New(`ent: missing required field "AmazonOrder.zip"`)}
	}
	if _, ok := aoc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "AmazonOrder.price"`)}
	}
	if _, ok := aoc.mutation.Tax(); !ok {
		return &ValidationError{Name: "tax", err: errors.New(`ent: missing required field "AmazonOrder.tax"`)}
	}
	if _, ok := aoc.mutation.Refund(); !ok {
		return &ValidationError{Name: "refund", err: errors.New(`ent: missing required field "AmazonOrder.refund"`)}
	}
	if _, ok := aoc.mutation.OrderedAt(); !ok {
		return &ValidationError{Name: "ordered_at", err: errors.New(`ent: missing required field "AmazonOrder.ordered_at"`)}
	}
	if _, ok := aoc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AmazonOrder.created_at"`)}
	}
	return nil
}

func (aoc *AmazonOrderCreate) sqlSave(ctx context.Context) (*AmazonOrder, error) {
	if err := aoc.check(); err != nil {
		return nil, err
	}
	_node, _spec := aoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, aoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	aoc.mutation.id = &_node.ID
	aoc.mutation.done = true
	return _node, nil
}

func (aoc *AmazonOrderCreate) createSpec() (*AmazonOrder, *sqlgraph.CreateSpec) {
	var (
		_node = &AmazonOrder{config: aoc.config}
		_spec = sqlgraph.NewCreateSpec(amazonorder.Table, sqlgraph.NewFieldSpec(amazonorder.FieldID, field.TypeInt))
	)
	if value, ok := aoc.mutation.Name(); ok {
		_spec.SetField(amazonorder.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := aoc.mutation.Category(); ok {
		_spec.SetField(amazonorder.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if value, ok := aoc.mutation.Brand(); ok {
		_spec.SetField(amazonorder.FieldBrand, field.TypeString, value)
		_node.Brand = value
	}
	if value, ok := aoc.mutation.Seller(); ok {
		_spec.SetField(amazonorder.FieldSeller, field.TypeString, value)
		_node.Seller = value
	}
	if value, ok := aoc.mutation.Address1(); ok {
		_spec.SetField(amazonorder.FieldAddress1, field.TypeString, value)
		_node.Address1 = value
	}
	if value, ok := aoc.mutation.Address2(); ok {
		_spec.SetField(amazonorder.FieldAddress2, field.TypeString, value)
		_node.Address2 = value
	}
	if value, ok := aoc.mutation.City(); ok {
		_spec.SetField(amazonorder.FieldCity, field.TypeString, value)
		_node.City = value
	}
	if value, ok := aoc.mutation.State(); ok {
		_spec.SetField(amazonorder.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := aoc.mutation.Zip(); ok {
		_spec.SetField(amazonorder.FieldZip, field.TypeString, value)
		_node.Zip = value
	}
	if value, ok := aoc.mutation.Price(); ok {
		_spec.SetField(amazonorder.FieldPrice, field.TypeFloat32, value)
		_node.Price = value
	}
	if value, ok := aoc.mutation.Tax(); ok {
		_spec.SetField(amazonorder.FieldTax, field.TypeFloat32, value)
		_node.Tax = value
	}
	if value, ok := aoc.mutation.Refund(); ok {
		_spec.SetField(amazonorder.FieldRefund, field.TypeBool, value)
		_node.Refund = value
	}
	if value, ok := aoc.mutation.OrderedAt(); ok {
		_spec.SetField(amazonorder.FieldOrderedAt, field.TypeTime, value)
		_node.OrderedAt = value
	}
	if value, ok := aoc.mutation.CreatedAt(); ok {
		_spec.SetField(amazonorder.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := aoc.mutation.AmazonListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   amazonorder.AmazonListTable,
			Columns: []string{amazonorder.AmazonListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonlist.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.amazon_list_amazon_orders = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AmazonOrderCreateBulk is the builder for creating many AmazonOrder entities in bulk.
type AmazonOrderCreateBulk struct {
	config
	builders []*AmazonOrderCreate
}

// Save creates the AmazonOrder entities in the database.
func (aocb *AmazonOrderCreateBulk) Save(ctx context.Context) ([]*AmazonOrder, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aocb.builders))
	nodes := make([]*AmazonOrder, len(aocb.builders))
	mutators := make([]Mutator, len(aocb.builders))
	for i := range aocb.builders {
		func(i int, root context.Context) {
			builder := aocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AmazonOrderMutation)
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
					_, err = mutators[i+1].Mutate(root, aocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, aocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aocb *AmazonOrderCreateBulk) SaveX(ctx context.Context) []*AmazonOrder {
	v, err := aocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aocb *AmazonOrderCreateBulk) Exec(ctx context.Context) error {
	_, err := aocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aocb *AmazonOrderCreateBulk) ExecX(ctx context.Context) {
	if err := aocb.Exec(ctx); err != nil {
		panic(err)
	}
}
