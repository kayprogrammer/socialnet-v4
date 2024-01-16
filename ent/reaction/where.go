// Code generated by ent, DO NOT EDIT.

package reaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldUserID, v))
}

// Rtype applies equality check predicate on the "rtype" field. It's identical to RtypeEQ.
func Rtype(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldRtype, v))
}

// PostID applies equality check predicate on the "post_id" field. It's identical to PostIDEQ.
func PostID(v uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldPostID, v))
}

// CommentID applies equality check predicate on the "comment_id" field. It's identical to CommentIDEQ.
func CommentID(v uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldCommentID, v))
}

// ReplyID applies equality check predicate on the "reply_id" field. It's identical to ReplyIDEQ.
func ReplyID(v uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldReplyID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Reaction {
	return predicate.Reaction(sql.FieldLTE(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldNotIn(FieldUserID, vs...))
}

// RtypeEQ applies the EQ predicate on the "rtype" field.
func RtypeEQ(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldRtype, v))
}

// RtypeNEQ applies the NEQ predicate on the "rtype" field.
func RtypeNEQ(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldNEQ(FieldRtype, v))
}

// RtypeIn applies the In predicate on the "rtype" field.
func RtypeIn(vs ...string) predicate.Reaction {
	return predicate.Reaction(sql.FieldIn(FieldRtype, vs...))
}

// RtypeNotIn applies the NotIn predicate on the "rtype" field.
func RtypeNotIn(vs ...string) predicate.Reaction {
	return predicate.Reaction(sql.FieldNotIn(FieldRtype, vs...))
}

// RtypeGT applies the GT predicate on the "rtype" field.
func RtypeGT(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldGT(FieldRtype, v))
}

// RtypeGTE applies the GTE predicate on the "rtype" field.
func RtypeGTE(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldGTE(FieldRtype, v))
}

// RtypeLT applies the LT predicate on the "rtype" field.
func RtypeLT(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldLT(FieldRtype, v))
}

// RtypeLTE applies the LTE predicate on the "rtype" field.
func RtypeLTE(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldLTE(FieldRtype, v))
}

// RtypeContains applies the Contains predicate on the "rtype" field.
func RtypeContains(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldContains(FieldRtype, v))
}

// RtypeHasPrefix applies the HasPrefix predicate on the "rtype" field.
func RtypeHasPrefix(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldHasPrefix(FieldRtype, v))
}

// RtypeHasSuffix applies the HasSuffix predicate on the "rtype" field.
func RtypeHasSuffix(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldHasSuffix(FieldRtype, v))
}

// RtypeEqualFold applies the EqualFold predicate on the "rtype" field.
func RtypeEqualFold(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldEqualFold(FieldRtype, v))
}

// RtypeContainsFold applies the ContainsFold predicate on the "rtype" field.
func RtypeContainsFold(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldContainsFold(FieldRtype, v))
}

// PostIDEQ applies the EQ predicate on the "post_id" field.
func PostIDEQ(v uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldPostID, v))
}

// PostIDNEQ applies the NEQ predicate on the "post_id" field.
func PostIDNEQ(v uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldNEQ(FieldPostID, v))
}

// PostIDIn applies the In predicate on the "post_id" field.
func PostIDIn(vs ...uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldIn(FieldPostID, vs...))
}

// PostIDNotIn applies the NotIn predicate on the "post_id" field.
func PostIDNotIn(vs ...uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldNotIn(FieldPostID, vs...))
}

// PostIDIsNil applies the IsNil predicate on the "post_id" field.
func PostIDIsNil() predicate.Reaction {
	return predicate.Reaction(sql.FieldIsNull(FieldPostID))
}

// PostIDNotNil applies the NotNil predicate on the "post_id" field.
func PostIDNotNil() predicate.Reaction {
	return predicate.Reaction(sql.FieldNotNull(FieldPostID))
}

// CommentIDEQ applies the EQ predicate on the "comment_id" field.
func CommentIDEQ(v uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldCommentID, v))
}

// CommentIDNEQ applies the NEQ predicate on the "comment_id" field.
func CommentIDNEQ(v uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldNEQ(FieldCommentID, v))
}

// CommentIDIn applies the In predicate on the "comment_id" field.
func CommentIDIn(vs ...uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldIn(FieldCommentID, vs...))
}

// CommentIDNotIn applies the NotIn predicate on the "comment_id" field.
func CommentIDNotIn(vs ...uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldNotIn(FieldCommentID, vs...))
}

// CommentIDIsNil applies the IsNil predicate on the "comment_id" field.
func CommentIDIsNil() predicate.Reaction {
	return predicate.Reaction(sql.FieldIsNull(FieldCommentID))
}

// CommentIDNotNil applies the NotNil predicate on the "comment_id" field.
func CommentIDNotNil() predicate.Reaction {
	return predicate.Reaction(sql.FieldNotNull(FieldCommentID))
}

// ReplyIDEQ applies the EQ predicate on the "reply_id" field.
func ReplyIDEQ(v uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldReplyID, v))
}

// ReplyIDNEQ applies the NEQ predicate on the "reply_id" field.
func ReplyIDNEQ(v uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldNEQ(FieldReplyID, v))
}

// ReplyIDIn applies the In predicate on the "reply_id" field.
func ReplyIDIn(vs ...uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldIn(FieldReplyID, vs...))
}

// ReplyIDNotIn applies the NotIn predicate on the "reply_id" field.
func ReplyIDNotIn(vs ...uuid.UUID) predicate.Reaction {
	return predicate.Reaction(sql.FieldNotIn(FieldReplyID, vs...))
}

// ReplyIDIsNil applies the IsNil predicate on the "reply_id" field.
func ReplyIDIsNil() predicate.Reaction {
	return predicate.Reaction(sql.FieldIsNull(FieldReplyID))
}

// ReplyIDNotNil applies the NotNil predicate on the "reply_id" field.
func ReplyIDNotNil() predicate.Reaction {
	return predicate.Reaction(sql.FieldNotNull(FieldReplyID))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPost applies the HasEdge predicate on the "post" edge.
func HasPost() predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PostTable, PostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostWith applies the HasEdge predicate on the "post" edge with a given conditions (other predicates).
func HasPostWith(preds ...predicate.Post) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := newPostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComment applies the HasEdge predicate on the "comment" edge.
func HasComment() predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CommentTable, CommentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentWith applies the HasEdge predicate on the "comment" edge with a given conditions (other predicates).
func HasCommentWith(preds ...predicate.Comment) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := newCommentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReply applies the HasEdge predicate on the "reply" edge.
func HasReply() predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReplyTable, ReplyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReplyWith applies the HasEdge predicate on the "reply" edge with a given conditions (other predicates).
func HasReplyWith(preds ...predicate.Reply) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := newReplyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Reaction) predicate.Reaction {
	return predicate.Reaction(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Reaction) predicate.Reaction {
	return predicate.Reaction(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Reaction) predicate.Reaction {
	return predicate.Reaction(sql.NotPredicates(p))
}