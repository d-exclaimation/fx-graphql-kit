// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/d-exclaimation/fx-graphql-kit/ent/predicate"
	"github.com/d-exclaimation/fx-graphql-kit/ent/thought"
)

// ThoughtUpdate is the builder for updating Thought entities.
type ThoughtUpdate struct {
	config
	hooks    []Hook
	mutation *ThoughtMutation
}

// Where adds a new predicate for the ThoughtUpdate builder.
func (tu *ThoughtUpdate) Where(ps ...predicate.Thought) *ThoughtUpdate {
	tu.mutation.predicates = append(tu.mutation.predicates, ps...)
	return tu
}

// SetTitle sets the "title" field.
func (tu *ThoughtUpdate) SetTitle(s string) *ThoughtUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetBody sets the "body" field.
func (tu *ThoughtUpdate) SetBody(s string) *ThoughtUpdate {
	tu.mutation.SetBody(s)
	return tu
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (tu *ThoughtUpdate) SetNillableBody(s *string) *ThoughtUpdate {
	if s != nil {
		tu.SetBody(*s)
	}
	return tu
}

// SetImageURL sets the "imageURL" field.
func (tu *ThoughtUpdate) SetImageURL(s string) *ThoughtUpdate {
	tu.mutation.SetImageURL(s)
	return tu
}

// SetUserId sets the "userId" field.
func (tu *ThoughtUpdate) SetUserId(i int64) *ThoughtUpdate {
	tu.mutation.ResetUserId()
	tu.mutation.SetUserId(i)
	return tu
}

// AddUserId adds i to the "userId" field.
func (tu *ThoughtUpdate) AddUserId(i int64) *ThoughtUpdate {
	tu.mutation.AddUserId(i)
	return tu
}

// Mutation returns the ThoughtMutation object of the builder.
func (tu *ThoughtUpdate) Mutation() *ThoughtMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *ThoughtUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThoughtMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *ThoughtUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *ThoughtUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *ThoughtUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *ThoughtUpdate) check() error {
	if v, ok := tu.mutation.Title(); ok {
		if err := thought.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Body(); ok {
		if err := thought.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf("ent: validator failed for field \"body\": %w", err)}
		}
	}
	return nil
}

func (tu *ThoughtUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   thought.Table,
			Columns: thought.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: thought.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thought.FieldTitle,
		})
	}
	if value, ok := tu.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thought.FieldBody,
		})
	}
	if value, ok := tu.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thought.FieldImageURL,
		})
	}
	if value, ok := tu.mutation.UserId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: thought.FieldUserId,
		})
	}
	if value, ok := tu.mutation.AddedUserId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: thought.FieldUserId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thought.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ThoughtUpdateOne is the builder for updating a single Thought entity.
type ThoughtUpdateOne struct {
	config
	hooks    []Hook
	mutation *ThoughtMutation
}

// SetTitle sets the "title" field.
func (tuo *ThoughtUpdateOne) SetTitle(s string) *ThoughtUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetBody sets the "body" field.
func (tuo *ThoughtUpdateOne) SetBody(s string) *ThoughtUpdateOne {
	tuo.mutation.SetBody(s)
	return tuo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (tuo *ThoughtUpdateOne) SetNillableBody(s *string) *ThoughtUpdateOne {
	if s != nil {
		tuo.SetBody(*s)
	}
	return tuo
}

// SetImageURL sets the "imageURL" field.
func (tuo *ThoughtUpdateOne) SetImageURL(s string) *ThoughtUpdateOne {
	tuo.mutation.SetImageURL(s)
	return tuo
}

// SetUserId sets the "userId" field.
func (tuo *ThoughtUpdateOne) SetUserId(i int64) *ThoughtUpdateOne {
	tuo.mutation.ResetUserId()
	tuo.mutation.SetUserId(i)
	return tuo
}

// AddUserId adds i to the "userId" field.
func (tuo *ThoughtUpdateOne) AddUserId(i int64) *ThoughtUpdateOne {
	tuo.mutation.AddUserId(i)
	return tuo
}

// Mutation returns the ThoughtMutation object of the builder.
func (tuo *ThoughtUpdateOne) Mutation() *ThoughtMutation {
	return tuo.mutation
}

// Save executes the query and returns the updated Thought entity.
func (tuo *ThoughtUpdateOne) Save(ctx context.Context) (*Thought, error) {
	var (
		err  error
		node *Thought
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThoughtMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *ThoughtUpdateOne) SaveX(ctx context.Context) *Thought {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *ThoughtUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *ThoughtUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *ThoughtUpdateOne) check() error {
	if v, ok := tuo.mutation.Title(); ok {
		if err := thought.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Body(); ok {
		if err := thought.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf("ent: validator failed for field \"body\": %w", err)}
		}
	}
	return nil
}

func (tuo *ThoughtUpdateOne) sqlSave(ctx context.Context) (_node *Thought, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   thought.Table,
			Columns: thought.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: thought.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Thought.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thought.FieldTitle,
		})
	}
	if value, ok := tuo.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thought.FieldBody,
		})
	}
	if value, ok := tuo.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thought.FieldImageURL,
		})
	}
	if value, ok := tuo.mutation.UserId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: thought.FieldUserId,
		})
	}
	if value, ok := tuo.mutation.AddedUserId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: thought.FieldUserId,
		})
	}
	_node = &Thought{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thought.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
