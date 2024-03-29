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
	"github.com/kayprogrammer/socialnet-v4/ent/post"
	"github.com/kayprogrammer/socialnet-v4/ent/predicate"
	"github.com/kayprogrammer/socialnet-v4/ent/reaction"
	"github.com/kayprogrammer/socialnet-v4/ent/reply"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
)

// ReactionUpdate is the builder for updating Reaction entities.
type ReactionUpdate struct {
	config
	hooks    []Hook
	mutation *ReactionMutation
}

// Where appends a list predicates to the ReactionUpdate builder.
func (ru *ReactionUpdate) Where(ps ...predicate.Reaction) *ReactionUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *ReactionUpdate) SetCreatedAt(t time.Time) *ReactionUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *ReactionUpdate) SetNillableCreatedAt(t *time.Time) *ReactionUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *ReactionUpdate) SetUpdatedAt(t time.Time) *ReactionUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetUserID sets the "user_id" field.
func (ru *ReactionUpdate) SetUserID(u uuid.UUID) *ReactionUpdate {
	ru.mutation.SetUserID(u)
	return ru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ru *ReactionUpdate) SetNillableUserID(u *uuid.UUID) *ReactionUpdate {
	if u != nil {
		ru.SetUserID(*u)
	}
	return ru
}

// SetRtype sets the "rtype" field.
func (ru *ReactionUpdate) SetRtype(r reaction.Rtype) *ReactionUpdate {
	ru.mutation.SetRtype(r)
	return ru
}

// SetNillableRtype sets the "rtype" field if the given value is not nil.
func (ru *ReactionUpdate) SetNillableRtype(r *reaction.Rtype) *ReactionUpdate {
	if r != nil {
		ru.SetRtype(*r)
	}
	return ru
}

