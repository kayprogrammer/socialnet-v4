// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent/comment"
	"github.com/kayprogrammer/socialnet-v4/ent/notification"
	"github.com/kayprogrammer/socialnet-v4/ent/post"
	"github.com/kayprogrammer/socialnet-v4/ent/predicate"
	"github.com/kayprogrammer/socialnet-v4/ent/reply"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
)

// NotificationUpdate is the builder for updating Notification entities.
type NotificationUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationMutation
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nu *NotificationUpdate) Where(ps ...predicate.Notification) *NotificationUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetCreatedAt sets the "created_at" field.
func (nu *NotificationUpdate) SetCreatedAt(t time.Time) *NotificationUpdate {
	nu.mutation.SetCreatedAt(t)
	return nu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableCreatedAt(t *time.Time) *NotificationUpdate {
	if t != nil {
		nu.SetCreatedAt(*t)
	}
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NotificationUpdate) SetUpdatedAt(t time.Time) *NotificationUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// SetSenderID sets the "sender_id" field.
func (nu *NotificationUpdate) SetSenderID(u uuid.UUID) *NotificationUpdate {
	nu.mutation.SetSenderID(u)
	return nu
}

// SetNillableSenderID sets the "sender_id" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableSenderID(u *uuid.UUID) *NotificationUpdate {
	if u != nil {
		nu.SetSenderID(*u)
	}
	return nu
}

// ClearSenderID clears the value of the "sender_id" field.
func (nu *NotificationUpdate) ClearSenderID() *NotificationUpdate {
	nu.mutation.ClearSenderID()
	return nu
}

// SetNtype sets the "ntype" field.
func (nu *NotificationUpdate) SetNtype(n notification.Ntype) *NotificationUpdate {
	nu.mutation.SetNtype(n)
	return nu
}

// SetNillableNtype sets the "ntype" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableNtype(n *notification.Ntype) *NotificationUpdate {
	if n != nil {
		nu.SetNtype(*n)
	}
	return nu
}

