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
	"github.com/kayprogrammer/socialnet-v4/ent/friend"
	"github.com/kayprogrammer/socialnet-v4/ent/predicate"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
)

// FriendUpdate is the builder for updating Friend entities.
type FriendUpdate struct {
	config
	hooks    []Hook
	mutation *FriendMutation
}

// Where appends a list predicates to the FriendUpdate builder.
func (fu *FriendUpdate) Where(ps ...predicate.Friend) *FriendUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetCreatedAt sets the "created_at" field.
func (fu *FriendUpdate) SetCreatedAt(t time.Time) *FriendUpdate {
	fu.mutation.SetCreatedAt(t)
	return fu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableCreatedAt(t *time.Time) *FriendUpdate {
	if t != nil {
		fu.SetCreatedAt(*t)
	}
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FriendUpdate) SetUpdatedAt(t time.Time) *FriendUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetRequesterID sets the "requester_id" field.
func (fu *FriendUpdate) SetRequesterID(u uuid.UUID) *FriendUpdate {
	fu.mutation.SetRequesterID(u)
	return fu
}

// SetNillableRequesterID sets the "requester_id" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableRequesterID(u *uuid.UUID) *FriendUpdate {
	if u != nil {
		fu.SetRequesterID(*u)
	}
	return fu
}

// SetRequesteeID sets the "requestee_id" field.
func (fu *FriendUpdate) SetRequesteeID(u uuid.UUID) *FriendUpdate {
	fu.mutation.SetRequesteeID(u)
	return fu
}

// SetNillableRequesteeID sets the "requestee_id" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableRequesteeID(u *uuid.UUID) *FriendUpdate {
	if u != nil {
		fu.SetRequesteeID(*u)
	}
	return fu
}

// SetStatus sets the "status" field.
func (fu *FriendUpdate) SetStatus(f friend.Status) *FriendUpdate {
	fu.mutation.SetStatus(f)
	return fu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableStatus(f *friend.Status) *FriendUpdate {
	if f != nil {
		fu.SetStatus(*f)
	}
	return fu
}

// SetRequester sets the "requester" edge to the User entity.
func (fu *FriendUpdate) SetRequester(u *User) *FriendUpdate {
	return fu.SetRequesterID(u.ID)
}

// SetRequestee sets the "requestee" edge to the User entity.
func (fu *FriendUpdate) SetRequestee(u *User) *FriendUpdate {
	return fu.SetRequesteeID(u.ID)
}

// Mutation returns the FriendMutation object of the builder.
func (fu *FriendUpdate) Mutation() *FriendMutation {
	return fu.mutation
}

// ClearRequester clears the "requester" edge to the User entity.
func (fu *FriendUpdate) ClearRequester() *FriendUpdate {
	fu.mutation.ClearRequester()
	return fu
}

// ClearRequestee clears the "requestee" edge to the User entity.
func (fu *FriendUpdate) ClearRequestee() *FriendUpdate {
	fu.mutation.ClearRequestee()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FriendUpdate) Save(ctx context.Context) (int, error) {
	fu.defaults()
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FriendUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FriendUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FriendUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FriendUpdate) defaults() {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		v := friend.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FriendUpdate) check() error {
	if v, ok := fu.mutation.Status(); ok {
		if err := friend.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Friend.status": %w`, err)}
		}
	}
	if _, ok := fu.mutation.RequesterID(); fu.mutation.RequesterCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friend.requester"`)
	}
	if _, ok := fu.mutation.RequesteeID(); fu.mutation.RequesteeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friend.requestee"`)
	}
	return nil
}

