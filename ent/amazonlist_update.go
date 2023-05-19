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
	"github.com/PaluMacil/dan2/ent/amazonlist"
	"github.com/PaluMacil/dan2/ent/amazonorder"
	"github.com/PaluMacil/dan2/ent/amazonshare"
	"github.com/PaluMacil/dan2/ent/predicate"
	"github.com/PaluMacil/dan2/ent/user"
)

// AmazonListUpdate is the builder for updating AmazonList entities.
type AmazonListUpdate struct {
	config
	hooks    []Hook
	mutation *AmazonListMutation
}

// Where appends a list predicates to the AmazonListUpdate builder.
func (alu *AmazonListUpdate) Where(ps ...predicate.AmazonList) *AmazonListUpdate {
	alu.mutation.Where(ps...)
	return alu
}

// SetCreatedAt sets the "created_at" field.
func (alu *AmazonListUpdate) SetCreatedAt(t time.Time) *AmazonListUpdate {
	alu.mutation.SetCreatedAt(t)
	return alu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (alu *AmazonListUpdate) SetNillableCreatedAt(t *time.Time) *AmazonListUpdate {
	if t != nil {
		alu.SetCreatedAt(*t)
	}
	return alu
}

// AddAmazonOrderIDs adds the "amazon_orders" edge to the AmazonOrder entity by IDs.
func (alu *AmazonListUpdate) AddAmazonOrderIDs(ids ...int) *AmazonListUpdate {
	alu.mutation.AddAmazonOrderIDs(ids...)
	return alu
}

// AddAmazonOrders adds the "amazon_orders" edges to the AmazonOrder entity.
func (alu *AmazonListUpdate) AddAmazonOrders(a ...*AmazonOrder) *AmazonListUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return alu.AddAmazonOrderIDs(ids...)
}

// AddOwnerIDs adds the "owner" edge to the User entity by IDs.
func (alu *AmazonListUpdate) AddOwnerIDs(ids ...int) *AmazonListUpdate {
	alu.mutation.AddOwnerIDs(ids...)
	return alu
}

// AddOwner adds the "owner" edges to the User entity.
func (alu *AmazonListUpdate) AddOwner(u ...*User) *AmazonListUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return alu.AddOwnerIDs(ids...)
}

// AddAmazonShareIDs adds the "amazon_shares" edge to the AmazonShare entity by IDs.
func (alu *AmazonListUpdate) AddAmazonShareIDs(ids ...int) *AmazonListUpdate {
	alu.mutation.AddAmazonShareIDs(ids...)
	return alu
}

// AddAmazonShares adds the "amazon_shares" edges to the AmazonShare entity.
func (alu *AmazonListUpdate) AddAmazonShares(a ...*AmazonShare) *AmazonListUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return alu.AddAmazonShareIDs(ids...)
}

// Mutation returns the AmazonListMutation object of the builder.
func (alu *AmazonListUpdate) Mutation() *AmazonListMutation {
	return alu.mutation
}

// ClearAmazonOrders clears all "amazon_orders" edges to the AmazonOrder entity.
func (alu *AmazonListUpdate) ClearAmazonOrders() *AmazonListUpdate {
	alu.mutation.ClearAmazonOrders()
	return alu
}

// RemoveAmazonOrderIDs removes the "amazon_orders" edge to AmazonOrder entities by IDs.
func (alu *AmazonListUpdate) RemoveAmazonOrderIDs(ids ...int) *AmazonListUpdate {
	alu.mutation.RemoveAmazonOrderIDs(ids...)
	return alu
}

// RemoveAmazonOrders removes "amazon_orders" edges to AmazonOrder entities.
func (alu *AmazonListUpdate) RemoveAmazonOrders(a ...*AmazonOrder) *AmazonListUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return alu.RemoveAmazonOrderIDs(ids...)
}

// ClearOwner clears all "owner" edges to the User entity.
func (alu *AmazonListUpdate) ClearOwner() *AmazonListUpdate {
	alu.mutation.ClearOwner()
	return alu
}

// RemoveOwnerIDs removes the "owner" edge to User entities by IDs.
func (alu *AmazonListUpdate) RemoveOwnerIDs(ids ...int) *AmazonListUpdate {
	alu.mutation.RemoveOwnerIDs(ids...)
	return alu
}

// RemoveOwner removes "owner" edges to User entities.
func (alu *AmazonListUpdate) RemoveOwner(u ...*User) *AmazonListUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return alu.RemoveOwnerIDs(ids...)
}

// ClearAmazonShares clears all "amazon_shares" edges to the AmazonShare entity.
func (alu *AmazonListUpdate) ClearAmazonShares() *AmazonListUpdate {
	alu.mutation.ClearAmazonShares()
	return alu
}

// RemoveAmazonShareIDs removes the "amazon_shares" edge to AmazonShare entities by IDs.
func (alu *AmazonListUpdate) RemoveAmazonShareIDs(ids ...int) *AmazonListUpdate {
	alu.mutation.RemoveAmazonShareIDs(ids...)
	return alu
}

