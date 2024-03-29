// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent/city"
	"github.com/kayprogrammer/socialnet-v4/ent/file"
	"github.com/kayprogrammer/socialnet-v4/ent/otp"
	"github.com/kayprogrammer/socialnet-v4/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// TermsAgreement holds the value of the "terms_agreement" field.
	TermsAgreement bool `json:"terms_agreement,omitempty"`
	// IsEmailVerified holds the value of the "is_email_verified" field.
	IsEmailVerified bool `json:"is_email_verified,omitempty"`
	// IsStaff holds the value of the "is_staff" field.
	IsStaff bool `json:"is_staff,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// Bio holds the value of the "bio" field.
	Bio *string `json:"bio,omitempty"`
	// Dob holds the value of the "dob" field.
	Dob *time.Time `json:"dob,omitempty"`
	// Access holds the value of the "access" field.
	Access *string `json:"access,omitempty"`
	// Refresh holds the value of the "refresh" field.
	Refresh *string `json:"refresh,omitempty"`
	// CityID holds the value of the "city_id" field.
	CityID *uuid.UUID `json:"city_id,omitempty"`
	// AvatarID holds the value of the "avatar_id" field.
	AvatarID *uuid.UUID `json:"avatar_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// City holds the value of the city edge.
	City *City `json:"city,omitempty"`
	// Avatar holds the value of the avatar edge.
	Avatar *File `json:"avatar,omitempty"`
	// Otp holds the value of the otp edge.
	Otp *Otp `json:"otp,omitempty"`
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// Reactions holds the value of the reactions edge.
	Reactions []*Reaction `json:"reactions,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// Replies holds the value of the replies edge.
	Replies []*Reply `json:"replies,omitempty"`
	// RequesterFriends holds the value of the requester_friends edge.
	RequesterFriends []*Friend `json:"requester_friends,omitempty"`
	// RequesteeFriends holds the value of the requestee_friends edge.
	RequesteeFriends []*Friend `json:"requestee_friends,omitempty"`
	// NotificationsFrom holds the value of the notifications_from edge.
	NotificationsFrom []*Notification `json:"notifications_from,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// NotificationsRead holds the value of the notifications_read edge.
	NotificationsRead []*Notification `json:"notifications_read,omitempty"`
	// OwnedChats holds the value of the owned_chats edge.
	OwnedChats []*Chat `json:"owned_chats,omitempty"`
	// MemberChats holds the value of the member_chats edge.
	MemberChats []*Chat `json:"member_chats,omitempty"`
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [15]bool
}

// CityOrErr returns the City value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CityOrErr() (*City, error) {
	if e.loadedTypes[0] {
		if e.City == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: city.Label}
		}
		return e.City, nil
	}
	return nil, &NotLoadedError{edge: "city"}
}

// AvatarOrErr returns the Avatar value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) AvatarOrErr() (*File, error) {
	if e.loadedTypes[1] {
		if e.Avatar == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: file.Label}
		}
		return e.Avatar, nil
	}
	return nil, &NotLoadedError{edge: "avatar"}
}

// OtpOrErr returns the Otp value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) OtpOrErr() (*Otp, error) {
	if e.loadedTypes[2] {
		if e.Otp == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: otp.Label}
		}
		return e.Otp, nil
	}
	return nil, &NotLoadedError{edge: "otp"}
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[3] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// ReactionsOrErr returns the Reactions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReactionsOrErr() ([]*Reaction, error) {
	if e.loadedTypes[4] {
		return e.Reactions, nil
	}
	return nil, &NotLoadedError{edge: "reactions"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[5] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// RepliesOrErr returns the Replies value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RepliesOrErr() ([]*Reply, error) {
	if e.loadedTypes[6] {
		return e.Replies, nil
	}
	return nil, &NotLoadedError{edge: "replies"}
}

// RequesterFriendsOrErr returns the RequesterFriends value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RequesterFriendsOrErr() ([]*Friend, error) {
	if e.loadedTypes[7] {
		return e.RequesterFriends, nil
	}
	return nil, &NotLoadedError{edge: "requester_friends"}
}

// RequesteeFriendsOrErr returns the RequesteeFriends value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RequesteeFriendsOrErr() ([]*Friend, error) {
	if e.loadedTypes[8] {
		return e.RequesteeFriends, nil
	}
	return nil, &NotLoadedError{edge: "requestee_friends"}
}

// NotificationsFromOrErr returns the NotificationsFrom value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) NotificationsFromOrErr() ([]*Notification, error) {
	if e.loadedTypes[9] {
		return e.NotificationsFrom, nil
	}
	return nil, &NotLoadedError{edge: "notifications_from"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[10] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// NotificationsReadOrErr returns the NotificationsRead value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) NotificationsReadOrErr() ([]*Notification, error) {
	if e.loadedTypes[11] {
		return e.NotificationsRead, nil
	}
	return nil, &NotLoadedError{edge: "notifications_read"}
}

// OwnedChatsOrErr returns the OwnedChats value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OwnedChatsOrErr() ([]*Chat, error) {
	if e.loadedTypes[12] {
		return e.OwnedChats, nil
	}
	return nil, &NotLoadedError{edge: "owned_chats"}
}

