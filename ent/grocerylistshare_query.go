// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PaluMacil/dan2/ent/grocerylist"
	"github.com/PaluMacil/dan2/ent/grocerylistshare"
	"github.com/PaluMacil/dan2/ent/predicate"
	"github.com/PaluMacil/dan2/ent/user"
)

// GroceryListShareQuery is the builder for querying GroceryListShare entities.
type GroceryListShareQuery struct {
	config
	ctx             *QueryContext
	order           []grocerylistshare.OrderOption
	inters          []Interceptor
	predicates      []predicate.GroceryListShare
	withUser        *UserQuery
	withGroceryList *GroceryListQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroceryListShareQuery builder.
func (glsq *GroceryListShareQuery) Where(ps ...predicate.GroceryListShare) *GroceryListShareQuery {
	glsq.predicates = append(glsq.predicates, ps...)
	return glsq
}

// Limit the number of records to be returned by this query.
func (glsq *GroceryListShareQuery) Limit(limit int) *GroceryListShareQuery {
	glsq.ctx.Limit = &limit
	return glsq
}

// Offset to start from.
func (glsq *GroceryListShareQuery) Offset(offset int) *GroceryListShareQuery {
	glsq.ctx.Offset = &offset
	return glsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (glsq *GroceryListShareQuery) Unique(unique bool) *GroceryListShareQuery {
	glsq.ctx.Unique = &unique
	return glsq
}

// Order specifies how the records should be ordered.
func (glsq *GroceryListShareQuery) Order(o ...grocerylistshare.OrderOption) *GroceryListShareQuery {
	glsq.order = append(glsq.order, o...)
	return glsq
}

// QueryUser chains the current query on the "user" edge.
func (glsq *GroceryListShareQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: glsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := glsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := glsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grocerylistshare.Table, grocerylistshare.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, grocerylistshare.UserTable, grocerylistshare.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(glsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroceryList chains the current query on the "grocery_list" edge.
func (glsq *GroceryListShareQuery) QueryGroceryList() *GroceryListQuery {
	query := (&GroceryListClient{config: glsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := glsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := glsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grocerylistshare.Table, grocerylistshare.FieldID, selector),
			sqlgraph.To(grocerylist.Table, grocerylist.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, grocerylistshare.GroceryListTable, grocerylistshare.GroceryListColumn),
		)
		fromU = sqlgraph.SetNeighbors(glsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroceryListShare entity from the query.
// Returns a *NotFoundError when no GroceryListShare was found.
func (glsq *GroceryListShareQuery) First(ctx context.Context) (*GroceryListShare, error) {
	nodes, err := glsq.Limit(1).All(setContextOp(ctx, glsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{grocerylistshare.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (glsq *GroceryListShareQuery) FirstX(ctx context.Context) *GroceryListShare {
	node, err := glsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroceryListShare ID from the query.
// Returns a *NotFoundError when no GroceryListShare ID was found.
func (glsq *GroceryListShareQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = glsq.Limit(1).IDs(setContextOp(ctx, glsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{grocerylistshare.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (glsq *GroceryListShareQuery) FirstIDX(ctx context.Context) int {
	id, err := glsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroceryListShare entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroceryListShare entity is found.
// Returns a *NotFoundError when no GroceryListShare entities are found.
func (glsq *GroceryListShareQuery) Only(ctx context.Context) (*GroceryListShare, error) {
	nodes, err := glsq.Limit(2).All(setContextOp(ctx, glsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{grocerylistshare.Label}
	default:
		return nil, &NotSingularError{grocerylistshare.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (glsq *GroceryListShareQuery) OnlyX(ctx context.Context) *GroceryListShare {
	node, err := glsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroceryListShare ID in the query.
// Returns a *NotSingularError when more than one GroceryListShare ID is found.
// Returns a *NotFoundError when no entities are found.
func (glsq *GroceryListShareQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = glsq.Limit(2).IDs(setContextOp(ctx, glsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{grocerylistshare.Label}
	default:
		err = &NotSingularError{grocerylistshare.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (glsq *GroceryListShareQuery) OnlyIDX(ctx context.Context) int {
	id, err := glsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroceryListShares.
func (glsq *GroceryListShareQuery) All(ctx context.Context) ([]*GroceryListShare, error) {
	ctx = setContextOp(ctx, glsq.ctx, "All")
	if err := glsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroceryListShare, *GroceryListShareQuery]()
	return withInterceptors[[]*GroceryListShare](ctx, glsq, qr, glsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (glsq *GroceryListShareQuery) AllX(ctx context.Context) []*GroceryListShare {
	nodes, err := glsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroceryListShare IDs.
func (glsq *GroceryListShareQuery) IDs(ctx context.Context) (ids []int, err error) {
	if glsq.ctx.Unique == nil && glsq.path != nil {
		glsq.Unique(true)
	}
	ctx = setContextOp(ctx, glsq.ctx, "IDs")
	if err = glsq.Select(grocerylistshare.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (glsq *GroceryListShareQuery) IDsX(ctx context.Context) []int {
	ids, err := glsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (glsq *GroceryListShareQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, glsq.ctx, "Count")
	if err := glsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, glsq, querierCount[*GroceryListShareQuery](), glsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (glsq *GroceryListShareQuery) CountX(ctx context.Context) int {
	count, err := glsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (glsq *GroceryListShareQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, glsq.ctx, "Exist")
	switch _, err := glsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (glsq *GroceryListShareQuery) ExistX(ctx context.Context) bool {
	exist, err := glsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroceryListShareQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (glsq *GroceryListShareQuery) Clone() *GroceryListShareQuery {
	if glsq == nil {
		return nil
	}
	return &GroceryListShareQuery{
		config:          glsq.config,
		ctx:             glsq.ctx.Clone(),
		order:           append([]grocerylistshare.OrderOption{}, glsq.order...),
		inters:          append([]Interceptor{}, glsq.inters...),
		predicates:      append([]predicate.GroceryListShare{}, glsq.predicates...),
		withUser:        glsq.withUser.Clone(),
		withGroceryList: glsq.withGroceryList.Clone(),
		// clone intermediate query.
		sql:  glsq.sql.Clone(),
		path: glsq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (glsq *GroceryListShareQuery) WithUser(opts ...func(*UserQuery)) *GroceryListShareQuery {
	query := (&UserClient{config: glsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	glsq.withUser = query
	return glsq
}

// WithGroceryList tells the query-builder to eager-load the nodes that are connected to
// the "grocery_list" edge. The optional arguments are used to configure the query builder of the edge.
func (glsq *GroceryListShareQuery) WithGroceryList(opts ...func(*GroceryListQuery)) *GroceryListShareQuery {
	query := (&GroceryListClient{config: glsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	glsq.withGroceryList = query
	return glsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CanEdit bool `json:"can_edit,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroceryListShare.Query().
//		GroupBy(grocerylistshare.FieldCanEdit).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (glsq *GroceryListShareQuery) GroupBy(field string, fields ...string) *GroceryListShareGroupBy {
	glsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroceryListShareGroupBy{build: glsq}
	grbuild.flds = &glsq.ctx.Fields
	grbuild.label = grocerylistshare.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CanEdit bool `json:"can_edit,omitempty"`
//	}
//
//	client.GroceryListShare.Query().
//		Select(grocerylistshare.FieldCanEdit).
//		Scan(ctx, &v)
func (glsq *GroceryListShareQuery) Select(fields ...string) *GroceryListShareSelect {
	glsq.ctx.Fields = append(glsq.ctx.Fields, fields...)
	sbuild := &GroceryListShareSelect{GroceryListShareQuery: glsq}
	sbuild.label = grocerylistshare.Label
	sbuild.flds, sbuild.scan = &glsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroceryListShareSelect configured with the given aggregations.
func (glsq *GroceryListShareQuery) Aggregate(fns ...AggregateFunc) *GroceryListShareSelect {
	return glsq.Select().Aggregate(fns...)
}

func (glsq *GroceryListShareQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range glsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, glsq); err != nil {
				return err
			}
		}
	}
	for _, f := range glsq.ctx.Fields {
		if !grocerylistshare.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if glsq.path != nil {
		prev, err := glsq.path(ctx)
		if err != nil {
			return err
		}
		glsq.sql = prev
	}
	return nil
}

func (glsq *GroceryListShareQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroceryListShare, error) {
	var (
		nodes       = []*GroceryListShare{}
		withFKs     = glsq.withFKs
		_spec       = glsq.querySpec()
		loadedTypes = [2]bool{
			glsq.withUser != nil,
			glsq.withGroceryList != nil,
		}
	)
	if glsq.withUser != nil || glsq.withGroceryList != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, grocerylistshare.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroceryListShare).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroceryListShare{config: glsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, glsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := glsq.withUser; query != nil {
		if err := glsq.loadUser(ctx, query, nodes, nil,
			func(n *GroceryListShare, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := glsq.withGroceryList; query != nil {
		if err := glsq.loadGroceryList(ctx, query, nodes, nil,
			func(n *GroceryListShare, e *GroceryList) { n.Edges.GroceryList = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (glsq *GroceryListShareQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*GroceryListShare, init func(*GroceryListShare), assign func(*GroceryListShare, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*GroceryListShare)
	for i := range nodes {
		if nodes[i].user_grocery_list_shares == nil {
			continue
		}
		fk := *nodes[i].user_grocery_list_shares
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_grocery_list_shares" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (glsq *GroceryListShareQuery) loadGroceryList(ctx context.Context, query *GroceryListQuery, nodes []*GroceryListShare, init func(*GroceryListShare), assign func(*GroceryListShare, *GroceryList)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*GroceryListShare)
	for i := range nodes {
		if nodes[i].grocery_list_grocery_list_shares == nil {
			continue
		}
		fk := *nodes[i].grocery_list_grocery_list_shares
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(grocerylist.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "grocery_list_grocery_list_shares" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (glsq *GroceryListShareQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := glsq.querySpec()
	_spec.Node.Columns = glsq.ctx.Fields
	if len(glsq.ctx.Fields) > 0 {
		_spec.Unique = glsq.ctx.Unique != nil && *glsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, glsq.driver, _spec)
}

func (glsq *GroceryListShareQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(grocerylistshare.Table, grocerylistshare.Columns, sqlgraph.NewFieldSpec(grocerylistshare.FieldID, field.TypeInt))
	_spec.From = glsq.sql
	if unique := glsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if glsq.path != nil {
		_spec.Unique = true
	}
	if fields := glsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grocerylistshare.FieldID)
		for i := range fields {
			if fields[i] != grocerylistshare.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := glsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := glsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := glsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := glsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (glsq *GroceryListShareQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(glsq.driver.Dialect())
	t1 := builder.Table(grocerylistshare.Table)
	columns := glsq.ctx.Fields
	if len(columns) == 0 {
		columns = grocerylistshare.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if glsq.sql != nil {
		selector = glsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if glsq.ctx.Unique != nil && *glsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range glsq.predicates {
		p(selector)
	}
	for _, p := range glsq.order {
		p(selector)
	}
	if offset := glsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := glsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroceryListShareGroupBy is the group-by builder for GroceryListShare entities.
type GroceryListShareGroupBy struct {
	selector
	build *GroceryListShareQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (glsgb *GroceryListShareGroupBy) Aggregate(fns ...AggregateFunc) *GroceryListShareGroupBy {
	glsgb.fns = append(glsgb.fns, fns...)
	return glsgb
}

// Scan applies the selector query and scans the result into the given value.
func (glsgb *GroceryListShareGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, glsgb.build.ctx, "GroupBy")
	if err := glsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroceryListShareQuery, *GroceryListShareGroupBy](ctx, glsgb.build, glsgb, glsgb.build.inters, v)
}

func (glsgb *GroceryListShareGroupBy) sqlScan(ctx context.Context, root *GroceryListShareQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(glsgb.fns))
	for _, fn := range glsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*glsgb.flds)+len(glsgb.fns))
		for _, f := range *glsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*glsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := glsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroceryListShareSelect is the builder for selecting fields of GroceryListShare entities.
type GroceryListShareSelect struct {
	*GroceryListShareQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (glss *GroceryListShareSelect) Aggregate(fns ...AggregateFunc) *GroceryListShareSelect {
	glss.fns = append(glss.fns, fns...)
	return glss
}

// Scan applies the selector query and scans the result into the given value.
func (glss *GroceryListShareSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, glss.ctx, "Select")
	if err := glss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroceryListShareQuery, *GroceryListShareSelect](ctx, glss.GroceryListShareQuery, glss, glss.inters, v)
}

func (glss *GroceryListShareSelect) sqlScan(ctx context.Context, root *GroceryListShareQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(glss.fns))
	for _, fn := range glss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*glss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := glss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
