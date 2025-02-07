// Code generated by ent, DO NOT EDIT.

package ent

import (
	"GoServer/ent/predicate"
	"GoServer/ent/todos"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodosQuery is the builder for querying Todos entities.
type TodosQuery struct {
	config
	ctx        *QueryContext
	order      []todos.OrderOption
	inters     []Interceptor
	predicates []predicate.Todos
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TodosQuery builder.
func (tq *TodosQuery) Where(ps ...predicate.Todos) *TodosQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TodosQuery) Limit(limit int) *TodosQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TodosQuery) Offset(offset int) *TodosQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TodosQuery) Unique(unique bool) *TodosQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TodosQuery) Order(o ...todos.OrderOption) *TodosQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// First returns the first Todos entity from the query.
// Returns a *NotFoundError when no Todos was found.
func (tq *TodosQuery) First(ctx context.Context) (*Todos, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{todos.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TodosQuery) FirstX(ctx context.Context) *Todos {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Todos ID from the query.
// Returns a *NotFoundError when no Todos ID was found.
func (tq *TodosQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{todos.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TodosQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Todos entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Todos entity is found.
// Returns a *NotFoundError when no Todos entities are found.
func (tq *TodosQuery) Only(ctx context.Context) (*Todos, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{todos.Label}
	default:
		return nil, &NotSingularError{todos.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TodosQuery) OnlyX(ctx context.Context) *Todos {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Todos ID in the query.
// Returns a *NotSingularError when more than one Todos ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TodosQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{todos.Label}
	default:
		err = &NotSingularError{todos.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TodosQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TodosSlice.
func (tq *TodosQuery) All(ctx context.Context) ([]*Todos, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryAll)
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Todos, *TodosQuery]()
	return withInterceptors[[]*Todos](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TodosQuery) AllX(ctx context.Context) []*Todos {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Todos IDs.
func (tq *TodosQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryIDs)
	if err = tq.Select(todos.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TodosQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TodosQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryCount)
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TodosQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TodosQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TodosQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryExist)
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TodosQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TodosQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TodosQuery) Clone() *TodosQuery {
	if tq == nil {
		return nil
	}
	return &TodosQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]todos.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Todos{}, tq.predicates...),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Todos.Query().
//		GroupBy(todos.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TodosQuery) GroupBy(field string, fields ...string) *TodosGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TodosGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = todos.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Todos.Query().
//		Select(todos.FieldTitle).
//		Scan(ctx, &v)
func (tq *TodosQuery) Select(fields ...string) *TodosSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TodosSelect{TodosQuery: tq}
	sbuild.label = todos.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TodosSelect configured with the given aggregations.
func (tq *TodosQuery) Aggregate(fns ...AggregateFunc) *TodosSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TodosQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !todos.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TodosQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Todos, error) {
	var (
		nodes = []*Todos{}
		_spec = tq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Todos).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Todos{config: tq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tq *TodosQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TodosQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(todos.Table, todos.Columns, sqlgraph.NewFieldSpec(todos.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todos.FieldID)
		for i := range fields {
			if fields[i] != todos.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TodosQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(todos.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = todos.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TodosGroupBy is the group-by builder for Todos entities.
type TodosGroupBy struct {
	selector
	build *TodosQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TodosGroupBy) Aggregate(fns ...AggregateFunc) *TodosGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TodosGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, ent.OpQueryGroupBy)
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TodosQuery, *TodosGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TodosGroupBy) sqlScan(ctx context.Context, root *TodosQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TodosSelect is the builder for selecting fields of Todos entities.
type TodosSelect struct {
	*TodosQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TodosSelect) Aggregate(fns ...AggregateFunc) *TodosSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TodosSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, ent.OpQuerySelect)
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TodosQuery, *TodosSelect](ctx, ts.TodosQuery, ts, ts.inters, v)
}

func (ts *TodosSelect) sqlScan(ctx context.Context, root *TodosQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
