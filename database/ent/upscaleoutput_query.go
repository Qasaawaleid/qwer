// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/stablecog/go-apps/database/ent/predicate"
	"github.com/stablecog/go-apps/database/ent/upscale"
	"github.com/stablecog/go-apps/database/ent/upscaleoutput"
)

// UpscaleOutputQuery is the builder for querying UpscaleOutput entities.
type UpscaleOutputQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.UpscaleOutput
	withUpscales *UpscaleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UpscaleOutputQuery builder.
func (uoq *UpscaleOutputQuery) Where(ps ...predicate.UpscaleOutput) *UpscaleOutputQuery {
	uoq.predicates = append(uoq.predicates, ps...)
	return uoq
}

// Limit the number of records to be returned by this query.
func (uoq *UpscaleOutputQuery) Limit(limit int) *UpscaleOutputQuery {
	uoq.ctx.Limit = &limit
	return uoq
}

// Offset to start from.
func (uoq *UpscaleOutputQuery) Offset(offset int) *UpscaleOutputQuery {
	uoq.ctx.Offset = &offset
	return uoq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uoq *UpscaleOutputQuery) Unique(unique bool) *UpscaleOutputQuery {
	uoq.ctx.Unique = &unique
	return uoq
}

// Order specifies how the records should be ordered.
func (uoq *UpscaleOutputQuery) Order(o ...OrderFunc) *UpscaleOutputQuery {
	uoq.order = append(uoq.order, o...)
	return uoq
}