// SetPostID sets the "post_id" field.
func (nu *NotificationUpdate) SetPostID(u uuid.UUID) *NotificationUpdate {
	nu.mutation.SetPostID(u)
	return nu
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillablePostID(u *uuid.UUID) *NotificationUpdate {
	if u != nil {
		nu.SetPostID(*u)
	}
	return nu
}

// ClearPostID clears the value of the "post_id" field.
func (nu *NotificationUpdate) ClearPostID() *NotificationUpdate {
	nu.mutation.ClearPostID()
	return nu
}

// SetCommentID sets the "comment_id" field.
func (nu *NotificationUpdate) SetCommentID(u uuid.UUID) *NotificationUpdate {
	nu.mutation.SetCommentID(u)
	return nu
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableCommentID(u *uuid.UUID) *NotificationUpdate {
	if u != nil {
		nu.SetCommentID(*u)
	}
	return nu
}

// ClearCommentID clears the value of the "comment_id" field.
func (nu *NotificationUpdate) ClearCommentID() *NotificationUpdate {
	nu.mutation.ClearCommentID()
	return nu
}

// SetReplyID sets the "reply_id" field.
func (nu *NotificationUpdate) SetReplyID(u uuid.UUID) *NotificationUpdate {
	nu.mutation.SetReplyID(u)
	return nu
}

// SetNillableReplyID sets the "reply_id" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableReplyID(u *uuid.UUID) *NotificationUpdate {
	if u != nil {
		nu.SetReplyID(*u)
	}
	return nu
}

// ClearReplyID clears the value of the "reply_id" field.
func (nu *NotificationUpdate) ClearReplyID() *NotificationUpdate {
	nu.mutation.ClearReplyID()
	return nu
}

// SetText sets the "text" field.
func (nu *NotificationUpdate) SetText(s string) *NotificationUpdate {
	nu.mutation.SetText(s)
	return nu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableText(s *string) *NotificationUpdate {
	if s != nil {
		nu.SetText(*s)
	}
	return nu
}

// ClearText clears the value of the "text" field.
func (nu *NotificationUpdate) ClearText() *NotificationUpdate {
	nu.mutation.ClearText()
	return nu
}

// SetSender sets the "sender" edge to the User entity.
func (nu *NotificationUpdate) SetSender(u *User) *NotificationUpdate {
	return nu.SetSenderID(u.ID)
}

// AddReceiverIDs adds the "receivers" edge to the User entity by IDs.
func (nu *NotificationUpdate) AddReceiverIDs(ids ...uuid.UUID) *NotificationUpdate {
	nu.mutation.AddReceiverIDs(ids...)
	return nu
}

// AddReceivers adds the "receivers" edges to the User entity.
func (nu *NotificationUpdate) AddReceivers(u ...*User) *NotificationUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nu.AddReceiverIDs(ids...)
}

// SetPost sets the "post" edge to the Post entity.
func (nu *NotificationUpdate) SetPost(p *Post) *NotificationUpdate {
	return nu.SetPostID(p.ID)
}

// SetComment sets the "comment" edge to the Comment entity.
func (nu *NotificationUpdate) SetComment(c *Comment) *NotificationUpdate {
	return nu.SetCommentID(c.ID)
}

// SetReply sets the "reply" edge to the Reply entity.
func (nu *NotificationUpdate) SetReply(r *Reply) *NotificationUpdate {
	return nu.SetReplyID(r.ID)
}

// AddReadByIDs adds the "read_by" edge to the User entity by IDs.
func (nu *NotificationUpdate) AddReadByIDs(ids ...uuid.UUID) *NotificationUpdate {
	nu.mutation.AddReadByIDs(ids...)
	return nu
}

// AddReadBy adds the "read_by" edges to the User entity.
func (nu *NotificationUpdate) AddReadBy(u ...*User) *NotificationUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nu.AddReadByIDs(ids...)
}

// Mutation returns the NotificationMutation object of the builder.
func (nu *NotificationUpdate) Mutation() *NotificationMutation {
	return nu.mutation
}

// ClearSender clears the "sender" edge to the User entity.
func (nu *NotificationUpdate) ClearSender() *NotificationUpdate {
	nu.mutation.ClearSender()
	return nu
}

// ClearReceivers clears all "receivers" edges to the User entity.
func (nu *NotificationUpdate) ClearReceivers() *NotificationUpdate {
	nu.mutation.ClearReceivers()
	return nu
}

// RemoveReceiverIDs removes the "receivers" edge to User entities by IDs.
func (nu *NotificationUpdate) RemoveReceiverIDs(ids ...uuid.UUID) *NotificationUpdate {
	nu.mutation.RemoveReceiverIDs(ids...)
	return nu
}

// RemoveReceivers removes "receivers" edges to User entities.
func (nu *NotificationUpdate) RemoveReceivers(u ...*User) *NotificationUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nu.RemoveReceiverIDs(ids...)
}

// ClearPost clears the "post" edge to the Post entity.
func (nu *NotificationUpdate) ClearPost() *NotificationUpdate {
	nu.mutation.ClearPost()
	return nu
}

// ClearComment clears the "comment" edge to the Comment entity.
func (nu *NotificationUpdate) ClearComment() *NotificationUpdate {
	nu.mutation.ClearComment()
	return nu
}

// ClearReply clears the "reply" edge to the Reply entity.
func (nu *NotificationUpdate) ClearReply() *NotificationUpdate {
	nu.mutation.ClearReply()
	return nu
}

// ClearReadBy clears all "read_by" edges to the User entity.
func (nu *NotificationUpdate) ClearReadBy() *NotificationUpdate {
	nu.mutation.ClearReadBy()
	return nu
}

// RemoveReadByIDs removes the "read_by" edge to User entities by IDs.
func (nu *NotificationUpdate) RemoveReadByIDs(ids ...uuid.UUID) *NotificationUpdate {
	nu.mutation.RemoveReadByIDs(ids...)
	return nu
}

