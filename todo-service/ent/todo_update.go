// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todoservice/ent/predicate"
	"todoservice/ent/todo"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoUpdate is the builder for updating Todo entities.
type TodoUpdate struct {
	config
	hooks    []Hook
	mutation *TodoMutation
}

// Where appends a list predicates to the TodoUpdate builder.
func (tu *TodoUpdate) Where(ps ...predicate.Todo) *TodoUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TodoUpdate) SetDescription(s string) *TodoUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableDescription(s *string) *TodoUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// SetDueAt sets the "due_at" field.
func (tu *TodoUpdate) SetDueAt(t time.Time) *TodoUpdate {
	tu.mutation.SetDueAt(t)
	return tu
}

// SetNillableDueAt sets the "due_at" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableDueAt(t *time.Time) *TodoUpdate {
	if t != nil {
		tu.SetDueAt(*t)
	}
	return tu
}

// ClearDueAt clears the value of the "due_at" field.
func (tu *TodoUpdate) ClearDueAt() *TodoUpdate {
	tu.mutation.ClearDueAt()
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TodoUpdate) SetCreatedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableCreatedAt(t *time.Time) *TodoUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TodoUpdate) SetUpdatedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableUpdatedAt(t *time.Time) *TodoUpdate {
	if t != nil {
		tu.SetUpdatedAt(*t)
	}
	return tu
}

// Mutation returns the TodoMutation object of the builder.
func (tu *TodoUpdate) Mutation() *TodoMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TodoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TodoUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TodoUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TodoUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TodoUpdate) check() error {
	if v, ok := tu.mutation.Description(); ok {
		if err := todo.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Todo.description": %w`, err)}
		}
	}
	return nil
}

func (tu *TodoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(todo.FieldDescription, field.TypeString, value)
	}
	if value, ok := tu.mutation.DueAt(); ok {
		_spec.SetField(todo.FieldDueAt, field.TypeTime, value)
	}
	if tu.mutation.DueAtCleared() {
		_spec.ClearField(todo.FieldDueAt, field.TypeTime)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(todo.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(todo.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TodoUpdateOne is the builder for updating a single Todo entity.
type TodoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoMutation
}

// SetDescription sets the "description" field.
func (tuo *TodoUpdateOne) SetDescription(s string) *TodoUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableDescription(s *string) *TodoUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// SetDueAt sets the "due_at" field.
func (tuo *TodoUpdateOne) SetDueAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetDueAt(t)
	return tuo
}

// SetNillableDueAt sets the "due_at" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableDueAt(t *time.Time) *TodoUpdateOne {
	if t != nil {
		tuo.SetDueAt(*t)
	}
	return tuo
}

// ClearDueAt clears the value of the "due_at" field.
func (tuo *TodoUpdateOne) ClearDueAt() *TodoUpdateOne {
	tuo.mutation.ClearDueAt()
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TodoUpdateOne) SetCreatedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableCreatedAt(t *time.Time) *TodoUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TodoUpdateOne) SetUpdatedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableUpdatedAt(t *time.Time) *TodoUpdateOne {
	if t != nil {
		tuo.SetUpdatedAt(*t)
	}
	return tuo
}

// Mutation returns the TodoMutation object of the builder.
func (tuo *TodoUpdateOne) Mutation() *TodoMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TodoUpdate builder.
func (tuo *TodoUpdateOne) Where(ps ...predicate.Todo) *TodoUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TodoUpdateOne) Select(field string, fields ...string) *TodoUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Todo entity.
func (tuo *TodoUpdateOne) Save(ctx context.Context) (*Todo, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TodoUpdateOne) SaveX(ctx context.Context) *Todo {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TodoUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TodoUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TodoUpdateOne) check() error {
	if v, ok := tuo.mutation.Description(); ok {
		if err := todo.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Todo.description": %w`, err)}
		}
	}
	return nil
}

func (tuo *TodoUpdateOne) sqlSave(ctx context.Context) (_node *Todo, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Todo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todo.FieldID)
		for _, f := range fields {
			if !todo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(todo.FieldDescription, field.TypeString, value)
	}
	if value, ok := tuo.mutation.DueAt(); ok {
		_spec.SetField(todo.FieldDueAt, field.TypeTime, value)
	}
	if tuo.mutation.DueAtCleared() {
		_spec.ClearField(todo.FieldDueAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(todo.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(todo.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Todo{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
