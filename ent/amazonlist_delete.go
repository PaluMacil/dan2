// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PaluMacil/dan2/ent/amazonlist"
	"github.com/PaluMacil/dan2/ent/predicate"
)

// AmazonListDelete is the builder for deleting a AmazonList entity.
type AmazonListDelete struct {
	config
	hooks    []Hook
	mutation *AmazonListMutation
}

// Where appends a list predicates to the AmazonListDelete builder.
func (ald *AmazonListDelete) Where(ps ...predicate.AmazonList) *AmazonListDelete {
	ald.mutation.Where(ps...)
	return ald
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ald *AmazonListDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ald.sqlExec, ald.mutation, ald.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ald *AmazonListDelete) ExecX(ctx context.Context) int {
	n, err := ald.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ald *AmazonListDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(amazonlist.Table, sqlgraph.NewFieldSpec(amazonlist.FieldID, field.TypeInt))
	if ps := ald.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ald.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ald.mutation.done = true
	return affected, err
}

// AmazonListDeleteOne is the builder for deleting a single AmazonList entity.
type AmazonListDeleteOne struct {
	ald *AmazonListDelete
}

// Where appends a list predicates to the AmazonListDelete builder.
func (aldo *AmazonListDeleteOne) Where(ps ...predicate.AmazonList) *AmazonListDeleteOne {
	aldo.ald.mutation.Where(ps...)
	return aldo
}

// Exec executes the deletion query.
func (aldo *AmazonListDeleteOne) Exec(ctx context.Context) error {
	n, err := aldo.ald.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{amazonlist.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aldo *AmazonListDeleteOne) ExecX(ctx context.Context) {
	if err := aldo.Exec(ctx); err != nil {
		panic(err)
	}
}