// RemoveReadBy removes "read_by" edges to User entities.
func (nu *NotificationUpdate) RemoveReadBy(u ...*User) *NotificationUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nu.RemoveReadByIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NotificationUpdate) Save(ctx context.Context) (int, error) {
	nu.defaults()
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NotificationUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NotificationUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NotificationUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NotificationUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		v := notification.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nu *NotificationUpdate) check() error {
	if v, ok := nu.mutation.Ntype(); ok {
		if err := notification.NtypeValidator(v); err != nil {
			return &ValidationError{Name: "ntype", err: fmt.Errorf(`ent: validator failed for field "Notification.ntype": %w`, err)}
		}
	}
	return nil
}

func (nu *NotificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.CreatedAt(); ok {
		_spec.SetField(notification.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(notification.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nu.mutation.Ntype(); ok {
		_spec.SetField(notification.FieldNtype, field.TypeEnum, value)
	}
	if value, ok := nu.mutation.Text(); ok {
		_spec.SetField(notification.FieldText, field.TypeString, value)
	}
	if nu.mutation.TextCleared() {
		_spec.ClearField(notification.FieldText, field.TypeString)
	}
	if nu.mutation.SenderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.SenderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.ReceiversCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedReceiversIDs(); len(nodes) > 0 && !nu.mutation.ReceiversCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.ReceiversIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.PostCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.PostIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.CommentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.CommentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.ReplyCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.ReplyIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.ReadByCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedReadByIDs(); len(nodes) > 0 && !nu.mutation.ReadByCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.ReadByIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NotificationUpdateOne is the builder for updating a single Notification entity.
type NotificationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationMutation
}

// SetCreatedAt sets the "created_at" field.
func (nuo *NotificationUpdateOne) SetCreatedAt(t time.Time) *NotificationUpdateOne {
	nuo.mutation.SetCreatedAt(t)
	return nuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableCreatedAt(t *time.Time) *NotificationUpdateOne {
	if t != nil {
		nuo.SetCreatedAt(*t)
	}
	return nuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NotificationUpdateOne) SetUpdatedAt(t time.Time) *NotificationUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// SetSenderID sets the "sender_id" field.
func (nuo *NotificationUpdateOne) SetSenderID(u uuid.UUID) *NotificationUpdateOne {
	nuo.mutation.SetSenderID(u)
	return nuo
}

// SetNillableSenderID sets the "sender_id" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableSenderID(u *uuid.UUID) *NotificationUpdateOne {
	if u != nil {
		nuo.SetSenderID(*u)
	}
	return nuo
}

// ClearSenderID clears the value of the "sender_id" field.
func (nuo *NotificationUpdateOne) ClearSenderID() *NotificationUpdateOne {
	nuo.mutation.ClearSenderID()
	return nuo
}

// SetNtype sets the "ntype" field.
func (nuo *NotificationUpdateOne) SetNtype(n notification.Ntype) *NotificationUpdateOne {
	nuo.mutation.SetNtype(n)
	return nuo
}

// SetNillableNtype sets the "ntype" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableNtype(n *notification.Ntype) *NotificationUpdateOne {
	if n != nil {
		nuo.SetNtype(*n)
	}
	return nuo
}

// SetPostID sets the "post_id" field.
func (nuo *NotificationUpdateOne) SetPostID(u uuid.UUID) *NotificationUpdateOne {
	nuo.mutation.SetPostID(u)
	return nuo
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillablePostID(u *uuid.UUID) *NotificationUpdateOne {
	if u != nil {
		nuo.SetPostID(*u)
	}
	return nuo
}

// ClearPostID clears the value of the "post_id" field.
func (nuo *NotificationUpdateOne) ClearPostID() *NotificationUpdateOne {
	nuo.mutation.ClearPostID()
	return nuo
}

// SetCommentID sets the "comment_id" field.
func (nuo *NotificationUpdateOne) SetCommentID(u uuid.UUID) *NotificationUpdateOne {
	nuo.mutation.SetCommentID(u)
	return nuo
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableCommentID(u *uuid.UUID) *NotificationUpdateOne {
	if u != nil {
		nuo.SetCommentID(*u)
	}
	return nuo
}

// ClearCommentID clears the value of the "comment_id" field.
func (nuo *NotificationUpdateOne) ClearCommentID() *NotificationUpdateOne {
	nuo.mutation.ClearCommentID()
	return nuo
}

// SetReplyID sets the "reply_id" field.
func (nuo *NotificationUpdateOne) SetReplyID(u uuid.UUID) *NotificationUpdateOne {
	nuo.mutation.SetReplyID(u)
	return nuo
}

// SetNillableReplyID sets the "reply_id" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableReplyID(u *uuid.UUID) *NotificationUpdateOne {
	if u != nil {
		nuo.SetReplyID(*u)
	}
	return nuo
}

// ClearReplyID clears the value of the "reply_id" field.
func (nuo *NotificationUpdateOne) ClearReplyID() *NotificationUpdateOne {
	nuo.mutation.ClearReplyID()
	return nuo
}

// SetText sets the "text" field.
func (nuo *NotificationUpdateOne) SetText(s string) *NotificationUpdateOne {
	nuo.mutation.SetText(s)
	return nuo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableText(s *string) *NotificationUpdateOne {
	if s != nil {
		nuo.SetText(*s)
	}
	return nuo
}

// ClearText clears the value of the "text" field.
func (nuo *NotificationUpdateOne) ClearText() *NotificationUpdateOne {
	nuo.mutation.ClearText()
	return nuo
}

// SetSender sets the "sender" edge to the User entity.
func (nuo *NotificationUpdateOne) SetSender(u *User) *NotificationUpdateOne {
	return nuo.SetSenderID(u.ID)
}

// AddReceiverIDs adds the "receivers" edge to the User entity by IDs.
func (nuo *NotificationUpdateOne) AddReceiverIDs(ids ...uuid.UUID) *NotificationUpdateOne {
	nuo.mutation.AddReceiverIDs(ids...)
	return nuo
}

// AddReceivers adds the "receivers" edges to the User entity.
func (nuo *NotificationUpdateOne) AddReceivers(u ...*User) *NotificationUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nuo.AddReceiverIDs(ids...)
}

// SetPost sets the "post" edge to the Post entity.
func (nuo *NotificationUpdateOne) SetPost(p *Post) *NotificationUpdateOne {
	return nuo.SetPostID(p.ID)
}

// SetComment sets the "comment" edge to the Comment entity.
func (nuo *NotificationUpdateOne) SetComment(c *Comment) *NotificationUpdateOne {
	return nuo.SetCommentID(c.ID)
}

// SetReply sets the "reply" edge to the Reply entity.
func (nuo *NotificationUpdateOne) SetReply(r *Reply) *NotificationUpdateOne {
	return nuo.SetReplyID(r.ID)
}

// AddReadByIDs adds the "read_by" edge to the User entity by IDs.
func (nuo *NotificationUpdateOne) AddReadByIDs(ids ...uuid.UUID) *NotificationUpdateOne {
	nuo.mutation.AddReadByIDs(ids...)
	return nuo
}

// AddReadBy adds the "read_by" edges to the User entity.
func (nuo *NotificationUpdateOne) AddReadBy(u ...*User) *NotificationUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nuo.AddReadByIDs(ids...)
}

// Mutation returns the NotificationMutation object of the builder.
func (nuo *NotificationUpdateOne) Mutation() *NotificationMutation {
	return nuo.mutation
}

// ClearSender clears the "sender" edge to the User entity.
func (nuo *NotificationUpdateOne) ClearSender() *NotificationUpdateOne {
	nuo.mutation.ClearSender()
	return nuo
}

// ClearReceivers clears all "receivers" edges to the User entity.
func (nuo *NotificationUpdateOne) ClearReceivers() *NotificationUpdateOne {
	nuo.mutation.ClearReceivers()
	return nuo
}

// RemoveReceiverIDs removes the "receivers" edge to User entities by IDs.
func (nuo *NotificationUpdateOne) RemoveReceiverIDs(ids ...uuid.UUID) *NotificationUpdateOne {
	nuo.mutation.RemoveReceiverIDs(ids...)
	return nuo
}

// RemoveReceivers removes "receivers" edges to User entities.
func (nuo *NotificationUpdateOne) RemoveReceivers(u ...*User) *NotificationUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nuo.RemoveReceiverIDs(ids...)
}

// ClearPost clears the "post" edge to the Post entity.
func (nuo *NotificationUpdateOne) ClearPost() *NotificationUpdateOne {
	nuo.mutation.ClearPost()
	return nuo
}

// ClearComment clears the "comment" edge to the Comment entity.
func (nuo *NotificationUpdateOne) ClearComment() *NotificationUpdateOne {
	nuo.mutation.ClearComment()
	return nuo
}

// ClearReply clears the "reply" edge to the Reply entity.
func (nuo *NotificationUpdateOne) ClearReply() *NotificationUpdateOne {
	nuo.mutation.ClearReply()
	return nuo
}

// ClearReadBy clears all "read_by" edges to the User entity.
func (nuo *NotificationUpdateOne) ClearReadBy() *NotificationUpdateOne {
	nuo.mutation.ClearReadBy()
	return nuo
}

// RemoveReadByIDs removes the "read_by" edge to User entities by IDs.
func (nuo *NotificationUpdateOne) RemoveReadByIDs(ids ...uuid.UUID) *NotificationUpdateOne {
	nuo.mutation.RemoveReadByIDs(ids...)
	return nuo
}

// RemoveReadBy removes "read_by" edges to User entities.
func (nuo *NotificationUpdateOne) RemoveReadBy(u ...*User) *NotificationUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nuo.RemoveReadByIDs(ids...)
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nuo *NotificationUpdateOne) Where(ps ...predicate.Notification) *NotificationUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NotificationUpdateOne) Select(field string, fields ...string) *NotificationUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Notification entity.
func (nuo *NotificationUpdateOne) Save(ctx context.Context) (*Notification, error) {
	nuo.defaults()
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NotificationUpdateOne) SaveX(ctx context.Context) *Notification {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NotificationUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NotificationUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NotificationUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		v := notification.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NotificationUpdateOne) check() error {
	if v, ok := nuo.mutation.Ntype(); ok {
		if err := notification.NtypeValidator(v); err != nil {
			return &ValidationError{Name: "ntype", err: fmt.Errorf(`ent: validator failed for field "Notification.ntype": %w`, err)}
		}
	}
	return nil
}

func (nuo *NotificationUpdateOne) sqlSave(ctx context.Context) (_node *Notification, err error) {
	if err := nuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Notification.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notification.FieldID)
		for _, f := range fields {
			if !notification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.CreatedAt(); ok {
		_spec.SetField(notification.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(notification.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.Ntype(); ok {
		_spec.SetField(notification.FieldNtype, field.TypeEnum, value)
	}
	if value, ok := nuo.mutation.Text(); ok {
		_spec.SetField(notification.FieldText, field.TypeString, value)
	}
	if nuo.mutation.TextCleared() {
		_spec.ClearField(notification.FieldText, field.TypeString)
	}
	if nuo.mutation.SenderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.SenderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.ReceiversCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedReceiversIDs(); len(nodes) > 0 && !nuo.mutation.ReceiversCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.ReceiversIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.PostCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.PostIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.CommentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.CommentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.ReplyCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.ReplyIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.ReadByCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedReadByIDs(); len(nodes) > 0 && !nuo.mutation.ReadByCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.ReadByIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Notification{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
