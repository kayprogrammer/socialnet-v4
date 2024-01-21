// Code generated by ent, DO NOT EDIT.

package reply

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the reply type in the database.
	Label = "reply"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldAuthorID holds the string denoting the author_id field in the database.
	FieldAuthorID = "author_id"
	// FieldCommentID holds the string denoting the comment_id field in the database.
	FieldCommentID = "comment_id"
	// EdgeReactions holds the string denoting the reactions edge name in mutations.
	EdgeReactions = "reactions"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeComment holds the string denoting the comment edge name in mutations.
	EdgeComment = "comment"
	// EdgeNotifications holds the string denoting the notifications edge name in mutations.
	EdgeNotifications = "notifications"
	// Table holds the table name of the reply in the database.
	Table = "replies"
	// ReactionsTable is the table that holds the reactions relation/edge.
	ReactionsTable = "reactions"
	// ReactionsInverseTable is the table name for the Reaction entity.
	// It exists in this package in order to avoid circular dependency with the "reaction" package.
	ReactionsInverseTable = "reactions"
	// ReactionsColumn is the table column denoting the reactions relation/edge.
	ReactionsColumn = "reply_id"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "replies"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "author_id"
	// CommentTable is the table that holds the comment relation/edge.
	CommentTable = "replies"
	// CommentInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentInverseTable = "comments"
	// CommentColumn is the table column denoting the comment relation/edge.
	CommentColumn = "comment_id"
	// NotificationsTable is the table that holds the notifications relation/edge.
	NotificationsTable = "notifications"
	// NotificationsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsInverseTable = "notifications"
	// NotificationsColumn is the table column denoting the notifications relation/edge.
	NotificationsColumn = "reply_id"
)

// Columns holds all SQL columns for reply fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldText,
	FieldSlug,
	FieldAuthorID,
	FieldCommentID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// TextValidator is a validator for the "text" field. It is called by the builders before save.
	TextValidator func(string) error
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Reply queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByAuthorID orders the results by the author_id field.
func ByAuthorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthorID, opts...).ToFunc()
}

// ByCommentID orders the results by the comment_id field.
func ByCommentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommentID, opts...).ToFunc()
}

// ByReactionsCount orders the results by reactions count.
func ByReactionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReactionsStep(), opts...)
	}
}

// ByReactions orders the results by reactions terms.
func ByReactions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReactionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAuthorField orders the results by author field.
func ByAuthorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), sql.OrderByField(field, opts...))
	}
}

// ByCommentField orders the results by comment field.
func ByCommentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentStep(), sql.OrderByField(field, opts...))
	}
}

// ByNotificationsCount orders the results by notifications count.
func ByNotificationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationsStep(), opts...)
	}
}

// ByNotifications orders the results by notifications terms.
func ByNotifications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newReactionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReactionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReactionsTable, ReactionsColumn),
	)
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
	)
}
func newCommentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CommentTable, CommentColumn),
	)
}
func newNotificationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotificationsTable, NotificationsColumn),
	)
}
