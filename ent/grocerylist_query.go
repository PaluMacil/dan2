// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PaluMacil/dan2/ent/grocerylist"
	"github.com/PaluMacil/dan2/ent/grocerylistitem"
	"github.com/PaluMacil/dan2/ent/grocerylistshare"
	"github.com/PaluMacil/dan2/ent/predicate"
	"github.com/PaluMacil/dan2/ent/user"
)

// GroceryListQuery is the builder for querying GroceryList entities.
type GroceryListQuery struct {
	config
	ctx                   *QueryContext
	order                 []grocerylist.OrderOption
	inters                []Interceptor
	predicates            []predicate.GroceryList
	withGroceryListItems  *GroceryListItemQuery
	withOwner             *UserQuery
	withGroceryListShares *GroceryListShareQuery
	withFKs               bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroceryListQuery builder.
func (glq *GroceryListQuery) Where(ps ...predicate.GroceryList) *GroceryListQuery {
	glq.predicates = append(glq.predicates, ps...)
	return glq
}

// Limit the number of records to be returned by this query.
func (glq *GroceryListQuery) Limit(limit int) *GroceryListQuery {
	glq.ctx.Limit = &limit
	return glq
}

// Offset to start from.
func (glq *GroceryListQuery) Offset(offset int) *GroceryListQuery {
	glq.ctx.Offset = &offset
	return glq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (glq *GroceryListQuery) Unique(unique bool) *GroceryListQuery {
	glq.ctx.Unique = &unique
	return glq
}

// Order specifies how the records should be ordered.
func (glq *GroceryListQuery) Order(o ...grocerylist.OrderOption) *GroceryListQuery {
	glq.order = append(glq.order, o...)
	return glq
}

// QueryGroceryListItems chains the current query on the "grocery_list_items" edge.
func (glq *GroceryListQuery) QueryGroceryListItems() *GroceryListItemQuery {
	query := (&GroceryListItemClient{config: glq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := glq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := glq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grocerylist.Table, grocerylist.FieldID, selector),
			sqlgraph.To(grocerylistitem.Table, grocerylistitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, grocerylist.GroceryListItemsTable, grocerylist.GroceryListItemsColumn),
		)
		fromU = sqlgraph.SetNeighbors(glq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwner chains the current query on the "owner" edge.
func (glq *GroceryListQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: glq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := glq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := glq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grocerylist.Table, grocerylist.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, grocerylist.OwnerTable, grocerylist.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(glq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroceryListShares chains the current query on the "grocery_list_shares" edge.
func (glq *GroceryListQuery) QueryGroceryListShares() *GroceryListShareQuery {
	query := (&GroceryListShareClient{config: glq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := glq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := glq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grocerylist.Table, grocerylist.FieldID, selector),
			sqlgraph.To(grocerylistshare.Table, grocerylistshare.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, grocerylist.GroceryListSharesTable, grocerylist.GroceryListSharesColumn),
		)
		fromU = sqlgraph.SetNeighbors(glq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroceryList entity from the query.
// Returns a *NotFoundError when no GroceryList was found.
func (glq *GroceryListQuery) First(ctx context.Context) (*GroceryList, error) {
	nodes, err := glq.Limit(1).All(setContextOp(ctx, glq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{grocerylist.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (glq *GroceryListQuery) FirstX(ctx context.Context) *GroceryList {
	node, err := glq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroceryList ID from the query.
// Returns a *NotFoundError when no GroceryList ID was found.
func (glq *GroceryListQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = glq.Limit(1).IDs(setContextOp(ctx, glq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{grocerylist.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (glq *GroceryListQuery) FirstIDX(ctx context.Context) int {
	id, err := glq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroceryList entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroceryList entity is found.
// Returns a *NotFoundError when no GroceryList entities are found.
func (glq *GroceryListQuery) Only(ctx context.Context) (*GroceryList, error) {
	nodes, err := glq.Limit(2).All(setContextOp(ctx, glq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{grocerylist.Label}
	default:
		return nil, &NotSingularError{grocerylist.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (glq *GroceryListQuery) OnlyX(ctx context.Context) *GroceryList {
	node, err := glq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroceryList ID in the query.
// Returns a *NotSingularError when more than one GroceryList ID is found.
// Returns a *NotFoundError when no entities are found.
func (glq *GroceryListQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = glq.Limit(2).IDs(setContextOp(ctx, glq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{grocerylist.Label}
	default:
		err = &NotSingularError{grocerylist.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (glq *GroceryListQuery) OnlyIDX(ctx context.Context) int {
	id, err := glq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroceryLists.
func (glq *GroceryListQuery) All(ctx context.Context) ([]*GroceryList, error) {
	ctx = setContextOp(ctx, glq.ctx, "All")
	if err := glq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroceryList, *GroceryListQuery]()
	return withInterceptors[[]*GroceryList](ctx, glq, qr, glq.inters)
}

// AllX is like All, but panics if an error occurs.
func (glq *GroceryListQuery) AllX(ctx context.Context) []*GroceryList {
	nodes, err := glq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroceryList IDs.
func (glq *GroceryListQuery) IDs(ctx context.Context) (ids []int, err error) {
	if glq.ctx.Unique == nil && glq.path != nil {
		glq.Unique(true)
	}
	ctx = setContextOp(ctx, glq.ctx, "IDs")
	if err = glq.Select(grocerylist.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (glq *GroceryListQuery) IDsX(ctx context.Context) []int {
	ids, err := glq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (glq *GroceryListQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, glq.ctx, "Count")
	if err := glq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, glq, querierCount[*GroceryListQuery](), glq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (glq *GroceryListQuery) CountX(ctx context.Context) int {
	count, err := glq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (glq *GroceryListQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, glq.ctx, "Exist")
	switch _, err := glq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (glq *GroceryListQuery) ExistX(ctx context.Context) bool {
	exist, err := glq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroceryListQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (glq *GroceryListQuery) Clone() *GroceryListQuery {
	if glq == nil {
		return nil
	}
	return &GroceryListQuery{
		config:                glq.config,
		ctx:                   glq.ctx.Clone(),
		order:                 append([]grocerylist.OrderOption{}, glq.order...),
		inters:                append([]Interceptor{}, glq.inters...),
		predicates:            append([]predicate.GroceryList{}, glq.predicates...),
		withGroceryListItems:  glq.withGroceryListItems.Clone(),
		withOwner:             glq.withOwner.Clone(),
		withGroceryListShares: glq.withGroceryListShares.Clone(),
		// clone intermediate query.
		sql:  glq.sql.Clone(),
		path: glq.path,
	}
}

// WithGroceryListItems tells the query-builder to eager-load the nodes that are connected to
// the "grocery_list_items" edge. The optional arguments are used to configure the query builder of the edge.
func (glq *GroceryListQuery) WithGroceryListItems(opts ...func(*GroceryListItemQuery)) *GroceryListQuery {
	query := (&GroceryListItemClient{config: glq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	glq.withGroceryListItems = query
	return glq
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (glq *GroceryListQuery) WithOwner(opts ...func(*UserQuery)) *GroceryListQuery {
	query := (&UserClient{config: glq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	glq.withOwner = query
	return glq
}

// WithGroceryListShares tells the query-builder to eager-load the nodes that are connected to
// the "grocery_list_shares" edge. The optional arguments are used to configure the query builder of the edge.
func (glq *GroceryListQuery) WithGroceryListShares(opts ...func(*GroceryListShareQuery)) *GroceryListQuery {
	query := (&GroceryListShareClient{config: glq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	glq.withGroceryListShares = query
	return glq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroceryList.Query().
//		GroupBy(grocerylist.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (glq *GroceryListQuery) GroupBy(field string, fields ...string) *GroceryListGroupBy {
	glq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroceryListGroupBy{build: glq}
	grbuild.flds = &glq.ctx.Fields
	grbuild.label = grocerylist.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.GroceryList.Query().
//		Select(grocerylist.FieldName).
//		Scan(ctx, &v)
func (glq *GroceryListQuery) Select(fields ...string) *GroceryListSelect {
	glq.ctx.Fields = append(glq.ctx.Fields, fields...)
	sbuild := &GroceryListSelect{GroceryListQuery: glq}
	sbuild.label = grocerylist.Label
	sbuild.flds, sbuild.scan = &glq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroceryListSelect configured with the given aggregations.
func (glq *GroceryListQuery) Aggregate(fns ...AggregateFunc) *GroceryListSelect {
	return glq.Select().Aggregate(fns...)
}

func (glq *GroceryListQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range glq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, glq); err != nil {
				return err
			}
		}
	}
	for _, f := range glq.ctx.Fields {
		if !grocerylist.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if glq.path != nil {
		prev, err := glq.path(ctx)
		if err != nil {
			return err
		}
		glq.sql = prev
	}
	return nil
}

func (glq *GroceryListQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroceryList, error) {
	var (
		nodes       = []*GroceryList{}
		withFKs     = glq.withFKs
		_spec       = glq.querySpec()
		loadedTypes = [3]bool{
			glq.withGroceryListItems != nil,
			glq.withOwner != nil,
			glq.withGroceryListShares != nil,
		}
	)
	if glq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, grocerylist.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroceryList).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroceryList{config: glq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, glq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := glq.withGroceryListItems; query != nil {
		if err := glq.loadGroceryListItems(ctx, query, nodes,
			func(n *GroceryList) { n.Edges.GroceryListItems = []*GroceryListItem{} },
			func(n *GroceryList, e *GroceryListItem) {
				n.Edges.GroceryListItems = append(n.Edges.GroceryListItems, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := glq.withOwner; query != nil {
		if err := glq.loadOwner(ctx, query, nodes, nil,
			func(n *GroceryList, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := glq.withGroceryListShares; query != nil {
		if err := glq.loadGroceryListShares(ctx, query, nodes,
			func(n *GroceryList) { n.Edges.GroceryListShares = []*GroceryListShare{} },
			func(n *GroceryList, e *GroceryListShare) {
				n.Edges.GroceryListShares = append(n.Edges.GroceryListShares, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (glq *GroceryListQuery) loadGroceryListItems(ctx context.Context, query *GroceryListItemQuery, nodes []*GroceryList, init func(*GroceryList), assign func(*GroceryList, *GroceryListItem)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*GroceryList)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.GroceryListItem(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(grocerylist.GroceryListItemsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.grocery_list_grocery_list_items
		if fk == nil {
			return fmt.Errorf(`foreign-key "grocery_list_grocery_list_items" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "grocery_list_grocery_list_items" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (glq *GroceryListQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*GroceryList, init func(*GroceryList), assign func(*GroceryList, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*GroceryList)
	for i := range nodes {
		if nodes[i].user_grocery_lists == nil {
			continue
		}
		fk := *nodes[i].user_grocery_lists
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
			return fmt.Errorf(`unexpected foreign-key "user_grocery_lists" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (glq *GroceryListQuery) loadGroceryListShares(ctx context.Context, query *GroceryListShareQuery, nodes []*GroceryList, init func(*GroceryList), assign func(*GroceryList, *GroceryListShare)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*GroceryList)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.GroceryListShare(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(grocerylist.GroceryListSharesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.grocery_list_grocery_list_shares
		if fk == nil {
			return fmt.Errorf(`foreign-key "grocery_list_grocery_list_shares" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "grocery_list_grocery_list_shares" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (glq *GroceryListQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := glq.querySpec()
	_spec.Node.Columns = glq.ctx.Fields
	if len(glq.ctx.Fields) > 0 {
		_spec.Unique = glq.ctx.Unique != nil && *glq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, glq.driver, _spec)
}

func (glq *GroceryListQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(grocerylist.Table, grocerylist.Columns, sqlgraph.NewFieldSpec(grocerylist.FieldID, field.TypeInt))
	_spec.From = glq.sql
	if unique := glq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if glq.path != nil {
		_spec.Unique = true
	}
	if fields := glq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grocerylist.FieldID)
		for i := range fields {
			if fields[i] != grocerylist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := glq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := glq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := glq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := glq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (glq *GroceryListQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(glq.driver.Dialect())
	t1 := builder.Table(grocerylist.Table)
	columns := glq.ctx.Fields
	if len(columns) == 0 {
		columns = grocerylist.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if glq.sql != nil {
		selector = glq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if glq.ctx.Unique != nil && *glq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range glq.predicates {
		p(selector)
	}
	for _, p := range glq.order {
		p(selector)
	}
	if offset := glq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := glq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroceryListGroupBy is the group-by builder for GroceryList entities.
type GroceryListGroupBy struct {
	selector
	build *GroceryListQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (glgb *GroceryListGroupBy) Aggregate(fns ...AggregateFunc) *GroceryListGroupBy {
	glgb.fns = append(glgb.fns, fns...)
	return glgb
}

// Scan applies the selector query and scans the result into the given value.
func (glgb *GroceryListGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, glgb.build.ctx, "GroupBy")
	if err := glgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroceryListQuery, *GroceryListGroupBy](ctx, glgb.build, glgb, glgb.build.inters, v)
}

func (glgb *GroceryListGroupBy) sqlScan(ctx context.Context, root *GroceryListQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(glgb.fns))
	for _, fn := range glgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*glgb.flds)+len(glgb.fns))
		for _, f := range *glgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*glgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := glgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroceryListSelect is the builder for selecting fields of GroceryList entities.
type GroceryListSelect struct {
	*GroceryListQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gls *GroceryListSelect) Aggregate(fns ...AggregateFunc) *GroceryListSelect {
	gls.fns = append(gls.fns, fns...)
	return gls
}

// Scan applies the selector query and scans the result into the given value.
func (gls *GroceryListSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gls.ctx, "Select")
	if err := gls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroceryListQuery, *GroceryListSelect](ctx, gls.GroceryListQuery, gls, gls.inters, v)
}

func (gls *GroceryListSelect) sqlScan(ctx context.Context, root *GroceryListQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gls.fns))
	for _, fn := range gls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
