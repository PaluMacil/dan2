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
	"github.com/PaluMacil/dan2/ent/amazonlist"
	"github.com/PaluMacil/dan2/ent/amazonorder"
	"github.com/PaluMacil/dan2/ent/predicate"
)

// AmazonOrderQuery is the builder for querying AmazonOrder entities.
type AmazonOrderQuery struct {
	config
	ctx            *QueryContext
	order          []amazonorder.OrderOption
	inters         []Interceptor
	predicates     []predicate.AmazonOrder
	withAmazonList *AmazonListQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AmazonOrderQuery builder.
func (aoq *AmazonOrderQuery) Where(ps ...predicate.AmazonOrder) *AmazonOrderQuery {
	aoq.predicates = append(aoq.predicates, ps...)
	return aoq
}

// Limit the number of records to be returned by this query.
func (aoq *AmazonOrderQuery) Limit(limit int) *AmazonOrderQuery {
	aoq.ctx.Limit = &limit
	return aoq
}

// Offset to start from.
func (aoq *AmazonOrderQuery) Offset(offset int) *AmazonOrderQuery {
	aoq.ctx.Offset = &offset
	return aoq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aoq *AmazonOrderQuery) Unique(unique bool) *AmazonOrderQuery {
	aoq.ctx.Unique = &unique
	return aoq
}

// Order specifies how the records should be ordered.
func (aoq *AmazonOrderQuery) Order(o ...amazonorder.OrderOption) *AmazonOrderQuery {
	aoq.order = append(aoq.order, o...)
	return aoq
}

// QueryAmazonList chains the current query on the "amazon_list" edge.
func (aoq *AmazonOrderQuery) QueryAmazonList() *AmazonListQuery {
	query := (&AmazonListClient{config: aoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(amazonorder.Table, amazonorder.FieldID, selector),
			sqlgraph.To(amazonlist.Table, amazonlist.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, amazonorder.AmazonListTable, amazonorder.AmazonListPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AmazonOrder entity from the query.
// Returns a *NotFoundError when no AmazonOrder was found.
func (aoq *AmazonOrderQuery) First(ctx context.Context) (*AmazonOrder, error) {
	nodes, err := aoq.Limit(1).All(setContextOp(ctx, aoq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{amazonorder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aoq *AmazonOrderQuery) FirstX(ctx context.Context) *AmazonOrder {
	node, err := aoq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AmazonOrder ID from the query.
// Returns a *NotFoundError when no AmazonOrder ID was found.
func (aoq *AmazonOrderQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aoq.Limit(1).IDs(setContextOp(ctx, aoq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{amazonorder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aoq *AmazonOrderQuery) FirstIDX(ctx context.Context) int {
	id, err := aoq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AmazonOrder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AmazonOrder entity is found.
// Returns a *NotFoundError when no AmazonOrder entities are found.
func (aoq *AmazonOrderQuery) Only(ctx context.Context) (*AmazonOrder, error) {
	nodes, err := aoq.Limit(2).All(setContextOp(ctx, aoq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{amazonorder.Label}
	default:
		return nil, &NotSingularError{amazonorder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aoq *AmazonOrderQuery) OnlyX(ctx context.Context) *AmazonOrder {
	node, err := aoq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AmazonOrder ID in the query.
// Returns a *NotSingularError when more than one AmazonOrder ID is found.
// Returns a *NotFoundError when no entities are found.
func (aoq *AmazonOrderQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aoq.Limit(2).IDs(setContextOp(ctx, aoq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{amazonorder.Label}
	default:
		err = &NotSingularError{amazonorder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aoq *AmazonOrderQuery) OnlyIDX(ctx context.Context) int {
	id, err := aoq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AmazonOrders.
func (aoq *AmazonOrderQuery) All(ctx context.Context) ([]*AmazonOrder, error) {
	ctx = setContextOp(ctx, aoq.ctx, "All")
	if err := aoq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AmazonOrder, *AmazonOrderQuery]()
	return withInterceptors[[]*AmazonOrder](ctx, aoq, qr, aoq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aoq *AmazonOrderQuery) AllX(ctx context.Context) []*AmazonOrder {
	nodes, err := aoq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AmazonOrder IDs.
func (aoq *AmazonOrderQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aoq.ctx.Unique == nil && aoq.path != nil {
		aoq.Unique(true)
	}
	ctx = setContextOp(ctx, aoq.ctx, "IDs")
	if err = aoq.Select(amazonorder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aoq *AmazonOrderQuery) IDsX(ctx context.Context) []int {
	ids, err := aoq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aoq *AmazonOrderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aoq.ctx, "Count")
	if err := aoq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aoq, querierCount[*AmazonOrderQuery](), aoq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aoq *AmazonOrderQuery) CountX(ctx context.Context) int {
	count, err := aoq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aoq *AmazonOrderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aoq.ctx, "Exist")
	switch _, err := aoq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aoq *AmazonOrderQuery) ExistX(ctx context.Context) bool {
	exist, err := aoq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AmazonOrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aoq *AmazonOrderQuery) Clone() *AmazonOrderQuery {
	if aoq == nil {
		return nil
	}
	return &AmazonOrderQuery{
		config:         aoq.config,
		ctx:            aoq.ctx.Clone(),
		order:          append([]amazonorder.OrderOption{}, aoq.order...),
		inters:         append([]Interceptor{}, aoq.inters...),
		predicates:     append([]predicate.AmazonOrder{}, aoq.predicates...),
		withAmazonList: aoq.withAmazonList.Clone(),
		// clone intermediate query.
		sql:  aoq.sql.Clone(),
		path: aoq.path,
	}
}

// WithAmazonList tells the query-builder to eager-load the nodes that are connected to
// the "amazon_list" edge. The optional arguments are used to configure the query builder of the edge.
func (aoq *AmazonOrderQuery) WithAmazonList(opts ...func(*AmazonListQuery)) *AmazonOrderQuery {
	query := (&AmazonListClient{config: aoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aoq.withAmazonList = query
	return aoq
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
//	client.AmazonOrder.Query().
//		GroupBy(amazonorder.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aoq *AmazonOrderQuery) GroupBy(field string, fields ...string) *AmazonOrderGroupBy {
	aoq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AmazonOrderGroupBy{build: aoq}
	grbuild.flds = &aoq.ctx.Fields
	grbuild.label = amazonorder.Label
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
//	client.AmazonOrder.Query().
//		Select(amazonorder.FieldName).
//		Scan(ctx, &v)
func (aoq *AmazonOrderQuery) Select(fields ...string) *AmazonOrderSelect {
	aoq.ctx.Fields = append(aoq.ctx.Fields, fields...)
	sbuild := &AmazonOrderSelect{AmazonOrderQuery: aoq}
	sbuild.label = amazonorder.Label
	sbuild.flds, sbuild.scan = &aoq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AmazonOrderSelect configured with the given aggregations.
func (aoq *AmazonOrderQuery) Aggregate(fns ...AggregateFunc) *AmazonOrderSelect {
	return aoq.Select().Aggregate(fns...)
}

func (aoq *AmazonOrderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aoq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aoq); err != nil {
				return err
			}
		}
	}
	for _, f := range aoq.ctx.Fields {
		if !amazonorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aoq.path != nil {
		prev, err := aoq.path(ctx)
		if err != nil {
			return err
		}
		aoq.sql = prev
	}
	return nil
}

func (aoq *AmazonOrderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AmazonOrder, error) {
	var (
		nodes       = []*AmazonOrder{}
		_spec       = aoq.querySpec()
		loadedTypes = [1]bool{
			aoq.withAmazonList != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AmazonOrder).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AmazonOrder{config: aoq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aoq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aoq.withAmazonList; query != nil {
		if err := aoq.loadAmazonList(ctx, query, nodes,
			func(n *AmazonOrder) { n.Edges.AmazonList = []*AmazonList{} },
			func(n *AmazonOrder, e *AmazonList) { n.Edges.AmazonList = append(n.Edges.AmazonList, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aoq *AmazonOrderQuery) loadAmazonList(ctx context.Context, query *AmazonListQuery, nodes []*AmazonOrder, init func(*AmazonOrder), assign func(*AmazonOrder, *AmazonList)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*AmazonOrder)
	nids := make(map[int]map[*AmazonOrder]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(amazonorder.AmazonListTable)
		s.Join(joinT).On(s.C(amazonlist.FieldID), joinT.C(amazonorder.AmazonListPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(amazonorder.AmazonListPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(amazonorder.AmazonListPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*AmazonOrder]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*AmazonList](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "amazon_list" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (aoq *AmazonOrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aoq.querySpec()
	_spec.Node.Columns = aoq.ctx.Fields
	if len(aoq.ctx.Fields) > 0 {
		_spec.Unique = aoq.ctx.Unique != nil && *aoq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aoq.driver, _spec)
}

func (aoq *AmazonOrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(amazonorder.Table, amazonorder.Columns, sqlgraph.NewFieldSpec(amazonorder.FieldID, field.TypeInt))
	_spec.From = aoq.sql
	if unique := aoq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aoq.path != nil {
		_spec.Unique = true
	}
	if fields := aoq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, amazonorder.FieldID)
		for i := range fields {
			if fields[i] != amazonorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aoq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aoq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aoq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aoq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aoq *AmazonOrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aoq.driver.Dialect())
	t1 := builder.Table(amazonorder.Table)
	columns := aoq.ctx.Fields
	if len(columns) == 0 {
		columns = amazonorder.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aoq.sql != nil {
		selector = aoq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aoq.ctx.Unique != nil && *aoq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aoq.predicates {
		p(selector)
	}
	for _, p := range aoq.order {
		p(selector)
	}
	if offset := aoq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aoq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AmazonOrderGroupBy is the group-by builder for AmazonOrder entities.
type AmazonOrderGroupBy struct {
	selector
	build *AmazonOrderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aogb *AmazonOrderGroupBy) Aggregate(fns ...AggregateFunc) *AmazonOrderGroupBy {
	aogb.fns = append(aogb.fns, fns...)
	return aogb
}

// Scan applies the selector query and scans the result into the given value.
func (aogb *AmazonOrderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aogb.build.ctx, "GroupBy")
	if err := aogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AmazonOrderQuery, *AmazonOrderGroupBy](ctx, aogb.build, aogb, aogb.build.inters, v)
}

func (aogb *AmazonOrderGroupBy) sqlScan(ctx context.Context, root *AmazonOrderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aogb.fns))
	for _, fn := range aogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aogb.flds)+len(aogb.fns))
		for _, f := range *aogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AmazonOrderSelect is the builder for selecting fields of AmazonOrder entities.
type AmazonOrderSelect struct {
	*AmazonOrderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aos *AmazonOrderSelect) Aggregate(fns ...AggregateFunc) *AmazonOrderSelect {
	aos.fns = append(aos.fns, fns...)
	return aos
}

// Scan applies the selector query and scans the result into the given value.
func (aos *AmazonOrderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aos.ctx, "Select")
	if err := aos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AmazonOrderQuery, *AmazonOrderSelect](ctx, aos.AmazonOrderQuery, aos, aos.inters, v)
}

func (aos *AmazonOrderSelect) sqlScan(ctx context.Context, root *AmazonOrderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aos.fns))
	for _, fn := range aos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}