// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kayprogrammer/socialnet-v4/ent/predicate"
	"github.com/kayprogrammer/socialnet-v4/ent/reaction"
)

// ReactionDelete is the builder for deleting a Reaction entity.
type ReactionDelete struct {
	config
	hooks    []Hook
	mutation *ReactionMutation
}

// Where appends a list predicates to the ReactionDelete builder.
func (rd *ReactionDelete) Where(ps ...predicate.Reaction) *ReactionDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *ReactionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rd.sqlExec, rd.mutation, rd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *ReactionDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *ReactionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(reaction.Table, sqlgraph.NewFieldSpec(reaction.FieldID, field.TypeUUID))
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rd.mutation.done = true
	return affected, err
}

// ReactionDeleteOne is the builder for deleting a single Reaction entity.
type ReactionDeleteOne struct {
	rd *ReactionDelete
}

// Where appends a list predicates to the ReactionDelete builder.
func (rdo *ReactionDeleteOne) Where(ps ...predicate.Reaction) *ReactionDeleteOne {
	rdo.rd.mutation.Where(ps...)
	return rdo
}

// Exec executes the deletion query.
func (rdo *ReactionDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{reaction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *ReactionDeleteOne) ExecX(ctx context.Context) {
	if err := rdo.Exec(ctx); err != nil {
		panic(err)
	}
}