// QueryUpscales chains the current query on the "upscales" edge.
func (uoq *UpscaleOutputQuery) QueryUpscales() *UpscaleQuery {
	query := (&UpscaleClient{config: uoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upscaleoutput.Table, upscaleoutput.FieldID, selector),
			sqlgraph.To(upscale.Table, upscale.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, upscaleoutput.UpscalesTable, upscaleoutput.UpscalesColumn),
		)
		fromU = sqlgraph.SetNeighbors(uoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UpscaleOutput entity from the query.
// Returns a *NotFoundError when no UpscaleOutput was found.
func (uoq *UpscaleOutputQuery) First(ctx context.Context) (*UpscaleOutput, error) {
	nodes, err := uoq.Limit(1).All(setContextOp(ctx, uoq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{upscaleoutput.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uoq *UpscaleOutputQuery) FirstX(ctx context.Context) *UpscaleOutput {
	node, err := uoq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UpscaleOutput ID from the query.
// Returns a *NotFoundError when no UpscaleOutput ID was found.
func (uoq *UpscaleOutputQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uoq.Limit(1).IDs(setContextOp(ctx, uoq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{upscaleoutput.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uoq *UpscaleOutputQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := uoq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UpscaleOutput entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UpscaleOutput entity is found.
// Returns a *NotFoundError when no UpscaleOutput entities are found.
func (uoq *UpscaleOutputQuery) Only(ctx context.Context) (*UpscaleOutput, error) {
	nodes, err := uoq.Limit(2).All(setContextOp(ctx, uoq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{upscaleoutput.Label}
	default:
		return nil, &NotSingularError{upscaleoutput.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uoq *UpscaleOutputQuery) OnlyX(ctx context.Context) *UpscaleOutput {
	node, err := uoq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UpscaleOutput ID in the query.
// Returns a *NotSingularError when more than one UpscaleOutput ID is found.
// Returns a *NotFoundError when no entities are found.
func (uoq *UpscaleOutputQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uoq.Limit(2).IDs(setContextOp(ctx, uoq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{upscaleoutput.Label}
	default:
		err = &NotSingularError{upscaleoutput.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uoq *UpscaleOutputQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := uoq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UpscaleOutputs.
func (uoq *UpscaleOutputQuery) All(ctx context.Context) ([]*UpscaleOutput, error) {
	ctx = setContextOp(ctx, uoq.ctx, "All")
	if err := uoq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UpscaleOutput, *UpscaleOutputQuery]()
	return withInterceptors[[]*UpscaleOutput](ctx, uoq, qr, uoq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uoq *UpscaleOutputQuery) AllX(ctx context.Context) []*UpscaleOutput {
	nodes, err := uoq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UpscaleOutput IDs.
func (uoq *UpscaleOutputQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, uoq.ctx, "IDs")
	if err := uoq.Select(upscaleoutput.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uoq *UpscaleOutputQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := uoq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uoq *UpscaleOutputQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uoq.ctx, "Count")
	if err := uoq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uoq, querierCount[*UpscaleOutputQuery](), uoq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uoq *UpscaleOutputQuery) CountX(ctx context.Context) int {
	count, err := uoq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uoq *UpscaleOutputQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uoq.ctx, "Exist")
	switch _, err := uoq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uoq *UpscaleOutputQuery) ExistX(ctx context.Context) bool {
	exist, err := uoq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UpscaleOutputQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uoq *UpscaleOutputQuery) Clone() *UpscaleOutputQuery {
	if uoq == nil {
		return nil
	}
	return &UpscaleOutputQuery{
		config:       uoq.config,
		ctx:          uoq.ctx.Clone(),
		order:        append([]OrderFunc{}, uoq.order...),
		inters:       append([]Interceptor{}, uoq.inters...),
		predicates:   append([]predicate.UpscaleOutput{}, uoq.predicates...),
		withUpscales: uoq.withUpscales.Clone(),
		// clone intermediate query.
		sql:  uoq.sql.Clone(),
		path: uoq.path,
	}
}

// WithUpscales tells the query-builder to eager-load the nodes that are connected to
// the "upscales" edge. The optional arguments are used to configure the query builder of the edge.
func (uoq *UpscaleOutputQuery) WithUpscales(opts ...func(*UpscaleQuery)) *UpscaleOutputQuery {
	query := (&UpscaleClient{config: uoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uoq.withUpscales = query
	return uoq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ImageURL string `json:"image_url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UpscaleOutput.Query().
//		GroupBy(upscaleoutput.FieldImageURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uoq *UpscaleOutputQuery) GroupBy(field string, fields ...string) *UpscaleOutputGroupBy {
	uoq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UpscaleOutputGroupBy{build: uoq}
	grbuild.flds = &uoq.ctx.Fields
	grbuild.label = upscaleoutput.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ImageURL string `json:"image_url,omitempty"`
//	}
//
//	client.UpscaleOutput.Query().
//		Select(upscaleoutput.FieldImageURL).
//		Scan(ctx, &v)
func (uoq *UpscaleOutputQuery) Select(fields ...string) *UpscaleOutputSelect {
	uoq.ctx.Fields = append(uoq.ctx.Fields, fields...)
	sbuild := &UpscaleOutputSelect{UpscaleOutputQuery: uoq}
	sbuild.label = upscaleoutput.Label
	sbuild.flds, sbuild.scan = &uoq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UpscaleOutputSelect configured with the given aggregations.
func (uoq *UpscaleOutputQuery) Aggregate(fns ...AggregateFunc) *UpscaleOutputSelect {
	return uoq.Select().Aggregate(fns...)
}

func (uoq *UpscaleOutputQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uoq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uoq); err != nil {
				return err
			}
		}
	}
	for _, f := range uoq.ctx.Fields {
		if !upscaleoutput.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uoq.path != nil {
		prev, err := uoq.path(ctx)
		if err != nil {
			return err
		}
		uoq.sql = prev
	}
	return nil
}

func (uoq *UpscaleOutputQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UpscaleOutput, error) {
	var (
		nodes       = []*UpscaleOutput{}
		_spec       = uoq.querySpec()
		loadedTypes = [1]bool{
			uoq.withUpscales != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UpscaleOutput).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UpscaleOutput{config: uoq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uoq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uoq.withUpscales; query != nil {
		if err := uoq.loadUpscales(ctx, query, nodes, nil,
			func(n *UpscaleOutput, e *Upscale) { n.Edges.Upscales = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uoq *UpscaleOutputQuery) loadUpscales(ctx context.Context, query *UpscaleQuery, nodes []*UpscaleOutput, init func(*UpscaleOutput), assign func(*UpscaleOutput, *Upscale)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UpscaleOutput)
	for i := range nodes {
		fk := nodes[i].UpscaleID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(upscale.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "upscale_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (uoq *UpscaleOutputQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uoq.querySpec()
	_spec.Node.Columns = uoq.ctx.Fields
	if len(uoq.ctx.Fields) > 0 {
		_spec.Unique = uoq.ctx.Unique != nil && *uoq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uoq.driver, _spec)
}

func (uoq *UpscaleOutputQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   upscaleoutput.Table,
			Columns: upscaleoutput.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscaleoutput.FieldID,
			},
		},
		From:   uoq.sql,
		Unique: true,
	}
	if unique := uoq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uoq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upscaleoutput.FieldID)
		for i := range fields {
			if fields[i] != upscaleoutput.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uoq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uoq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uoq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uoq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uoq *UpscaleOutputQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uoq.driver.Dialect())
	t1 := builder.Table(upscaleoutput.Table)
	columns := uoq.ctx.Fields
	if len(columns) == 0 {
		columns = upscaleoutput.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uoq.sql != nil {
		selector = uoq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uoq.ctx.Unique != nil && *uoq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uoq.predicates {
		p(selector)
	}
	for _, p := range uoq.order {
		p(selector)
	}
	if offset := uoq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uoq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UpscaleOutputGroupBy is the group-by builder for UpscaleOutput entities.
type UpscaleOutputGroupBy struct {
	selector
	build *UpscaleOutputQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uogb *UpscaleOutputGroupBy) Aggregate(fns ...AggregateFunc) *UpscaleOutputGroupBy {
	uogb.fns = append(uogb.fns, fns...)
	return uogb
}

// Scan applies the selector query and scans the result into the given value.
func (uogb *UpscaleOutputGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uogb.build.ctx, "GroupBy")
	if err := uogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpscaleOutputQuery, *UpscaleOutputGroupBy](ctx, uogb.build, uogb, uogb.build.inters, v)
}

func (uogb *UpscaleOutputGroupBy) sqlScan(ctx context.Context, root *UpscaleOutputQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(uogb.fns))
	for _, fn := range uogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*uogb.flds)+len(uogb.fns))
		for _, f := range *uogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*uogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UpscaleOutputSelect is the builder for selecting fields of UpscaleOutput entities.
type UpscaleOutputSelect struct {
	*UpscaleOutputQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uos *UpscaleOutputSelect) Aggregate(fns ...AggregateFunc) *UpscaleOutputSelect {
	uos.fns = append(uos.fns, fns...)
	return uos
}

// Scan applies the selector query and scans the result into the given value.
func (uos *UpscaleOutputSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uos.ctx, "Select")
	if err := uos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpscaleOutputQuery, *UpscaleOutputSelect](ctx, uos.UpscaleOutputQuery, uos, uos.inters, v)
}

func (uos *UpscaleOutputSelect) sqlScan(ctx context.Context, root *UpscaleOutputQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uos.fns))
	for _, fn := range uos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
