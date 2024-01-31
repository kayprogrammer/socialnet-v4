// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent/chat"
	"github.com/kayprogrammer/socialnet-v4/ent/file"
	"github.com/kayprogrammer/socialnet-v4/ent/message"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Text holds the value of the "text" field.
	Text *string `json:"text,omitempty"`
	// SenderID holds the value of the "sender_id" field.
	SenderID uuid.UUID `json:"sender_id,omitempty"`
	// ChatID holds the value of the "chat_id" field.
	ChatID uuid.UUID `json:"chat_id,omitempty"`
	// FileID holds the value of the "file_id" field.
	FileID *uuid.UUID `json:"file_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessageQuery when eager-loading is set.
	Edges        MessageEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MessageEdges holds the relations/edges for other nodes in the graph.
type MessageEdges struct {
	// Sender holds the value of the sender edge.
	Sender *User `json:"sender,omitempty"`
	// Chat holds the value of the chat edge.
	Chat *Chat `json:"chat,omitempty"`
	// File holds the value of the file edge.
	File *File `json:"file,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SenderOrErr returns the Sender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) SenderOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Sender == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Sender, nil
	}
	return nil, &NotLoadedError{edge: "sender"}
}

// ChatOrErr returns the Chat value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) ChatOrErr() (*Chat, error) {
	if e.loadedTypes[1] {
		if e.Chat == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: chat.Label}
		}
		return e.Chat, nil
	}
	return nil, &NotLoadedError{edge: "chat"}
}

// FileOrErr returns the File value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) FileOrErr() (*File, error) {
	if e.loadedTypes[2] {
		if e.File == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: file.Label}
		}
		return e.File, nil
	}
	return nil, &NotLoadedError{edge: "file"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldFileID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case message.FieldText:
			values[i] = new(sql.NullString)
		case message.FieldCreatedAt, message.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case message.FieldID, message.FieldSenderID, message.FieldChatID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case message.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case message.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case message.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				m.Text = new(string)
				*m.Text = value.String
			}
		case message.FieldSenderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field sender_id", values[i])
			} else if value != nil {
				m.SenderID = *value
			}
		case message.FieldChatID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field chat_id", values[i])
			} else if value != nil {
				m.ChatID = *value
			}
		case message.FieldFileID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field file_id", values[i])
			} else if value.Valid {
				m.FileID = new(uuid.UUID)
				*m.FileID = *value.S.(*uuid.UUID)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Message.
// This includes values selected through modifiers, order, etc.
func (m *Message) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QuerySender queries the "sender" edge of the Message entity.
func (m *Message) QuerySender() *UserQuery {
	return NewMessageClient(m.config).QuerySender(m)
}

// QueryChat queries the "chat" edge of the Message entity.
func (m *Message) QueryChat() *ChatQuery {
	return NewMessageClient(m.config).QueryChat(m)
}

// QueryFile queries the "file" edge of the Message entity.
func (m *Message) QueryFile() *FileQuery {
	return NewMessageClient(m.config).QueryFile(m)
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return NewMessageClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := m.Text; v != nil {
		builder.WriteString("text=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("sender_id=")
	builder.WriteString(fmt.Sprintf("%v", m.SenderID))
	builder.WriteString(", ")
	builder.WriteString("chat_id=")
	builder.WriteString(fmt.Sprintf("%v", m.ChatID))
	builder.WriteString(", ")
	if v := m.FileID; v != nil {
		builder.WriteString("file_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message