// MemberChatsOrErr returns the MemberChats value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MemberChatsOrErr() ([]*Chat, error) {
	if e.loadedTypes[13] {
		return e.MemberChats, nil
	}
	return nil, &NotLoadedError{edge: "member_chats"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[14] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldCityID, user.FieldAvatarID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case user.FieldTermsAgreement, user.FieldIsEmailVerified, user.FieldIsStaff, user.FieldIsActive:
			values[i] = new(sql.NullBool)
		case user.FieldFirstName, user.FieldLastName, user.FieldUsername, user.FieldEmail, user.FieldPassword, user.FieldBio, user.FieldAccess, user.FieldRefresh:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDob:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldTermsAgreement:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field terms_agreement", values[i])
			} else if value.Valid {
				u.TermsAgreement = value.Bool
			}
		case user.FieldIsEmailVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_email_verified", values[i])
			} else if value.Valid {
				u.IsEmailVerified = value.Bool
			}
		case user.FieldIsStaff:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_staff", values[i])
			} else if value.Valid {
				u.IsStaff = value.Bool
			}
		case user.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				u.IsActive = value.Bool
			}
		case user.FieldBio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bio", values[i])
			} else if value.Valid {
				u.Bio = new(string)
				*u.Bio = value.String
			}
		case user.FieldDob:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dob", values[i])
			} else if value.Valid {
				u.Dob = new(time.Time)
				*u.Dob = value.Time
			}
		case user.FieldAccess:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access", values[i])
			} else if value.Valid {
				u.Access = new(string)
				*u.Access = value.String
			}
		case user.FieldRefresh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refresh", values[i])
			} else if value.Valid {
				u.Refresh = new(string)
				*u.Refresh = value.String
			}
		case user.FieldCityID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field city_id", values[i])
			} else if value.Valid {
				u.CityID = new(uuid.UUID)
				*u.CityID = *value.S.(*uuid.UUID)
			}
		case user.FieldAvatarID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_id", values[i])
			} else if value.Valid {
				u.AvatarID = new(uuid.UUID)
				*u.AvatarID = *value.S.(*uuid.UUID)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryCity queries the "city" edge of the User entity.
func (u *User) QueryCity() *CityQuery {
	return NewUserClient(u.config).QueryCity(u)
}

// QueryAvatar queries the "avatar" edge of the User entity.
func (u *User) QueryAvatar() *FileQuery {
	return NewUserClient(u.config).QueryAvatar(u)
}

// QueryOtp queries the "otp" edge of the User entity.
func (u *User) QueryOtp() *OtpQuery {
	return NewUserClient(u.config).QueryOtp(u)
}

// QueryPosts queries the "posts" edge of the User entity.
func (u *User) QueryPosts() *PostQuery {
	return NewUserClient(u.config).QueryPosts(u)
}

// QueryReactions queries the "reactions" edge of the User entity.
func (u *User) QueryReactions() *ReactionQuery {
	return NewUserClient(u.config).QueryReactions(u)
}

// QueryComments queries the "comments" edge of the User entity.
func (u *User) QueryComments() *CommentQuery {
	return NewUserClient(u.config).QueryComments(u)
}

// QueryReplies queries the "replies" edge of the User entity.
func (u *User) QueryReplies() *ReplyQuery {
	return NewUserClient(u.config).QueryReplies(u)
}

// QueryRequesterFriends queries the "requester_friends" edge of the User entity.
func (u *User) QueryRequesterFriends() *FriendQuery {
	return NewUserClient(u.config).QueryRequesterFriends(u)
}

// QueryRequesteeFriends queries the "requestee_friends" edge of the User entity.
func (u *User) QueryRequesteeFriends() *FriendQuery {
	return NewUserClient(u.config).QueryRequesteeFriends(u)
}

// QueryNotificationsFrom queries the "notifications_from" edge of the User entity.
func (u *User) QueryNotificationsFrom() *NotificationQuery {
	return NewUserClient(u.config).QueryNotificationsFrom(u)
}

// QueryNotifications queries the "notifications" edge of the User entity.
func (u *User) QueryNotifications() *NotificationQuery {
	return NewUserClient(u.config).QueryNotifications(u)
}

// QueryNotificationsRead queries the "notifications_read" edge of the User entity.
func (u *User) QueryNotificationsRead() *NotificationQuery {
	return NewUserClient(u.config).QueryNotificationsRead(u)
}

// QueryOwnedChats queries the "owned_chats" edge of the User entity.
func (u *User) QueryOwnedChats() *ChatQuery {
	return NewUserClient(u.config).QueryOwnedChats(u)
}

// QueryMemberChats queries the "member_chats" edge of the User entity.
func (u *User) QueryMemberChats() *ChatQuery {
	return NewUserClient(u.config).QueryMemberChats(u)
}

// QueryMessages queries the "messages" edge of the User entity.
func (u *User) QueryMessages() *MessageQuery {
	return NewUserClient(u.config).QueryMessages(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("terms_agreement=")
	builder.WriteString(fmt.Sprintf("%v", u.TermsAgreement))
	builder.WriteString(", ")
	builder.WriteString("is_email_verified=")
	builder.WriteString(fmt.Sprintf("%v", u.IsEmailVerified))
	builder.WriteString(", ")
	builder.WriteString("is_staff=")
	builder.WriteString(fmt.Sprintf("%v", u.IsStaff))
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", u.IsActive))
	builder.WriteString(", ")
	if v := u.Bio; v != nil {
		builder.WriteString("bio=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Dob; v != nil {
		builder.WriteString("dob=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.Access; v != nil {
		builder.WriteString("access=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Refresh; v != nil {
		builder.WriteString("refresh=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.CityID; v != nil {
		builder.WriteString("city_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.AvatarID; v != nil {
		builder.WriteString("avatar_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
