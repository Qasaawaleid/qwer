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
	"github.com/yekta/stablecog/go-apps/database/ent/generation"
	"github.com/yekta/stablecog/go-apps/database/ent/model"
	"github.com/yekta/stablecog/go-apps/database/ent/negativeprompt"
	"github.com/yekta/stablecog/go-apps/database/ent/predicate"
	"github.com/yekta/stablecog/go-apps/database/ent/prompt"
	"github.com/yekta/stablecog/go-apps/database/ent/scheduler"
	"github.com/yekta/stablecog/go-apps/database/ent/user"
)

// GenerationQuery is the builder for querying Generation entities.
type GenerationQuery struct {
	config
	limit              *int
	offset             *int
	unique             *bool
	order              []OrderFunc
	fields             []string
	inters             []Interceptor
	predicates         []predicate.Generation
	withUser           *UserQuery
	withModel          *ModelQuery
	withPrompt         *PromptQuery
	withNegativePrompt *NegativePromptQuery
	withScheduler      *SchedulerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GenerationQuery builder.
func (gq *GenerationQuery) Where(ps ...predicate.Generation) *GenerationQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit the number of records to be returned by this query.
func (gq *GenerationQuery) Limit(limit int) *GenerationQuery {
	gq.limit = &limit
	return gq
}

// Offset to start from.
func (gq *GenerationQuery) Offset(offset int) *GenerationQuery {
	gq.offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GenerationQuery) Unique(unique bool) *GenerationQuery {
	gq.unique = &unique
	return gq
}

// Order specifies how the records should be ordered.
func (gq *GenerationQuery) Order(o ...OrderFunc) *GenerationQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryUser chains the current query on the "user" edge.
func (gq *GenerationQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(generation.Table, generation.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, generation.UserTable, generation.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryModel chains the current query on the "model" edge.
func (gq *GenerationQuery) QueryModel() *ModelQuery {
	query := (&ModelClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(generation.Table, generation.FieldID, selector),
			sqlgraph.To(model.Table, model.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, generation.ModelTable, generation.ModelColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPrompt chains the current query on the "prompt" edge.
func (gq *GenerationQuery) QueryPrompt() *PromptQuery {
	query := (&PromptClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(generation.Table, generation.FieldID, selector),
			sqlgraph.To(prompt.Table, prompt.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, generation.PromptTable, generation.PromptColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNegativePrompt chains the current query on the "negative_prompt" edge.
func (gq *GenerationQuery) QueryNegativePrompt() *NegativePromptQuery {
	query := (&NegativePromptClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(generation.Table, generation.FieldID, selector),
			sqlgraph.To(negativeprompt.Table, negativeprompt.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, generation.NegativePromptTable, generation.NegativePromptColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryScheduler chains the current query on the "scheduler" edge.
func (gq *GenerationQuery) QueryScheduler() *SchedulerQuery {
	query := (&SchedulerClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(generation.Table, generation.FieldID, selector),
			sqlgraph.To(scheduler.Table, scheduler.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, generation.SchedulerTable, generation.SchedulerColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Generation entity from the query.
// Returns a *NotFoundError when no Generation was found.
func (gq *GenerationQuery) First(ctx context.Context) (*Generation, error) {
	nodes, err := gq.Limit(1).All(newQueryContext(ctx, TypeGeneration, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{generation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GenerationQuery) FirstX(ctx context.Context) *Generation {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Generation ID from the query.
// Returns a *NotFoundError when no Generation ID was found.
func (gq *GenerationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gq.Limit(1).IDs(newQueryContext(ctx, TypeGeneration, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{generation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GenerationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Generation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Generation entity is found.
// Returns a *NotFoundError when no Generation entities are found.
func (gq *GenerationQuery) Only(ctx context.Context) (*Generation, error) {
	nodes, err := gq.Limit(2).All(newQueryContext(ctx, TypeGeneration, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{generation.Label}
	default:
		return nil, &NotSingularError{generation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GenerationQuery) OnlyX(ctx context.Context) *Generation {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Generation ID in the query.
// Returns a *NotSingularError when more than one Generation ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GenerationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gq.Limit(2).IDs(newQueryContext(ctx, TypeGeneration, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{generation.Label}
	default:
		err = &NotSingularError{generation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GenerationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Generations.
func (gq *GenerationQuery) All(ctx context.Context) ([]*Generation, error) {
	ctx = newQueryContext(ctx, TypeGeneration, "All")
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Generation, *GenerationQuery]()
	return withInterceptors[[]*Generation](ctx, gq, qr, gq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gq *GenerationQuery) AllX(ctx context.Context) []*Generation {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Generation IDs.
func (gq *GenerationQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeGeneration, "IDs")
	if err := gq.Select(generation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GenerationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GenerationQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeGeneration, "Count")
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gq, querierCount[*GenerationQuery](), gq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GenerationQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GenerationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeGeneration, "Exist")
	switch _, err := gq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GenerationQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GenerationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GenerationQuery) Clone() *GenerationQuery {
	if gq == nil {
		return nil
	}
	return &GenerationQuery{
		config:             gq.config,
		limit:              gq.limit,
		offset:             gq.offset,
		order:              append([]OrderFunc{}, gq.order...),
		inters:             append([]Interceptor{}, gq.inters...),
		predicates:         append([]predicate.Generation{}, gq.predicates...),
		withUser:           gq.withUser.Clone(),
		withModel:          gq.withModel.Clone(),
		withPrompt:         gq.withPrompt.Clone(),
		withNegativePrompt: gq.withNegativePrompt.Clone(),
		withScheduler:      gq.withScheduler.Clone(),
		// clone intermediate query.
		sql:    gq.sql.Clone(),
		path:   gq.path,
		unique: gq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GenerationQuery) WithUser(opts ...func(*UserQuery)) *GenerationQuery {
	query := (&UserClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withUser = query
	return gq
}

// WithModel tells the query-builder to eager-load the nodes that are connected to
// the "model" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GenerationQuery) WithModel(opts ...func(*ModelQuery)) *GenerationQuery {
	query := (&ModelClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withModel = query
	return gq
}

// WithPrompt tells the query-builder to eager-load the nodes that are connected to
// the "prompt" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GenerationQuery) WithPrompt(opts ...func(*PromptQuery)) *GenerationQuery {
	query := (&PromptClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withPrompt = query
	return gq
}

// WithNegativePrompt tells the query-builder to eager-load the nodes that are connected to
// the "negative_prompt" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GenerationQuery) WithNegativePrompt(opts ...func(*NegativePromptQuery)) *GenerationQuery {
	query := (&NegativePromptClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withNegativePrompt = query
	return gq
}

// WithScheduler tells the query-builder to eager-load the nodes that are connected to
// the "scheduler" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GenerationQuery) WithScheduler(opts ...func(*SchedulerQuery)) *GenerationQuery {
	query := (&SchedulerClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withScheduler = query
	return gq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PromptID uuid.UUID `json:"prompt_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Generation.Query().
//		GroupBy(generation.FieldPromptID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gq *GenerationQuery) GroupBy(field string, fields ...string) *GenerationGroupBy {
	gq.fields = append([]string{field}, fields...)
	grbuild := &GenerationGroupBy{build: gq}
	grbuild.flds = &gq.fields
	grbuild.label = generation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PromptID uuid.UUID `json:"prompt_id,omitempty"`
//	}
//
//	client.Generation.Query().
//		Select(generation.FieldPromptID).
//		Scan(ctx, &v)
func (gq *GenerationQuery) Select(fields ...string) *GenerationSelect {
	gq.fields = append(gq.fields, fields...)
	sbuild := &GenerationSelect{GenerationQuery: gq}
	sbuild.label = generation.Label
	sbuild.flds, sbuild.scan = &gq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GenerationSelect configured with the given aggregations.
func (gq *GenerationQuery) Aggregate(fns ...AggregateFunc) *GenerationSelect {
	return gq.Select().Aggregate(fns...)
}

func (gq *GenerationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gq); err != nil {
				return err
			}
		}
	}
	for _, f := range gq.fields {
		if !generation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GenerationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Generation, error) {
	var (
		nodes       = []*Generation{}
		_spec       = gq.querySpec()
		loadedTypes = [5]bool{
			gq.withUser != nil,
			gq.withModel != nil,
			gq.withPrompt != nil,
			gq.withNegativePrompt != nil,
			gq.withScheduler != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Generation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Generation{config: gq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gq.withUser; query != nil {
		if err := gq.loadUser(ctx, query, nodes, nil,
			func(n *Generation, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := gq.withModel; query != nil {
		if err := gq.loadModel(ctx, query, nodes, nil,
			func(n *Generation, e *Model) { n.Edges.Model = e }); err != nil {
			return nil, err
		}
	}
	if query := gq.withPrompt; query != nil {
		if err := gq.loadPrompt(ctx, query, nodes, nil,
			func(n *Generation, e *Prompt) { n.Edges.Prompt = e }); err != nil {
			return nil, err
		}
	}
	if query := gq.withNegativePrompt; query != nil {
		if err := gq.loadNegativePrompt(ctx, query, nodes, nil,
			func(n *Generation, e *NegativePrompt) { n.Edges.NegativePrompt = e }); err != nil {
			return nil, err
		}
	}
	if query := gq.withScheduler; query != nil {
		if err := gq.loadScheduler(ctx, query, nodes, nil,
			func(n *Generation, e *Scheduler) { n.Edges.Scheduler = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gq *GenerationQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Generation, init func(*Generation), assign func(*Generation, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Generation)
	for i := range nodes {
		if nodes[i].UserID == nil {
			continue
		}
		fk := *nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (gq *GenerationQuery) loadModel(ctx context.Context, query *ModelQuery, nodes []*Generation, init func(*Generation), assign func(*Generation, *Model)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Generation)
	for i := range nodes {
		fk := nodes[i].ModelID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(model.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "model_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (gq *GenerationQuery) loadPrompt(ctx context.Context, query *PromptQuery, nodes []*Generation, init func(*Generation), assign func(*Generation, *Prompt)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Generation)
	for i := range nodes {
		if nodes[i].PromptID == nil {
			continue
		}
		fk := *nodes[i].PromptID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(prompt.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "prompt_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (gq *GenerationQuery) loadNegativePrompt(ctx context.Context, query *NegativePromptQuery, nodes []*Generation, init func(*Generation), assign func(*Generation, *NegativePrompt)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Generation)
	for i := range nodes {
		if nodes[i].NegativePromptID == nil {
			continue
		}
		fk := *nodes[i].NegativePromptID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(negativeprompt.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "negative_prompt_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (gq *GenerationQuery) loadScheduler(ctx context.Context, query *SchedulerQuery, nodes []*Generation, init func(*Generation), assign func(*Generation, *Scheduler)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Generation)
	for i := range nodes {
		fk := nodes[i].SchedulerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(scheduler.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "scheduler_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gq *GenerationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	_spec.Node.Columns = gq.fields
	if len(gq.fields) > 0 {
		_spec.Unique = gq.unique != nil && *gq.unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GenerationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   generation.Table,
			Columns: generation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: generation.FieldID,
			},
		},
		From:   gq.sql,
		Unique: true,
	}
	if unique := gq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, generation.FieldID)
		for i := range fields {
			if fields[i] != generation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GenerationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(generation.Table)
	columns := gq.fields
	if len(columns) == 0 {
		columns = generation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.unique != nil && *gq.unique {
		selector.Distinct()
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GenerationGroupBy is the group-by builder for Generation entities.
type GenerationGroupBy struct {
	selector
	build *GenerationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GenerationGroupBy) Aggregate(fns ...AggregateFunc) *GenerationGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the selector query and scans the result into the given value.
func (ggb *GenerationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeGeneration, "GroupBy")
	if err := ggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GenerationQuery, *GenerationGroupBy](ctx, ggb.build, ggb, ggb.build.inters, v)
}

func (ggb *GenerationGroupBy) sqlScan(ctx context.Context, root *GenerationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ggb.flds)+len(ggb.fns))
		for _, f := range *ggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GenerationSelect is the builder for selecting fields of Generation entities.
type GenerationSelect struct {
	*GenerationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GenerationSelect) Aggregate(fns ...AggregateFunc) *GenerationSelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GenerationSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeGeneration, "Select")
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GenerationQuery, *GenerationSelect](ctx, gs.GenerationQuery, gs, gs.inters, v)
}

func (gs *GenerationSelect) sqlScan(ctx context.Context, root *GenerationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
