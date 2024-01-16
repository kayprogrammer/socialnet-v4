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
	"github.com/kayprogrammer/socialnet-v4/ent/file"
	"github.com/kayprogrammer/socialnet-v4/ent/post"
	"github.com/kayprogrammer/socialnet-v4/ent/predicate"
	"github.com/kayprogrammer/socialnet-v4/ent/reaction"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PostUpdate) SetCreatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PostUpdate) SetNillableCreatedAt(t *time.Time) *PostUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PostUpdate) SetUpdatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetText sets the "text" field.
func (pu *PostUpdate) SetText(s string) *PostUpdate {
	pu.mutation.SetText(s)
	return pu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (pu *PostUpdate) SetNillableText(s *string) *PostUpdate {
	if s != nil {
		pu.SetText(*s)
	}
	return pu
}

// SetSlug sets the "slug" field.
func (pu *PostUpdate) SetSlug(s string) *PostUpdate {
	pu.mutation.SetSlug(s)
	return pu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (pu *PostUpdate) SetNillableSlug(s *string) *PostUpdate {
	if s != nil {
		pu.SetSlug(*s)
	}
	return pu
}

// SetAuthorID sets the "author_id" field.
func (pu *PostUpdate) SetAuthorID(u uuid.UUID) *PostUpdate {
	pu.mutation.SetAuthorID(u)
	return pu
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillableAuthorID(u *uuid.UUID) *PostUpdate {
	if u != nil {
		pu.SetAuthorID(*u)
	}
	return pu
}

// SetImageID sets the "image_id" field.
func (pu *PostUpdate) SetImageID(u uuid.UUID) *PostUpdate {
	pu.mutation.SetImageID(u)
	return pu
}

// SetNillableImageID sets the "image_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillableImageID(u *uuid.UUID) *PostUpdate {
	if u != nil {
		pu.SetImageID(*u)
	}
	return pu
}

// ClearImageID clears the value of the "image_id" field.
func (pu *PostUpdate) ClearImageID() *PostUpdate {
	pu.mutation.ClearImageID()
	return pu
}

// AddReactionIDs adds the "reactions" edge to the Reaction entity by IDs.
func (pu *PostUpdate) AddReactionIDs(ids ...uuid.UUID) *PostUpdate {
	pu.mutation.AddReactionIDs(ids...)
	return pu
}

// AddReactions adds the "reactions" edges to the Reaction entity.
func (pu *PostUpdate) AddReactions(r ...*Reaction) *PostUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddReactionIDs(ids...)
}

// SetAuthor sets the "author" edge to the User entity.
func (pu *PostUpdate) SetAuthor(u *User) *PostUpdate {
	return pu.SetAuthorID(u.ID)
}

// SetImage sets the "image" edge to the File entity.
func (pu *PostUpdate) SetImage(f *File) *PostUpdate {
	return pu.SetImageID(f.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (pu *PostUpdate) AddCommentIDs(ids ...uuid.UUID) *PostUpdate {
	pu.mutation.AddCommentIDs(ids...)
	return pu
}

// AddComments adds the "comments" edges to the Comment entity.
func (pu *PostUpdate) AddComments(c ...*Comment) *PostUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddCommentIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// ClearReactions clears all "reactions" edges to the Reaction entity.
func (pu *PostUpdate) ClearReactions() *PostUpdate {
	pu.mutation.ClearReactions()
	return pu
}

// RemoveReactionIDs removes the "reactions" edge to Reaction entities by IDs.
func (pu *PostUpdate) RemoveReactionIDs(ids ...uuid.UUID) *PostUpdate {
	pu.mutation.RemoveReactionIDs(ids...)
	return pu
}

// RemoveReactions removes "reactions" edges to Reaction entities.
func (pu *PostUpdate) RemoveReactions(r ...*Reaction) *PostUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveReactionIDs(ids...)
}

// ClearAuthor clears the "author" edge to the User entity.
func (pu *PostUpdate) ClearAuthor() *PostUpdate {
	pu.mutation.ClearAuthor()
	return pu
}

// ClearImage clears the "image" edge to the File entity.
func (pu *PostUpdate) ClearImage() *PostUpdate {
	pu.mutation.ClearImage()
	return pu
}

// ClearComments clears all "comments" edges to the Comment entity.
func (pu *PostUpdate) ClearComments() *PostUpdate {
	pu.mutation.ClearComments()
	return pu
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (pu *PostUpdate) RemoveCommentIDs(ids ...uuid.UUID) *PostUpdate {
	pu.mutation.RemoveCommentIDs(ids...)
	return pu
}

// RemoveComments removes "comments" edges to Comment entities.
func (pu *PostUpdate) RemoveComments(c ...*Comment) *PostUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCommentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PostUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := post.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PostUpdate) check() error {
	if v, ok := pu.mutation.Text(); ok {
		if err := post.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Post.text": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Slug(); ok {
		if err := post.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Post.slug": %w`, err)}
		}
	}
	if _, ok := pu.mutation.AuthorID(); pu.mutation.AuthorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Post.author"`)
	}
	return nil
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Text(); ok {
		_spec.SetField(post.FieldText, field.TypeString, value)
	}
	if value, ok := pu.mutation.Slug(); ok {
		_spec.SetField(post.FieldSlug, field.TypeString, value)
	}
	if pu.mutation.ReactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.ReactionsTable,
			Columns: []string{post.ReactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reaction.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedReactionsIDs(); len(nodes) > 0 && !pu.mutation.ReactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.ReactionsTable,
			Columns: []string{post.ReactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ReactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.ReactionsTable,
			Columns: []string{post.ReactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
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
	if pu.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.ImageTable,
			Columns: []string{post.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.ImageTable,
			Columns: []string{post.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !pu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

// SetCreatedAt sets the "created_at" field.
func (puo *PostUpdateOne) SetCreatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableCreatedAt(t *time.Time) *PostUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PostUpdateOne) SetUpdatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetText sets the "text" field.
func (puo *PostUpdateOne) SetText(s string) *PostUpdateOne {
	puo.mutation.SetText(s)
	return puo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableText(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetText(*s)
	}
	return puo
}

// SetSlug sets the "slug" field.
func (puo *PostUpdateOne) SetSlug(s string) *PostUpdateOne {
	puo.mutation.SetSlug(s)
	return puo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableSlug(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetSlug(*s)
	}
	return puo
}

// SetAuthorID sets the "author_id" field.
func (puo *PostUpdateOne) SetAuthorID(u uuid.UUID) *PostUpdateOne {
	puo.mutation.SetAuthorID(u)
	return puo
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableAuthorID(u *uuid.UUID) *PostUpdateOne {
	if u != nil {
		puo.SetAuthorID(*u)
	}
	return puo
}

// SetImageID sets the "image_id" field.
func (puo *PostUpdateOne) SetImageID(u uuid.UUID) *PostUpdateOne {
	puo.mutation.SetImageID(u)
	return puo
}

// SetNillableImageID sets the "image_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableImageID(u *uuid.UUID) *PostUpdateOne {
	if u != nil {
		puo.SetImageID(*u)
	}
	return puo
}

// ClearImageID clears the value of the "image_id" field.
func (puo *PostUpdateOne) ClearImageID() *PostUpdateOne {
	puo.mutation.ClearImageID()
	return puo
}

// AddReactionIDs adds the "reactions" edge to the Reaction entity by IDs.
func (puo *PostUpdateOne) AddReactionIDs(ids ...uuid.UUID) *PostUpdateOne {
	puo.mutation.AddReactionIDs(ids...)
	return puo
}

// AddReactions adds the "reactions" edges to the Reaction entity.
func (puo *PostUpdateOne) AddReactions(r ...*Reaction) *PostUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddReactionIDs(ids...)
}

// SetAuthor sets the "author" edge to the User entity.
func (puo *PostUpdateOne) SetAuthor(u *User) *PostUpdateOne {
	return puo.SetAuthorID(u.ID)
}

// SetImage sets the "image" edge to the File entity.
func (puo *PostUpdateOne) SetImage(f *File) *PostUpdateOne {
	return puo.SetImageID(f.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (puo *PostUpdateOne) AddCommentIDs(ids ...uuid.UUID) *PostUpdateOne {
	puo.mutation.AddCommentIDs(ids...)
	return puo
}

// AddComments adds the "comments" edges to the Comment entity.
func (puo *PostUpdateOne) AddComments(c ...*Comment) *PostUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddCommentIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// ClearReactions clears all "reactions" edges to the Reaction entity.
func (puo *PostUpdateOne) ClearReactions() *PostUpdateOne {
	puo.mutation.ClearReactions()
	return puo
}

// RemoveReactionIDs removes the "reactions" edge to Reaction entities by IDs.
func (puo *PostUpdateOne) RemoveReactionIDs(ids ...uuid.UUID) *PostUpdateOne {
	puo.mutation.RemoveReactionIDs(ids...)
	return puo
}

// RemoveReactions removes "reactions" edges to Reaction entities.
func (puo *PostUpdateOne) RemoveReactions(r ...*Reaction) *PostUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveReactionIDs(ids...)
}

// ClearAuthor clears the "author" edge to the User entity.
func (puo *PostUpdateOne) ClearAuthor() *PostUpdateOne {
	puo.mutation.ClearAuthor()
	return puo
}

// ClearImage clears the "image" edge to the File entity.
func (puo *PostUpdateOne) ClearImage() *PostUpdateOne {
	puo.mutation.ClearImage()
	return puo
}

// ClearComments clears all "comments" edges to the Comment entity.
func (puo *PostUpdateOne) ClearComments() *PostUpdateOne {
	puo.mutation.ClearComments()
	return puo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (puo *PostUpdateOne) RemoveCommentIDs(ids ...uuid.UUID) *PostUpdateOne {
	puo.mutation.RemoveCommentIDs(ids...)
	return puo
}

// RemoveComments removes "comments" edges to Comment entities.
func (puo *PostUpdateOne) RemoveComments(c ...*Comment) *PostUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCommentIDs(ids...)
}

// Where appends a list predicates to the PostUpdate builder.
func (puo *PostUpdateOne) Where(ps ...predicate.Post) *PostUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PostUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := post.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PostUpdateOne) check() error {
	if v, ok := puo.mutation.Text(); ok {
		if err := post.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Post.text": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Slug(); ok {
		if err := post.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Post.slug": %w`, err)}
		}
	}
	if _, ok := puo.mutation.AuthorID(); puo.mutation.AuthorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Post.author"`)
	}
	return nil
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Text(); ok {
		_spec.SetField(post.FieldText, field.TypeString, value)
	}
	if value, ok := puo.mutation.Slug(); ok {
		_spec.SetField(post.FieldSlug, field.TypeString, value)
	}
	if puo.mutation.ReactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.ReactionsTable,
			Columns: []string{post.ReactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reaction.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedReactionsIDs(); len(nodes) > 0 && !puo.mutation.ReactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.ReactionsTable,
			Columns: []string{post.ReactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ReactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.ReactionsTable,
			Columns: []string{post.ReactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
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
	if puo.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.ImageTable,
			Columns: []string{post.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.ImageTable,
			Columns: []string{post.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !puo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
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
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}