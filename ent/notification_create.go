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
	"github.com/kayprogrammer/socialnet-v4/ent/notification"
	"github.com/kayprogrammer/socialnet-v4/ent/post"
	"github.com/kayprogrammer/socialnet-v4/ent/reply"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
)

// NotificationCreate is the builder for creating a Notification entity.
type NotificationCreate struct {
	config
	mutation *NotificationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (nc *NotificationCreate) SetCreatedAt(t time.Time) *NotificationCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableCreatedAt(t *time.Time) *NotificationCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updated_at" field.
func (nc *NotificationCreate) SetUpdatedAt(t time.Time) *NotificationCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableUpdatedAt(t *time.Time) *NotificationCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetSenderID sets the "sender_id" field.
func (nc *NotificationCreate) SetSenderID(u uuid.UUID) *NotificationCreate {
	nc.mutation.SetSenderID(u)
	return nc
}

// SetNtype sets the "ntype" field.
func (nc *NotificationCreate) SetNtype(n notification.Ntype) *NotificationCreate {
	nc.mutation.SetNtype(n)
	return nc
}

// SetPostID sets the "post_id" field.
func (nc *NotificationCreate) SetPostID(u uuid.UUID) *NotificationCreate {
	nc.mutation.SetPostID(u)
	return nc
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (nc *NotificationCreate) SetNillablePostID(u *uuid.UUID) *NotificationCreate {
	if u != nil {
		nc.SetPostID(*u)
	}
	return nc
}

// SetCommentID sets the "comment_id" field.
func (nc *NotificationCreate) SetCommentID(u uuid.UUID) *NotificationCreate {
	nc.mutation.SetCommentID(u)
	return nc
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableCommentID(u *uuid.UUID) *NotificationCreate {
	if u != nil {
		nc.SetCommentID(*u)
	}
	return nc
}

// SetReplyID sets the "reply_id" field.
func (nc *NotificationCreate) SetReplyID(u uuid.UUID) *NotificationCreate {
	nc.mutation.SetReplyID(u)
	return nc
}

// SetNillableReplyID sets the "reply_id" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableReplyID(u *uuid.UUID) *NotificationCreate {
	if u != nil {
		nc.SetReplyID(*u)
	}
	return nc
}

// SetText sets the "text" field.
func (nc *NotificationCreate) SetText(s string) *NotificationCreate {
	nc.mutation.SetText(s)
	return nc
}

// SetNillableText sets the "text" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableText(s *string) *NotificationCreate {
	if s != nil {
		nc.SetText(*s)
	}
	return nc
}

// SetID sets the "id" field.
func (nc *NotificationCreate) SetID(u uuid.UUID) *NotificationCreate {
	nc.mutation.SetID(u)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableID(u *uuid.UUID) *NotificationCreate {
	if u != nil {
		nc.SetID(*u)
	}
	return nc
}

// SetSender sets the "sender" edge to the User entity.
func (nc *NotificationCreate) SetSender(u *User) *NotificationCreate {
	return nc.SetSenderID(u.ID)
}

// AddReceiverIDs adds the "receivers" edge to the User entity by IDs.
func (nc *NotificationCreate) AddReceiverIDs(ids ...uuid.UUID) *NotificationCreate {
	nc.mutation.AddReceiverIDs(ids...)
	return nc
}

// AddReceivers adds the "receivers" edges to the User entity.
func (nc *NotificationCreate) AddReceivers(u ...*User) *NotificationCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nc.AddReceiverIDs(ids...)
}

// SetPost sets the "post" edge to the Post entity.
func (nc *NotificationCreate) SetPost(p *Post) *NotificationCreate {
	return nc.SetPostID(p.ID)
}

// SetComment sets the "comment" edge to the Comment entity.
func (nc *NotificationCreate) SetComment(c *Comment) *NotificationCreate {
	return nc.SetCommentID(c.ID)
}

// SetReply sets the "reply" edge to the Reply entity.
func (nc *NotificationCreate) SetReply(r *Reply) *NotificationCreate {
	return nc.SetReplyID(r.ID)
}

// AddReadByIDs adds the "read_by" edge to the User entity by IDs.
func (nc *NotificationCreate) AddReadByIDs(ids ...uuid.UUID) *NotificationCreate {
	nc.mutation.AddReadByIDs(ids...)
	return nc
}

// AddReadBy adds the "read_by" edges to the User entity.
func (nc *NotificationCreate) AddReadBy(u ...*User) *NotificationCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nc.AddReadByIDs(ids...)
}

// Mutation returns the NotificationMutation object of the builder.
func (nc *NotificationCreate) Mutation() *NotificationMutation {
	return nc.mutation
}

// Save creates the Notification in the database.
func (nc *NotificationCreate) Save(ctx context.Context) (*Notification, error) {
	nc.defaults()
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NotificationCreate) SaveX(ctx context.Context) *Notification {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NotificationCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NotificationCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NotificationCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := notification.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		v := notification.DefaultUpdatedAt()
		nc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nc.mutation.ID(); !ok {
		v := notification.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NotificationCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Notification.created_at"`)}
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Notification.updated_at"`)}
	}
	if _, ok := nc.mutation.SenderID(); !ok {
		return &ValidationError{Name: "sender_id", err: errors.New(`ent: missing required field "Notification.sender_id"`)}
	}
	if _, ok := nc.mutation.Ntype(); !ok {
		return &ValidationError{Name: "ntype", err: errors.New(`ent: missing required field "Notification.ntype"`)}
	}
	if v, ok := nc.mutation.Ntype(); ok {
		if err := notification.NtypeValidator(v); err != nil {
			return &ValidationError{Name: "ntype", err: fmt.Errorf(`ent: validator failed for field "Notification.ntype": %w`, err)}
		}
	}
	if _, ok := nc.mutation.SenderID(); !ok {
		return &ValidationError{Name: "sender", err: errors.New(`ent: missing required edge "Notification.sender"`)}
	}
	return nil
}

func (nc *NotificationCreate) sqlSave(ctx context.Context) (*Notification, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
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
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NotificationCreate) createSpec() (*Notification, *sqlgraph.CreateSpec) {
	var (
		_node = &Notification{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(notification.Table, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID))
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(notification.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.SetField(notification.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.Ntype(); ok {
		_spec.SetField(notification.FieldNtype, field.TypeEnum, value)
		_node.Ntype = value
	}
	if value, ok := nc.mutation.Text(); ok {
		_spec.SetField(notification.FieldText, field.TypeString, value)
		_node.Text = &value
	}
	if nodes := nc.mutation.SenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.SenderTable,
			Columns: []string{notification.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SenderID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.ReceiversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.ReceiversTable,
			Columns: notification.ReceiversPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.PostTable,
			Columns: []string{notification.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PostID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.CommentTable,
			Columns: []string{notification.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CommentID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.ReplyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.ReplyTable,
			Columns: []string{notification.ReplyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ReplyID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.ReadByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.ReadByTable,
			Columns: notification.ReadByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NotificationCreateBulk is the builder for creating many Notification entities in bulk.
type NotificationCreateBulk struct {
	config
	err      error
	builders []*NotificationCreate
}

// Save creates the Notification entities in the database.
func (ncb *NotificationCreateBulk) Save(ctx context.Context) ([]*Notification, error) {
	if ncb.err != nil {
		return nil, ncb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Notification, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotificationMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NotificationCreateBulk) SaveX(ctx context.Context) []*Notification {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NotificationCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NotificationCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}