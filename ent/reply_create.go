// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent/comment"
	"github.com/kayprogrammer/socialnet-v4/ent/reaction"
	"github.com/kayprogrammer/socialnet-v4/ent/reply"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
)

// ReplyCreate is the builder for creating a Reply entity.
type ReplyCreate struct {
	config
	mutation *ReplyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rc *ReplyCreate) SetCreatedAt(t time.Time) *ReplyCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *ReplyCreate) SetNillableCreatedAt(t *time.Time) *ReplyCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *ReplyCreate) SetUpdatedAt(t time.Time) *ReplyCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *ReplyCreate) SetNillableUpdatedAt(t *time.Time) *ReplyCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetText sets the "text" field.
func (rc *ReplyCreate) SetText(s string) *ReplyCreate {
	rc.mutation.SetText(s)
	return rc
}

// SetSlug sets the "slug" field.
func (rc *ReplyCreate) SetSlug(s string) *ReplyCreate {
	rc.mutation.SetSlug(s)
	return rc
}

// SetAuthorID sets the "author_id" field.
func (rc *ReplyCreate) SetAuthorID(u uuid.UUID) *ReplyCreate {
	rc.mutation.SetAuthorID(u)
	return rc
}

// SetCommentID sets the "comment_id" field.
func (rc *ReplyCreate) SetCommentID(u uuid.UUID) *ReplyCreate {
	rc.mutation.SetCommentID(u)
	return rc
}

// SetID sets the "id" field.
func (rc *ReplyCreate) SetID(u uuid.UUID) *ReplyCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *ReplyCreate) SetNillableID(u *uuid.UUID) *ReplyCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// AddReactionIDs adds the "reactions" edge to the Reaction entity by IDs.
func (rc *ReplyCreate) AddReactionIDs(ids ...uuid.UUID) *ReplyCreate {
	rc.mutation.AddReactionIDs(ids...)
	return rc
}

// AddReactions adds the "reactions" edges to the Reaction entity.
func (rc *ReplyCreate) AddReactions(r ...*Reaction) *ReplyCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddReactionIDs(ids...)
}

// SetAuthor sets the "author" edge to the User entity.
func (rc *ReplyCreate) SetAuthor(u *User) *ReplyCreate {
	return rc.SetAuthorID(u.ID)
}

// SetComment sets the "comment" edge to the Comment entity.
func (rc *ReplyCreate) SetComment(c *Comment) *ReplyCreate {
	return rc.SetCommentID(c.ID)
}

// Mutation returns the ReplyMutation object of the builder.
func (rc *ReplyCreate) Mutation() *ReplyMutation {
	return rc.mutation
}

// Save creates the Reply in the database.
func (rc *ReplyCreate) Save(ctx context.Context) (*Reply, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReplyCreate) SaveX(ctx context.Context) *Reply {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReplyCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReplyCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReplyCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := reply.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := reply.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := reply.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReplyCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Reply.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Reply.updated_at"`)}
	}
	if _, ok := rc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Reply.text"`)}
	}
	if v, ok := rc.mutation.Text(); ok {
		if err := reply.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Reply.text": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Reply.slug"`)}
	}
	if v, ok := rc.mutation.Slug(); ok {
		if err := reply.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Reply.slug": %w`, err)}
		}
	}
	if _, ok := rc.mutation.AuthorID(); !ok {
		return &ValidationError{Name: "author_id", err: errors.New(`ent: missing required field "Reply.author_id"`)}
	}
	if _, ok := rc.mutation.CommentID(); !ok {
		return &ValidationError{Name: "comment_id", err: errors.New(`ent: missing required field "Reply.comment_id"`)}
	}
	if _, ok := rc.mutation.AuthorID(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required edge "Reply.author"`)}
	}
	if _, ok := rc.mutation.CommentID(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required edge "Reply.comment"`)}
	}
	return nil
}

func (rc *ReplyCreate) sqlSave(ctx context.Context) (*Reply, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReplyCreate) createSpec() (*Reply, *sqlgraph.CreateSpec) {
	var (
		_node = &Reply{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(reply.Table, sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUUID))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(reply.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(reply.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.Text(); ok {
		_spec.SetField(reply.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := rc.mutation.Slug(); ok {
		_spec.SetField(reply.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if nodes := rc.mutation.ReactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   reply.ReactionsTable,
			Columns: []string{reply.ReactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reply.AuthorTable,
			Columns: []string{reply.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AuthorID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reply.CommentTable,
			Columns: []string{reply.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CommentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReplyCreateBulk is the builder for creating many Reply entities in bulk.
type ReplyCreateBulk struct {
	config
	err      error
	builders []*ReplyCreate
}

// Save creates the Reply entities in the database.
func (rcb *ReplyCreateBulk) Save(ctx context.Context) ([]*Reply, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Reply, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReplyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReplyCreateBulk) SaveX(ctx context.Context) []*Reply {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReplyCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReplyCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}