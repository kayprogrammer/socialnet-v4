// Code generated by ent, DO NOT EDIT.

package notification

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldUpdatedAt, v))
}

// SenderID applies equality check predicate on the "sender_id" field. It's identical to SenderIDEQ.
func SenderID(v uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldSenderID, v))
}

// PostID applies equality check predicate on the "post_id" field. It's identical to PostIDEQ.
func PostID(v uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldPostID, v))
}

// CommentID applies equality check predicate on the "comment_id" field. It's identical to CommentIDEQ.
func CommentID(v uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldCommentID, v))
}

// ReplyID applies equality check predicate on the "reply_id" field. It's identical to ReplyIDEQ.
func ReplyID(v uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldReplyID, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldText, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldUpdatedAt, v))
}

// SenderIDEQ applies the EQ predicate on the "sender_id" field.
func SenderIDEQ(v uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldSenderID, v))
}

// SenderIDNEQ applies the NEQ predicate on the "sender_id" field.
func SenderIDNEQ(v uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldSenderID, v))
}

// SenderIDIn applies the In predicate on the "sender_id" field.
func SenderIDIn(vs ...uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldSenderID, vs...))
}

// SenderIDNotIn applies the NotIn predicate on the "sender_id" field.
func SenderIDNotIn(vs ...uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldSenderID, vs...))
}

// NtypeEQ applies the EQ predicate on the "ntype" field.
func NtypeEQ(v Ntype) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldNtype, v))
}

// NtypeNEQ applies the NEQ predicate on the "ntype" field.
func NtypeNEQ(v Ntype) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldNtype, v))
}

// NtypeIn applies the In predicate on the "ntype" field.
func NtypeIn(vs ...Ntype) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldNtype, vs...))
}

// NtypeNotIn applies the NotIn predicate on the "ntype" field.
func NtypeNotIn(vs ...Ntype) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldNtype, vs...))
}

// PostIDEQ applies the EQ predicate on the "post_id" field.
func PostIDEQ(v uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldPostID, v))
}

// PostIDNEQ applies the NEQ predicate on the "post_id" field.
func PostIDNEQ(v uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldPostID, v))
}

// PostIDIn applies the In predicate on the "post_id" field.
func PostIDIn(vs ...uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldPostID, vs...))
}

// PostIDNotIn applies the NotIn predicate on the "post_id" field.
func PostIDNotIn(vs ...uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldPostID, vs...))
}

// PostIDIsNil applies the IsNil predicate on the "post_id" field.
func PostIDIsNil() predicate.Notification {
	return predicate.Notification(sql.FieldIsNull(FieldPostID))
}

// PostIDNotNil applies the NotNil predicate on the "post_id" field.
func PostIDNotNil() predicate.Notification {
	return predicate.Notification(sql.FieldNotNull(FieldPostID))
}

// CommentIDEQ applies the EQ predicate on the "comment_id" field.
func CommentIDEQ(v uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldCommentID, v))
}

// CommentIDNEQ applies the NEQ predicate on the "comment_id" field.
func CommentIDNEQ(v uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldCommentID, v))
}

// CommentIDIn applies the In predicate on the "comment_id" field.
func CommentIDIn(vs ...uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldCommentID, vs...))
}

// CommentIDNotIn applies the NotIn predicate on the "comment_id" field.
func CommentIDNotIn(vs ...uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldCommentID, vs...))
}

// CommentIDIsNil applies the IsNil predicate on the "comment_id" field.
func CommentIDIsNil() predicate.Notification {
	return predicate.Notification(sql.FieldIsNull(FieldCommentID))
}

// CommentIDNotNil applies the NotNil predicate on the "comment_id" field.
func CommentIDNotNil() predicate.Notification {
	return predicate.Notification(sql.FieldNotNull(FieldCommentID))
}

// ReplyIDEQ applies the EQ predicate on the "reply_id" field.
func ReplyIDEQ(v uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldReplyID, v))
}

// ReplyIDNEQ applies the NEQ predicate on the "reply_id" field.
func ReplyIDNEQ(v uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldReplyID, v))
}

// ReplyIDIn applies the In predicate on the "reply_id" field.
func ReplyIDIn(vs ...uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldReplyID, vs...))
}

// ReplyIDNotIn applies the NotIn predicate on the "reply_id" field.
func ReplyIDNotIn(vs ...uuid.UUID) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldReplyID, vs...))
}

// ReplyIDIsNil applies the IsNil predicate on the "reply_id" field.
func ReplyIDIsNil() predicate.Notification {
	return predicate.Notification(sql.FieldIsNull(FieldReplyID))
}

// ReplyIDNotNil applies the NotNil predicate on the "reply_id" field.
func ReplyIDNotNil() predicate.Notification {
	return predicate.Notification(sql.FieldNotNull(FieldReplyID))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldText, v))
}

// TextIsNil applies the IsNil predicate on the "text" field.
func TextIsNil() predicate.Notification {
	return predicate.Notification(sql.FieldIsNull(FieldText))
}

// TextNotNil applies the NotNil predicate on the "text" field.
func TextNotNil() predicate.Notification {
	return predicate.Notification(sql.FieldNotNull(FieldText))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldText, v))
}

// HasSender applies the HasEdge predicate on the "sender" edge.
func HasSender() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SenderTable, SenderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSenderWith applies the HasEdge predicate on the "sender" edge with a given conditions (other predicates).
func HasSenderWith(preds ...predicate.User) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := newSenderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReceivers applies the HasEdge predicate on the "receivers" edge.
func HasReceivers() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ReceiversTable, ReceiversPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReceiversWith applies the HasEdge predicate on the "receivers" edge with a given conditions (other predicates).
func HasReceiversWith(preds ...predicate.User) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := newReceiversStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPost applies the HasEdge predicate on the "post" edge.
func HasPost() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PostTable, PostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostWith applies the HasEdge predicate on the "post" edge with a given conditions (other predicates).
func HasPostWith(preds ...predicate.Post) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := newPostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComment applies the HasEdge predicate on the "comment" edge.
func HasComment() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CommentTable, CommentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentWith applies the HasEdge predicate on the "comment" edge with a given conditions (other predicates).
func HasCommentWith(preds ...predicate.Comment) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := newCommentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReply applies the HasEdge predicate on the "reply" edge.
func HasReply() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReplyTable, ReplyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReplyWith applies the HasEdge predicate on the "reply" edge with a given conditions (other predicates).
func HasReplyWith(preds ...predicate.Reply) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := newReplyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReadBy applies the HasEdge predicate on the "read_by" edge.
func HasReadBy() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ReadByTable, ReadByPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReadByWith applies the HasEdge predicate on the "read_by" edge with a given conditions (other predicates).
func HasReadByWith(preds ...predicate.User) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := newReadByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Notification) predicate.Notification {
	return predicate.Notification(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Notification) predicate.Notification {
	return predicate.Notification(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Notification) predicate.Notification {
	return predicate.Notification(sql.NotPredicates(p))
}