func (fu *FriendUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(friend.Table, friend.Columns, sqlgraph.NewFieldSpec(friend.FieldID, field.TypeUUID))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.CreatedAt(); ok {
		_spec.SetField(friend.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(friend.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fu.mutation.Status(); ok {
		_spec.SetField(friend.FieldStatus, field.TypeEnum, value)
	}
	if fu.mutation.RequesterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.RequesterTable,
			Columns: []string{friend.RequesterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RequesterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.RequesterTable,
			Columns: []string{friend.RequesterColumn},
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
	if fu.mutation.RequesteeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.RequesteeTable,
			Columns: []string{friend.RequesteeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RequesteeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.RequesteeTable,
			Columns: []string{friend.RequesteeColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FriendUpdateOne is the builder for updating a single Friend entity.
type FriendUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FriendMutation
}

// SetCreatedAt sets the "created_at" field.
func (fuo *FriendUpdateOne) SetCreatedAt(t time.Time) *FriendUpdateOne {
	fuo.mutation.SetCreatedAt(t)
	return fuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableCreatedAt(t *time.Time) *FriendUpdateOne {
	if t != nil {
		fuo.SetCreatedAt(*t)
	}
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FriendUpdateOne) SetUpdatedAt(t time.Time) *FriendUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetRequesterID sets the "requester_id" field.
func (fuo *FriendUpdateOne) SetRequesterID(u uuid.UUID) *FriendUpdateOne {
	fuo.mutation.SetRequesterID(u)
	return fuo
}

// SetNillableRequesterID sets the "requester_id" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableRequesterID(u *uuid.UUID) *FriendUpdateOne {
	if u != nil {
		fuo.SetRequesterID(*u)
	}
	return fuo
}

// SetRequesteeID sets the "requestee_id" field.
func (fuo *FriendUpdateOne) SetRequesteeID(u uuid.UUID) *FriendUpdateOne {
	fuo.mutation.SetRequesteeID(u)
	return fuo
}

// SetNillableRequesteeID sets the "requestee_id" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableRequesteeID(u *uuid.UUID) *FriendUpdateOne {
	if u != nil {
		fuo.SetRequesteeID(*u)
	}
	return fuo
}

// SetStatus sets the "status" field.
func (fuo *FriendUpdateOne) SetStatus(f friend.Status) *FriendUpdateOne {
	fuo.mutation.SetStatus(f)
	return fuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableStatus(f *friend.Status) *FriendUpdateOne {
	if f != nil {
		fuo.SetStatus(*f)
	}
	return fuo
}

// SetRequester sets the "requester" edge to the User entity.
func (fuo *FriendUpdateOne) SetRequester(u *User) *FriendUpdateOne {
	return fuo.SetRequesterID(u.ID)
}

// SetRequestee sets the "requestee" edge to the User entity.
func (fuo *FriendUpdateOne) SetRequestee(u *User) *FriendUpdateOne {
	return fuo.SetRequesteeID(u.ID)
}

// Mutation returns the FriendMutation object of the builder.
func (fuo *FriendUpdateOne) Mutation() *FriendMutation {
	return fuo.mutation
}

// ClearRequester clears the "requester" edge to the User entity.
func (fuo *FriendUpdateOne) ClearRequester() *FriendUpdateOne {
	fuo.mutation.ClearRequester()
	return fuo
}

// ClearRequestee clears the "requestee" edge to the User entity.
func (fuo *FriendUpdateOne) ClearRequestee() *FriendUpdateOne {
	fuo.mutation.ClearRequestee()
	return fuo
}

// Where appends a list predicates to the FriendUpdate builder.
func (fuo *FriendUpdateOne) Where(ps ...predicate.Friend) *FriendUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FriendUpdateOne) Select(field string, fields ...string) *FriendUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Friend entity.
func (fuo *FriendUpdateOne) Save(ctx context.Context) (*Friend, error) {
	fuo.defaults()
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FriendUpdateOne) SaveX(ctx context.Context) *Friend {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FriendUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FriendUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FriendUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		v := friend.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FriendUpdateOne) check() error {
	if v, ok := fuo.mutation.Status(); ok {
		if err := friend.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Friend.status": %w`, err)}
		}
	}
	if _, ok := fuo.mutation.RequesterID(); fuo.mutation.RequesterCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friend.requester"`)
	}
	if _, ok := fuo.mutation.RequesteeID(); fuo.mutation.RequesteeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friend.requestee"`)
	}
	return nil
}

func (fuo *FriendUpdateOne) sqlSave(ctx context.Context) (_node *Friend, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(friend.Table, friend.Columns, sqlgraph.NewFieldSpec(friend.FieldID, field.TypeUUID))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Friend.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friend.FieldID)
		for _, f := range fields {
			if !friend.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != friend.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.CreatedAt(); ok {
		_spec.SetField(friend.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(friend.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fuo.mutation.Status(); ok {
		_spec.SetField(friend.FieldStatus, field.TypeEnum, value)
	}
	if fuo.mutation.RequesterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.RequesterTable,
			Columns: []string{friend.RequesterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RequesterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.RequesterTable,
			Columns: []string{friend.RequesterColumn},
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
	if fuo.mutation.RequesteeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.RequesteeTable,
			Columns: []string{friend.RequesteeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RequesteeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.RequesteeTable,
			Columns: []string{friend.RequesteeColumn},
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
	_node = &Friend{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
