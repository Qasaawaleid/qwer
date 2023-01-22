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
	"github.com/google/uuid"
	"github.com/stablecog/go-apps/database/ent/generation"
	"github.com/stablecog/go-apps/database/ent/generationmodel"
	"github.com/stablecog/go-apps/database/ent/predicate"
)

// GenerationModelQuery is the builder for querying GenerationModel entities.
type GenerationModelQuery struct {
	config
	limit           *int
	offset          *int
	unique          *bool
	order           []OrderFunc
	fields          []string
	inters          []Interceptor
	predicates      []predicate.GenerationModel
	withGenerations *GenerationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GenerationModelQuery builder.
func (gmq *GenerationModelQuery) Where(ps ...predicate.GenerationModel) *GenerationModelQuery {
	gmq.predicates = append(gmq.predicates, ps...)
	return gmq
}

// Limit the number of records to be returned by this query.
func (gmq *GenerationModelQuery) Limit(limit int) *GenerationModelQuery {
	gmq.limit = &limit
	return gmq
}

// Offset to start from.
func (gmq *GenerationModelQuery) Offset(offset int) *GenerationModelQuery {
	gmq.offset = &offset
	return gmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gmq *GenerationModelQuery) Unique(unique bool) *GenerationModelQuery {
	gmq.unique = &unique
	return gmq
}

// Order specifies how the records should be ordered.
func (gmq *GenerationModelQuery) Order(o ...OrderFunc) *GenerationModelQuery {
	gmq.order = append(gmq.order, o...)
	return gmq
}

// QueryGenerations chains the current query on the "generations" edge.
func (gmq *GenerationModelQuery) QueryGenerations() *GenerationQuery {
	query := (&GenerationClient{config: gmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(generationmodel.Table, generationmodel.FieldID, selector),
			sqlgraph.To(generation.Table, generation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, generationmodel.GenerationsTable, generationmodel.GenerationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(gmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GenerationModel entity from the query.
// Returns a *NotFoundError when no GenerationModel was found.
func (gmq *GenerationModelQuery) First(ctx context.Context) (*GenerationModel, error) {
	nodes, err := gmq.Limit(1).All(newQueryContext(ctx, TypeGenerationModel, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{generationmodel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gmq *GenerationModelQuery) FirstX(ctx context.Context) *GenerationModel {
	node, err := gmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GenerationModel ID from the query.
// Returns a *NotFoundError when no GenerationModel ID was found.
func (gmq *GenerationModelQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gmq.Limit(1).IDs(newQueryContext(ctx, TypeGenerationModel, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{generationmodel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gmq *GenerationModelQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GenerationModel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GenerationModel entity is found.
// Returns a *NotFoundError when no GenerationModel entities are found.
func (gmq *GenerationModelQuery) Only(ctx context.Context) (*GenerationModel, error) {
	nodes, err := gmq.Limit(2).All(newQueryContext(ctx, TypeGenerationModel, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{generationmodel.Label}
	default:
		return nil, &NotSingularError{generationmodel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gmq *GenerationModelQuery) OnlyX(ctx context.Context) *GenerationModel {
	node, err := gmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GenerationModel ID in the query.
// Returns a *NotSingularError when more than one GenerationModel ID is found.
// Returns a *NotFoundError when no entities are found.
func (gmq *GenerationModelQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gmq.Limit(2).IDs(newQueryContext(ctx, TypeGenerationModel, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{generationmodel.Label}
	default:
		err = &NotSingularError{generationmodel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gmq *GenerationModelQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GenerationModels.
func (gmq *GenerationModelQuery) All(ctx context.Context) ([]*GenerationModel, error) {
	ctx = newQueryContext(ctx, TypeGenerationModel, "All")
	if err := gmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GenerationModel, *GenerationModelQuery]()
	return withInterceptors[[]*GenerationModel](ctx, gmq, qr, gmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gmq *GenerationModelQuery) AllX(ctx context.Context) []*GenerationModel {
	nodes, err := gmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GenerationModel IDs.
func (gmq *GenerationModelQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeGenerationModel, "IDs")
	if err := gmq.Select(generationmodel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gmq *GenerationModelQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gmq *GenerationModelQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeGenerationModel, "Count")
	if err := gmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gmq, querierCount[*GenerationModelQuery](), gmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gmq *GenerationModelQuery) CountX(ctx context.Context) int {
	count, err := gmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gmq *GenerationModelQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeGenerationModel, "Exist")
	switch _, err := gmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gmq *GenerationModelQuery) ExistX(ctx context.Context) bool {
	exist, err := gmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GenerationModelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gmq *GenerationModelQuery) Clone() *GenerationModelQuery {
	if gmq == nil {
		return nil
	}
	return &GenerationModelQuery{
		config:          gmq.config,
		limit:           gmq.limit,
		offset:          gmq.offset,
		order:           append([]OrderFunc{}, gmq.order...),
		inters:          append([]Interceptor{}, gmq.inters...),
		predicates:      append([]predicate.GenerationModel{}, gmq.predicates...),
		withGenerations: gmq.withGenerations.Clone(),
		// clone intermediate query.
		sql:    gmq.sql.Clone(),
		path:   gmq.path,
		unique: gmq.unique,
	}
}

// WithGenerations tells the query-builder to eager-load the nodes that are connected to
// the "generations" edge. The optional arguments are used to configure the query builder of the edge.
func (gmq *GenerationModelQuery) WithGenerations(opts ...func(*GenerationQuery)) *GenerationModelQuery {
	query := (&GenerationClient{config: gmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gmq.withGenerations = query
	return gmq
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
//	client.GenerationModel.Query().
//		GroupBy(generationmodel.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gmq *GenerationModelQuery) GroupBy(field string, fields ...string) *GenerationModelGroupBy {
	gmq.fields = append([]string{field}, fields...)
	grbuild := &GenerationModelGroupBy{build: gmq}
	grbuild.flds = &gmq.fields
	grbuild.label = generationmodel.Label
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
//	client.GenerationModel.Query().
//		Select(generationmodel.FieldName).
//		Scan(ctx, &v)
func (gmq *GenerationModelQuery) Select(fields ...string) *GenerationModelSelect {
	gmq.fields = append(gmq.fields, fields...)
	sbuild := &GenerationModelSelect{GenerationModelQuery: gmq}
	sbuild.label = generationmodel.Label
	sbuild.flds, sbuild.scan = &gmq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GenerationModelSelect configured with the given aggregations.
func (gmq *GenerationModelQuery) Aggregate(fns ...AggregateFunc) *GenerationModelSelect {
	return gmq.Select().Aggregate(fns...)
}

func (gmq *GenerationModelQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gmq); err != nil {
				return err
			}
		}
	}
	for _, f := range gmq.fields {
		if !generationmodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gmq.path != nil {
		prev, err := gmq.path(ctx)
		if err != nil {
			return err
		}
		gmq.sql = prev
	}
	return nil
}

func (gmq *GenerationModelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GenerationModel, error) {
	var (
		nodes       = []*GenerationModel{}
		_spec       = gmq.querySpec()
		loadedTypes = [1]bool{
			gmq.withGenerations != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GenerationModel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GenerationModel{config: gmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gmq.withGenerations; query != nil {
		if err := gmq.loadGenerations(ctx, query, nodes,
			func(n *GenerationModel) { n.Edges.Generations = []*Generation{} },
			func(n *GenerationModel, e *Generation) { n.Edges.Generations = append(n.Edges.Generations, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gmq *GenerationModelQuery) loadGenerations(ctx context.Context, query *GenerationQuery, nodes []*GenerationModel, init func(*GenerationModel), assign func(*GenerationModel, *Generation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*GenerationModel)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Generation(func(s *sql.Selector) {
		s.Where(sql.InValues(generationmodel.GenerationsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ModelID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "model_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (gmq *GenerationModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gmq.querySpec()
	_spec.Node.Columns = gmq.fields
	if len(gmq.fields) > 0 {
		_spec.Unique = gmq.unique != nil && *gmq.unique
	}
	return sqlgraph.CountNodes(ctx, gmq.driver, _spec)
}

func (gmq *GenerationModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   generationmodel.Table,
			Columns: generationmodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: generationmodel.FieldID,
			},
		},
		From:   gmq.sql,
		Unique: true,
	}
	if unique := gmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, generationmodel.FieldID)
		for i := range fields {
			if fields[i] != generationmodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gmq *GenerationModelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gmq.driver.Dialect())
	t1 := builder.Table(generationmodel.Table)
	columns := gmq.fields
	if len(columns) == 0 {
		columns = generationmodel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gmq.sql != nil {
		selector = gmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gmq.unique != nil && *gmq.unique {
		selector.Distinct()
	}
	for _, p := range gmq.predicates {
		p(selector)
	}
	for _, p := range gmq.order {
		p(selector)
	}
	if offset := gmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GenerationModelGroupBy is the group-by builder for GenerationModel entities.
type GenerationModelGroupBy struct {
	selector
	build *GenerationModelQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gmgb *GenerationModelGroupBy) Aggregate(fns ...AggregateFunc) *GenerationModelGroupBy {
	gmgb.fns = append(gmgb.fns, fns...)
	return gmgb
}

// Scan applies the selector query and scans the result into the given value.
func (gmgb *GenerationModelGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeGenerationModel, "GroupBy")
	if err := gmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GenerationModelQuery, *GenerationModelGroupBy](ctx, gmgb.build, gmgb, gmgb.build.inters, v)
}

func (gmgb *GenerationModelGroupBy) sqlScan(ctx context.Context, root *GenerationModelQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gmgb.fns))
	for _, fn := range gmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gmgb.flds)+len(gmgb.fns))
		for _, f := range *gmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GenerationModelSelect is the builder for selecting fields of GenerationModel entities.
type GenerationModelSelect struct {
	*GenerationModelQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gms *GenerationModelSelect) Aggregate(fns ...AggregateFunc) *GenerationModelSelect {
	gms.fns = append(gms.fns, fns...)
	return gms
}

// Scan applies the selector query and scans the result into the given value.
func (gms *GenerationModelSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeGenerationModel, "Select")
	if err := gms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GenerationModelQuery, *GenerationModelSelect](ctx, gms.GenerationModelQuery, gms, gms.inters, v)
}

func (gms *GenerationModelSelect) sqlScan(ctx context.Context, root *GenerationModelQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gms.fns))
	for _, fn := range gms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}