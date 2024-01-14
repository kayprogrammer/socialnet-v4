// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kayprogrammer/socialnet-v4/ent/predicate"
	"github.com/kayprogrammer/socialnet-v4/ent/reply"
)

// ReplyDelete is the builder for deleting a Reply entity.
type ReplyDelete struct {
	config
	hooks    []Hook
	mutation *ReplyMutation
}

// Where appends a list predicates to the ReplyDelete builder.
func (rd *ReplyDelete) Where(ps ...predicate.Reply) *ReplyDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *ReplyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rd.sqlExec, rd.mutation, rd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *ReplyDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *ReplyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(reply.Table, sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUUID))
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

// ReplyDeleteOne is the builder for deleting a single Reply entity.
type ReplyDeleteOne struct {
	rd *ReplyDelete
}

// Where appends a list predicates to the ReplyDelete builder.
func (rdo *ReplyDeleteOne) Where(ps ...predicate.Reply) *ReplyDeleteOne {
	rdo.rd.mutation.Where(ps...)
	return rdo
}

// Exec executes the deletion query.
func (rdo *ReplyDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{reply.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *ReplyDeleteOne) ExecX(ctx context.Context) {
	if err := rdo.Exec(ctx); err != nil {
		panic(err)
	}
}