// RemoveAmazonShares removes "amazon_shares" edges to AmazonShare entities.
func (alu *AmazonListUpdate) RemoveAmazonShares(a ...*AmazonShare) *AmazonListUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return alu.RemoveAmazonShareIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (alu *AmazonListUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, alu.sqlSave, alu.mutation, alu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (alu *AmazonListUpdate) SaveX(ctx context.Context) int {
	affected, err := alu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (alu *AmazonListUpdate) Exec(ctx context.Context) error {
	_, err := alu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alu *AmazonListUpdate) ExecX(ctx context.Context) {
	if err := alu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (alu *AmazonListUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(amazonlist.Table, amazonlist.Columns, sqlgraph.NewFieldSpec(amazonlist.FieldID, field.TypeInt))
	if ps := alu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := alu.mutation.CreatedAt(); ok {
		_spec.SetField(amazonlist.FieldCreatedAt, field.TypeTime, value)
	}
	if alu.mutation.AmazonOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   amazonlist.AmazonOrdersTable,
			Columns: []string{amazonlist.AmazonOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonorder.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := alu.mutation.RemovedAmazonOrdersIDs(); len(nodes) > 0 && !alu.mutation.AmazonOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   amazonlist.AmazonOrdersTable,
			Columns: []string{amazonlist.AmazonOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonorder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := alu.mutation.AmazonOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   amazonlist.AmazonOrdersTable,
			Columns: []string{amazonlist.AmazonOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonorder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if alu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   amazonlist.OwnerTable,
			Columns: amazonlist.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := alu.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !alu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   amazonlist.OwnerTable,
			Columns: amazonlist.OwnerPrimaryKey,
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
	if nodes := alu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   amazonlist.OwnerTable,
			Columns: amazonlist.OwnerPrimaryKey,
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
	if alu.mutation.AmazonSharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   amazonlist.AmazonSharesTable,
			Columns: []string{amazonlist.AmazonSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonshare.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := alu.mutation.RemovedAmazonSharesIDs(); len(nodes) > 0 && !alu.mutation.AmazonSharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   amazonlist.AmazonSharesTable,
			Columns: []string{amazonlist.AmazonSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := alu.mutation.AmazonSharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   amazonlist.AmazonSharesTable,
			Columns: []string{amazonlist.AmazonSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, alu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{amazonlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	alu.mutation.done = true
	return n, nil
}

// AmazonListUpdateOne is the builder for updating a single AmazonList entity.
type AmazonListUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AmazonListMutation
}

// SetCreatedAt sets the "created_at" field.
func (aluo *AmazonListUpdateOne) SetCreatedAt(t time.Time) *AmazonListUpdateOne {
	aluo.mutation.SetCreatedAt(t)
	return aluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aluo *AmazonListUpdateOne) SetNillableCreatedAt(t *time.Time) *AmazonListUpdateOne {
	if t != nil {
		aluo.SetCreatedAt(*t)
	}
	return aluo
}

// AddAmazonOrderIDs adds the "amazon_orders" edge to the AmazonOrder entity by IDs.
func (aluo *AmazonListUpdateOne) AddAmazonOrderIDs(ids ...int) *AmazonListUpdateOne {
	aluo.mutation.AddAmazonOrderIDs(ids...)
	return aluo
}

// AddAmazonOrders adds the "amazon_orders" edges to the AmazonOrder entity.
func (aluo *AmazonListUpdateOne) AddAmazonOrders(a ...*AmazonOrder) *AmazonListUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aluo.AddAmazonOrderIDs(ids...)
}

// AddOwnerIDs adds the "owner" edge to the User entity by IDs.
func (aluo *AmazonListUpdateOne) AddOwnerIDs(ids ...int) *AmazonListUpdateOne {
	aluo.mutation.AddOwnerIDs(ids...)
	return aluo
}

// AddOwner adds the "owner" edges to the User entity.
func (aluo *AmazonListUpdateOne) AddOwner(u ...*User) *AmazonListUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return aluo.AddOwnerIDs(ids...)
}

// AddAmazonShareIDs adds the "amazon_shares" edge to the AmazonShare entity by IDs.
func (aluo *AmazonListUpdateOne) AddAmazonShareIDs(ids ...int) *AmazonListUpdateOne {
	aluo.mutation.AddAmazonShareIDs(ids...)
	return aluo
}

// AddAmazonShares adds the "amazon_shares" edges to the AmazonShare entity.
func (aluo *AmazonListUpdateOne) AddAmazonShares(a ...*AmazonShare) *AmazonListUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aluo.AddAmazonShareIDs(ids...)
}

// Mutation returns the AmazonListMutation object of the builder.
func (aluo *AmazonListUpdateOne) Mutation() *AmazonListMutation {
	return aluo.mutation
}

// ClearAmazonOrders clears all "amazon_orders" edges to the AmazonOrder entity.
func (aluo *AmazonListUpdateOne) ClearAmazonOrders() *AmazonListUpdateOne {
	aluo.mutation.ClearAmazonOrders()
	return aluo
}

// RemoveAmazonOrderIDs removes the "amazon_orders" edge to AmazonOrder entities by IDs.
func (aluo *AmazonListUpdateOne) RemoveAmazonOrderIDs(ids ...int) *AmazonListUpdateOne {
	aluo.mutation.RemoveAmazonOrderIDs(ids...)
	return aluo
}

// RemoveAmazonOrders removes "amazon_orders" edges to AmazonOrder entities.
func (aluo *AmazonListUpdateOne) RemoveAmazonOrders(a ...*AmazonOrder) *AmazonListUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aluo.RemoveAmazonOrderIDs(ids...)
}

// ClearOwner clears all "owner" edges to the User entity.
func (aluo *AmazonListUpdateOne) ClearOwner() *AmazonListUpdateOne {
	aluo.mutation.ClearOwner()
	return aluo
}

// RemoveOwnerIDs removes the "owner" edge to User entities by IDs.
func (aluo *AmazonListUpdateOne) RemoveOwnerIDs(ids ...int) *AmazonListUpdateOne {
	aluo.mutation.RemoveOwnerIDs(ids...)
	return aluo
}

// RemoveOwner removes "owner" edges to User entities.
func (aluo *AmazonListUpdateOne) RemoveOwner(u ...*User) *AmazonListUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return aluo.RemoveOwnerIDs(ids...)
}

// ClearAmazonShares clears all "amazon_shares" edges to the AmazonShare entity.
func (aluo *AmazonListUpdateOne) ClearAmazonShares() *AmazonListUpdateOne {
	aluo.mutation.ClearAmazonShares()
	return aluo
}

// RemoveAmazonShareIDs removes the "amazon_shares" edge to AmazonShare entities by IDs.
func (aluo *AmazonListUpdateOne) RemoveAmazonShareIDs(ids ...int) *AmazonListUpdateOne {
	aluo.mutation.RemoveAmazonShareIDs(ids...)
	return aluo
}

// RemoveAmazonShares removes "amazon_shares" edges to AmazonShare entities.
func (aluo *AmazonListUpdateOne) RemoveAmazonShares(a ...*AmazonShare) *AmazonListUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aluo.RemoveAmazonShareIDs(ids...)
}

// Where appends a list predicates to the AmazonListUpdate builder.
func (aluo *AmazonListUpdateOne) Where(ps ...predicate.AmazonList) *AmazonListUpdateOne {
	aluo.mutation.Where(ps...)
	return aluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aluo *AmazonListUpdateOne) Select(field string, fields ...string) *AmazonListUpdateOne {
	aluo.fields = append([]string{field}, fields...)
	return aluo
}

// Save executes the query and returns the updated AmazonList entity.
func (aluo *AmazonListUpdateOne) Save(ctx context.Context) (*AmazonList, error) {
	return withHooks(ctx, aluo.sqlSave, aluo.mutation, aluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aluo *AmazonListUpdateOne) SaveX(ctx context.Context) *AmazonList {
	node, err := aluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aluo *AmazonListUpdateOne) Exec(ctx context.Context) error {
	_, err := aluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aluo *AmazonListUpdateOne) ExecX(ctx context.Context) {
	if err := aluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aluo *AmazonListUpdateOne) sqlSave(ctx context.Context) (_node *AmazonList, err error) {
	_spec := sqlgraph.NewUpdateSpec(amazonlist.Table, amazonlist.Columns, sqlgraph.NewFieldSpec(amazonlist.FieldID, field.TypeInt))
	id, ok := aluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AmazonList.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, amazonlist.FieldID)
		for _, f := range fields {
			if !amazonlist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != amazonlist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aluo.mutation.CreatedAt(); ok {
		_spec.SetField(amazonlist.FieldCreatedAt, field.TypeTime, value)
	}
	if aluo.mutation.AmazonOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   amazonlist.AmazonOrdersTable,
			Columns: []string{amazonlist.AmazonOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonorder.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aluo.mutation.RemovedAmazonOrdersIDs(); len(nodes) > 0 && !aluo.mutation.AmazonOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   amazonlist.AmazonOrdersTable,
			Columns: []string{amazonlist.AmazonOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonorder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aluo.mutation.AmazonOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   amazonlist.AmazonOrdersTable,
			Columns: []string{amazonlist.AmazonOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonorder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aluo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   amazonlist.OwnerTable,
			Columns: amazonlist.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aluo.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !aluo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   amazonlist.OwnerTable,
			Columns: amazonlist.OwnerPrimaryKey,
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
	if nodes := aluo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   amazonlist.OwnerTable,
			Columns: amazonlist.OwnerPrimaryKey,
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
	if aluo.mutation.AmazonSharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   amazonlist.AmazonSharesTable,
			Columns: []string{amazonlist.AmazonSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonshare.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aluo.mutation.RemovedAmazonSharesIDs(); len(nodes) > 0 && !aluo.mutation.AmazonSharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   amazonlist.AmazonSharesTable,
			Columns: []string{amazonlist.AmazonSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aluo.mutation.AmazonSharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   amazonlist.AmazonSharesTable,
			Columns: []string{amazonlist.AmazonSharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amazonshare.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AmazonList{config: aluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{amazonlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aluo.mutation.done = true
	return _node, nil
}