// SetPostID sets the "post_id" field.
func (ru *ReactionUpdate) SetPostID(u uuid.UUID) *ReactionUpdate {
	ru.mutation.SetPostID(u)
	return ru
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (ru *ReactionUpdate) SetNillablePostID(u *uuid.UUID) *ReactionUpdate {
	if u != nil {
		ru.SetPostID(*u)
	}
	return ru
}

// ClearPostID clears the value of the "post_id" field.
func (ru *ReactionUpdate) ClearPostID() *ReactionUpdate {
	ru.mutation.ClearPostID()
	return ru
}

// SetCommentID sets the "comment_id" field.
func (ru *ReactionUpdate) SetCommentID(u uuid.UUID) *ReactionUpdate {
	ru.mutation.SetCommentID(u)
	return ru
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (ru *ReactionUpdate) SetNillableCommentID(u *uuid.UUID) *ReactionUpdate {
	if u != nil {
		ru.SetCommentID(*u)
	}
	return ru
}

// ClearCommentID clears the value of the "comment_id" field.
func (ru *ReactionUpdate) ClearCommentID() *ReactionUpdate {
	ru.mutation.ClearCommentID()
	return ru
}

// SetReplyID sets the "reply_id" field.
func (ru *ReactionUpdate) SetReplyID(u uuid.UUID) *ReactionUpdate {
	ru.mutation.SetReplyID(u)
	return ru
}

// SetNillableReplyID sets the "reply_id" field if the given value is not nil.
func (ru *ReactionUpdate) SetNillableReplyID(u *uuid.UUID) *ReactionUpdate {
	if u != nil {
		ru.SetReplyID(*u)
	}
	return ru
}

// ClearReplyID clears the value of the "reply_id" field.
func (ru *ReactionUpdate) ClearReplyID() *ReactionUpdate {
	ru.mutation.ClearReplyID()
	return ru
}

// SetUser sets the "user" edge to the User entity.
func (ru *ReactionUpdate) SetUser(u *User) *ReactionUpdate {
	return ru.SetUserID(u.ID)
}

// SetPost sets the "post" edge to the Post entity.
func (ru *ReactionUpdate) SetPost(p *Post) *ReactionUpdate {
	return ru.SetPostID(p.ID)
}

// SetComment sets the "comment" edge to the Comment entity.
func (ru *ReactionUpdate) SetComment(c *Comment) *ReactionUpdate {
	return ru.SetCommentID(c.ID)
}

// SetReply sets the "reply" edge to the Reply entity.
func (ru *ReactionUpdate) SetReply(r *Reply) *ReactionUpdate {
	return ru.SetReplyID(r.ID)
}

// Mutation returns the ReactionMutation object of the builder.
func (ru *ReactionUpdate) Mutation() *ReactionMutation {
	return ru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ru *ReactionUpdate) ClearUser() *ReactionUpdate {
	ru.mutation.ClearUser()
	return ru
}

// ClearPost clears the "post" edge to the Post entity.
func (ru *ReactionUpdate) ClearPost() *ReactionUpdate {
	ru.mutation.ClearPost()
	return ru
}

// ClearComment clears the "comment" edge to the Comment entity.
func (ru *ReactionUpdate) ClearComment() *ReactionUpdate {
	ru.mutation.ClearComment()
	return ru
}

// ClearReply clears the "reply" edge to the Reply entity.
func (ru *ReactionUpdate) ClearReply() *ReactionUpdate {
	ru.mutation.ClearReply()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReactionUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReactionUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReactionUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReactionUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ReactionUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := reaction.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReactionUpdate) check() error {
	if v, ok := ru.mutation.Rtype(); ok {
		if err := reaction.RtypeValidator(v); err != nil {
			return &ValidationError{Name: "rtype", err: fmt.Errorf(`ent: validator failed for field "Reaction.rtype": %w`, err)}
		}
	}
	if _, ok := ru.mutation.UserID(); ru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reaction.user"`)
	}
	return nil
}

func (ru *ReactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(reaction.Table, reaction.Columns, sqlgraph.NewFieldSpec(reaction.FieldID, field.TypeUUID))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.SetField(reaction.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(reaction.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Rtype(); ok {
		_spec.SetField(reaction.FieldRtype, field.TypeEnum, value)
	}
	if ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.UserTable,
			Columns: []string{reaction.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.UserTable,
			Columns: []string{reaction.UserColumn},
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
	if ru.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.PostTable,
			Columns: []string{reaction.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.PostTable,
			Columns: []string{reaction.PostColumn},
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
	if ru.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.CommentTable,
			Columns: []string{reaction.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.CommentTable,
			Columns: []string{reaction.CommentColumn},
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
	if ru.mutation.ReplyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.ReplyTable,
			Columns: []string{reaction.ReplyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ReplyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.ReplyTable,
			Columns: []string{reaction.ReplyColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ReactionUpdateOne is the builder for updating a single Reaction entity.
type ReactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReactionMutation
}

// SetCreatedAt sets the "created_at" field.
func (ruo *ReactionUpdateOne) SetCreatedAt(t time.Time) *ReactionUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *ReactionUpdateOne) SetNillableCreatedAt(t *time.Time) *ReactionUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *ReactionUpdateOne) SetUpdatedAt(t time.Time) *ReactionUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetUserID sets the "user_id" field.
func (ruo *ReactionUpdateOne) SetUserID(u uuid.UUID) *ReactionUpdateOne {
	ruo.mutation.SetUserID(u)
	return ruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ruo *ReactionUpdateOne) SetNillableUserID(u *uuid.UUID) *ReactionUpdateOne {
	if u != nil {
		ruo.SetUserID(*u)
	}
	return ruo
}

// SetRtype sets the "rtype" field.
func (ruo *ReactionUpdateOne) SetRtype(r reaction.Rtype) *ReactionUpdateOne {
	ruo.mutation.SetRtype(r)
	return ruo
}

// SetNillableRtype sets the "rtype" field if the given value is not nil.
func (ruo *ReactionUpdateOne) SetNillableRtype(r *reaction.Rtype) *ReactionUpdateOne {
	if r != nil {
		ruo.SetRtype(*r)
	}
	return ruo
}

// SetPostID sets the "post_id" field.
func (ruo *ReactionUpdateOne) SetPostID(u uuid.UUID) *ReactionUpdateOne {
	ruo.mutation.SetPostID(u)
	return ruo
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (ruo *ReactionUpdateOne) SetNillablePostID(u *uuid.UUID) *ReactionUpdateOne {
	if u != nil {
		ruo.SetPostID(*u)
	}
	return ruo
}

// ClearPostID clears the value of the "post_id" field.
func (ruo *ReactionUpdateOne) ClearPostID() *ReactionUpdateOne {
	ruo.mutation.ClearPostID()
	return ruo
}

// SetCommentID sets the "comment_id" field.
func (ruo *ReactionUpdateOne) SetCommentID(u uuid.UUID) *ReactionUpdateOne {
	ruo.mutation.SetCommentID(u)
	return ruo
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (ruo *ReactionUpdateOne) SetNillableCommentID(u *uuid.UUID) *ReactionUpdateOne {
	if u != nil {
		ruo.SetCommentID(*u)
	}
	return ruo
}

// ClearCommentID clears the value of the "comment_id" field.
func (ruo *ReactionUpdateOne) ClearCommentID() *ReactionUpdateOne {
	ruo.mutation.ClearCommentID()
	return ruo
}

// SetReplyID sets the "reply_id" field.
func (ruo *ReactionUpdateOne) SetReplyID(u uuid.UUID) *ReactionUpdateOne {
	ruo.mutation.SetReplyID(u)
	return ruo
}

// SetNillableReplyID sets the "reply_id" field if the given value is not nil.
func (ruo *ReactionUpdateOne) SetNillableReplyID(u *uuid.UUID) *ReactionUpdateOne {
	if u != nil {
		ruo.SetReplyID(*u)
	}
	return ruo
}

// ClearReplyID clears the value of the "reply_id" field.
func (ruo *ReactionUpdateOne) ClearReplyID() *ReactionUpdateOne {
	ruo.mutation.ClearReplyID()
	return ruo
}

// SetUser sets the "user" edge to the User entity.
func (ruo *ReactionUpdateOne) SetUser(u *User) *ReactionUpdateOne {
	return ruo.SetUserID(u.ID)
}

// SetPost sets the "post" edge to the Post entity.
func (ruo *ReactionUpdateOne) SetPost(p *Post) *ReactionUpdateOne {
	return ruo.SetPostID(p.ID)
}

// SetComment sets the "comment" edge to the Comment entity.
func (ruo *ReactionUpdateOne) SetComment(c *Comment) *ReactionUpdateOne {
	return ruo.SetCommentID(c.ID)
}

// SetReply sets the "reply" edge to the Reply entity.
func (ruo *ReactionUpdateOne) SetReply(r *Reply) *ReactionUpdateOne {
	return ruo.SetReplyID(r.ID)
}

// Mutation returns the ReactionMutation object of the builder.
func (ruo *ReactionUpdateOne) Mutation() *ReactionMutation {
	return ruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ruo *ReactionUpdateOne) ClearUser() *ReactionUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// ClearPost clears the "post" edge to the Post entity.
func (ruo *ReactionUpdateOne) ClearPost() *ReactionUpdateOne {
	ruo.mutation.ClearPost()
	return ruo
}

// ClearComment clears the "comment" edge to the Comment entity.
func (ruo *ReactionUpdateOne) ClearComment() *ReactionUpdateOne {
	ruo.mutation.ClearComment()
	return ruo
}

// ClearReply clears the "reply" edge to the Reply entity.
func (ruo *ReactionUpdateOne) ClearReply() *ReactionUpdateOne {
	ruo.mutation.ClearReply()
	return ruo
}

// Where appends a list predicates to the ReactionUpdate builder.
func (ruo *ReactionUpdateOne) Where(ps ...predicate.Reaction) *ReactionUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReactionUpdateOne) Select(field string, fields ...string) *ReactionUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Reaction entity.
func (ruo *ReactionUpdateOne) Save(ctx context.Context) (*Reaction, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReactionUpdateOne) SaveX(ctx context.Context) *Reaction {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReactionUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReactionUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ReactionUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := reaction.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReactionUpdateOne) check() error {
	if v, ok := ruo.mutation.Rtype(); ok {
		if err := reaction.RtypeValidator(v); err != nil {
			return &ValidationError{Name: "rtype", err: fmt.Errorf(`ent: validator failed for field "Reaction.rtype": %w`, err)}
		}
	}
	if _, ok := ruo.mutation.UserID(); ruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reaction.user"`)
	}
	return nil
}

func (ruo *ReactionUpdateOne) sqlSave(ctx context.Context) (_node *Reaction, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(reaction.Table, reaction.Columns, sqlgraph.NewFieldSpec(reaction.FieldID, field.TypeUUID))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Reaction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reaction.FieldID)
		for _, f := range fields {
			if !reaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != reaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.SetField(reaction.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(reaction.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Rtype(); ok {
		_spec.SetField(reaction.FieldRtype, field.TypeEnum, value)
	}
	if ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.UserTable,
			Columns: []string{reaction.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.UserTable,
			Columns: []string{reaction.UserColumn},
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
	if ruo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.PostTable,
			Columns: []string{reaction.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.PostTable,
			Columns: []string{reaction.PostColumn},
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
	if ruo.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.CommentTable,
			Columns: []string{reaction.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.CommentTable,
			Columns: []string{reaction.CommentColumn},
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
	if ruo.mutation.ReplyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.ReplyTable,
			Columns: []string{reaction.ReplyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ReplyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reaction.ReplyTable,
			Columns: []string{reaction.ReplyColumn},
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
	_node = &Reaction{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
