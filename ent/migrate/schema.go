// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChatsColumns holds the columns for the "chats" table.
	ChatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "ctype", Type: field.TypeEnum, Enums: []string{"DM", "GROUP"}},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "image_id", Type: field.TypeUUID, Nullable: true},
		{Name: "owner_id", Type: field.TypeUUID},
	}
	// ChatsTable holds the schema information for the "chats" table.
	ChatsTable = &schema.Table{
		Name:       "chats",
		Columns:    ChatsColumns,
		PrimaryKey: []*schema.Column{ChatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "chats_files_chats",
				Columns:    []*schema.Column{ChatsColumns[6]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "chats_users_owned_chats",
				Columns:    []*schema.Column{ChatsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CitiesColumns holds the columns for the "cities" table.
	CitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "country_id", Type: field.TypeUUID},
		{Name: "region_id", Type: field.TypeUUID, Nullable: true},
	}
	// CitiesTable holds the schema information for the "cities" table.
	CitiesTable = &schema.Table{
		Name:       "cities",
		Columns:    CitiesColumns,
		PrimaryKey: []*schema.Column{CitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cities_countries_cities",
				Columns:    []*schema.Column{CitiesColumns[4]},
				RefColumns: []*schema.Column{CountriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "cities_regions_cities",
				Columns:    []*schema.Column{CitiesColumns[5]},
				RefColumns: []*schema.Column{RegionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "post_id", Type: field.TypeUUID},
		{Name: "author_id", Type: field.TypeUUID},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_posts_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CountriesColumns holds the columns for the "countries" table.
	CountriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
	}
	// CountriesTable holds the schema information for the "countries" table.
	CountriesTable = &schema.Table{
		Name:       "countries",
		Columns:    CountriesColumns,
		PrimaryKey: []*schema.Column{CountriesColumns[0]},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "resource_type", Type: field.TypeString},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
	}
	// FriendsColumns holds the columns for the "friends" table.
	FriendsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PENDING", "ACCEPTED"}, Default: "PENDING"},
		{Name: "requester_id", Type: field.TypeUUID},
		{Name: "requestee_id", Type: field.TypeUUID},
	}
	// FriendsTable holds the schema information for the "friends" table.
	FriendsTable = &schema.Table{
		Name:       "friends",
		Columns:    FriendsColumns,
		PrimaryKey: []*schema.Column{FriendsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "friends_users_requester_friends",
				Columns:    []*schema.Column{FriendsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "friends_users_requestee_friends",
				Columns:    []*schema.Column{FriendsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "unique_requester_requestee_combination",
				Unique:  true,
				Columns: []*schema.Column{FriendsColumns[4], FriendsColumns[5]},
			},
			{
				Name:    "unique_requestee_requester_combination",
				Unique:  true,
				Columns: []*schema.Column{FriendsColumns[5], FriendsColumns[4]},
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString, Nullable: true},
		{Name: "chat_id", Type: field.TypeUUID},
		{Name: "file_id", Type: field.TypeUUID, Nullable: true},
		{Name: "sender_id", Type: field.TypeUUID},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_chats_messages",
				Columns:    []*schema.Column{MessagesColumns[4]},
				RefColumns: []*schema.Column{ChatsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "messages_files_messages",
				Columns:    []*schema.Column{MessagesColumns[5]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "messages_users_messages",
				Columns:    []*schema.Column{MessagesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NotificationsColumns holds the columns for the "notifications" table.
	NotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "ntype", Type: field.TypeEnum, Enums: []string{"REACTION", "COMMENT", "REPLY", "ADMIN"}},
		{Name: "text", Type: field.TypeString, Nullable: true},
		{Name: "comment_id", Type: field.TypeUUID, Nullable: true},
		{Name: "post_id", Type: field.TypeUUID, Nullable: true},
		{Name: "reply_id", Type: field.TypeUUID, Nullable: true},
		{Name: "sender_id", Type: field.TypeUUID, Nullable: true},
	}
	// NotificationsTable holds the schema information for the "notifications" table.
	NotificationsTable = &schema.Table{
		Name:       "notifications",
		Columns:    NotificationsColumns,
		PrimaryKey: []*schema.Column{NotificationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notifications_comments_notifications",
				Columns:    []*schema.Column{NotificationsColumns[5]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "notifications_posts_notifications",
				Columns:    []*schema.Column{NotificationsColumns[6]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "notifications_replies_notifications",
				Columns:    []*schema.Column{NotificationsColumns[7]},
				RefColumns: []*schema.Column{RepliesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "notifications_users_notifications_from",
				Columns:    []*schema.Column{NotificationsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// OtpsColumns holds the columns for the "otps" table.
	OtpsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "code", Type: field.TypeUint32},
		{Name: "user_id", Type: field.TypeUUID, Unique: true},
	}
	// OtpsTable holds the schema information for the "otps" table.
	OtpsTable = &schema.Table{
		Name:       "otps",
		Columns:    OtpsColumns,
		PrimaryKey: []*schema.Column{OtpsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "otps_users_otp",
				Columns:    []*schema.Column{OtpsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "image_id", Type: field.TypeUUID, Nullable: true},
		{Name: "author_id", Type: field.TypeUUID},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_files_posts",
				Columns:    []*schema.Column{PostsColumns[5]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ReactionsColumns holds the columns for the "reactions" table.
	ReactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "rtype", Type: field.TypeEnum, Enums: []string{"LIKE", "LOVE", "HAHA", "WOW", "SAD", "ANGRY"}, Default: "LIKE"},
		{Name: "comment_id", Type: field.TypeUUID, Nullable: true},
		{Name: "post_id", Type: field.TypeUUID, Nullable: true},
		{Name: "reply_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// ReactionsTable holds the schema information for the "reactions" table.
	ReactionsTable = &schema.Table{
		Name:       "reactions",
		Columns:    ReactionsColumns,
		PrimaryKey: []*schema.Column{ReactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reactions_comments_reactions",
				Columns:    []*schema.Column{ReactionsColumns[4]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reactions_posts_reactions",
				Columns:    []*schema.Column{ReactionsColumns[5]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reactions_replies_reactions",
				Columns:    []*schema.Column{ReactionsColumns[6]},
				RefColumns: []*schema.Column{RepliesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reactions_users_reactions",
				Columns:    []*schema.Column{ReactionsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "reaction_user_id_post_id",
				Unique:  true,
				Columns: []*schema.Column{ReactionsColumns[7], ReactionsColumns[5]},
			},
			{
				Name:    "reaction_user_id_comment_id",
				Unique:  true,
				Columns: []*schema.Column{ReactionsColumns[7], ReactionsColumns[4]},
			},
			{
				Name:    "reaction_user_id_reply_id",
				Unique:  true,
				Columns: []*schema.Column{ReactionsColumns[7], ReactionsColumns[6]},
			},
		},
	}
	// RegionsColumns holds the columns for the "regions" table.
	RegionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "country_id", Type: field.TypeUUID},
	}
	// RegionsTable holds the schema information for the "regions" table.
	RegionsTable = &schema.Table{
		Name:       "regions",
		Columns:    RegionsColumns,
		PrimaryKey: []*schema.Column{RegionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "regions_countries_regions",
				Columns:    []*schema.Column{RegionsColumns[4]},
				RefColumns: []*schema.Column{CountriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RepliesColumns holds the columns for the "replies" table.
	RepliesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "comment_id", Type: field.TypeUUID},
		{Name: "author_id", Type: field.TypeUUID},
	}
	// RepliesTable holds the schema information for the "replies" table.
	RepliesTable = &schema.Table{
		Name:       "replies",
		Columns:    RepliesColumns,
		PrimaryKey: []*schema.Column{RepliesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "replies_comments_replies",
				Columns:    []*schema.Column{RepliesColumns[5]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "replies_users_replies",
				Columns:    []*schema.Column{RepliesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SiteDetailsColumns holds the columns for the "site_details" table.
	SiteDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Default: "SocialNet"},
		{Name: "email", Type: field.TypeString, Default: "kayprogrammer1@gmail.com"},
		{Name: "phone", Type: field.TypeString, Default: "+2348133831036"},
		{Name: "address", Type: field.TypeString, Default: "234, Lagos, Nigeria"},
		{Name: "fb", Type: field.TypeString, Default: "https://facebook.com"},
		{Name: "tw", Type: field.TypeString, Default: "https://twitter.com"},
		{Name: "wh", Type: field.TypeString, Default: "https://wa.me/2348133831036"},
		{Name: "ig", Type: field.TypeString, Default: "https://instagram.com"},
	}
	// SiteDetailsTable holds the schema information for the "site_details" table.
	SiteDetailsTable = &schema.Table{
		Name:       "site_details",
		Columns:    SiteDetailsColumns,
		PrimaryKey: []*schema.Column{SiteDetailsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "terms_agreement", Type: field.TypeBool, Default: false},
		{Name: "is_email_verified", Type: field.TypeBool, Default: false},
		{Name: "is_staff", Type: field.TypeBool, Default: false},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "bio", Type: field.TypeString, Nullable: true},
		{Name: "dob", Type: field.TypeTime, Nullable: true},
		{Name: "access", Type: field.TypeString, Nullable: true},
		{Name: "refresh", Type: field.TypeString, Nullable: true},
		{Name: "city_id", Type: field.TypeUUID, Nullable: true},
		{Name: "avatar_id", Type: field.TypeUUID, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_cities_users",
				Columns:    []*schema.Column{UsersColumns[16]},
				RefColumns: []*schema.Column{CitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_files_users",
				Columns:    []*schema.Column{UsersColumns[17]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserNotificationsColumns holds the columns for the "user_notifications" table.
	UserNotificationsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "notification_id", Type: field.TypeUUID},
	}
	// UserNotificationsTable holds the schema information for the "user_notifications" table.
	UserNotificationsTable = &schema.Table{
		Name:       "user_notifications",
		Columns:    UserNotificationsColumns,
		PrimaryKey: []*schema.Column{UserNotificationsColumns[0], UserNotificationsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_notifications_user_id",
				Columns:    []*schema.Column{UserNotificationsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_notifications_notification_id",
				Columns:    []*schema.Column{UserNotificationsColumns[1]},
				RefColumns: []*schema.Column{NotificationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserNotificationsReadColumns holds the columns for the "user_notifications_read" table.
	UserNotificationsReadColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "notification_id", Type: field.TypeUUID},
	}
	// UserNotificationsReadTable holds the schema information for the "user_notifications_read" table.
	UserNotificationsReadTable = &schema.Table{
		Name:       "user_notifications_read",
		Columns:    UserNotificationsReadColumns,
		PrimaryKey: []*schema.Column{UserNotificationsReadColumns[0], UserNotificationsReadColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_notifications_read_user_id",
				Columns:    []*schema.Column{UserNotificationsReadColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_notifications_read_notification_id",
				Columns:    []*schema.Column{UserNotificationsReadColumns[1]},
				RefColumns: []*schema.Column{NotificationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserMemberChatsColumns holds the columns for the "user_member_chats" table.
	UserMemberChatsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "chat_id", Type: field.TypeUUID},
	}
	// UserMemberChatsTable holds the schema information for the "user_member_chats" table.
	UserMemberChatsTable = &schema.Table{
		Name:       "user_member_chats",
		Columns:    UserMemberChatsColumns,
		PrimaryKey: []*schema.Column{UserMemberChatsColumns[0], UserMemberChatsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_member_chats_user_id",
				Columns:    []*schema.Column{UserMemberChatsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_member_chats_chat_id",
				Columns:    []*schema.Column{UserMemberChatsColumns[1]},
				RefColumns: []*schema.Column{ChatsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChatsTable,
		CitiesTable,
		CommentsTable,
		CountriesTable,
		FilesTable,
		FriendsTable,
		MessagesTable,
		NotificationsTable,
		OtpsTable,
		PostsTable,
		ReactionsTable,
		RegionsTable,
		RepliesTable,
		SiteDetailsTable,
		UsersTable,
		UserNotificationsTable,
		UserNotificationsReadTable,
		UserMemberChatsTable,
	}
)

func init() {
	ChatsTable.ForeignKeys[0].RefTable = FilesTable
	ChatsTable.ForeignKeys[1].RefTable = UsersTable
	ChatsTable.Annotation = &entsql.Annotation{}
	ChatsTable.Annotation.Checks = map[string]string{
		"dm_chat_constraints":    "(ctype = 'DM' AND name IS NULL AND description IS NULL AND image_id IS NULL) OR (ctype = 'GROUP')",
		"group_chat_constraints": "(ctype = 'GROUP' AND name IS NOT NULL) OR (ctype = 'DM')",
	}
	CitiesTable.ForeignKeys[0].RefTable = CountriesTable
	CitiesTable.ForeignKeys[1].RefTable = RegionsTable
	CommentsTable.ForeignKeys[0].RefTable = PostsTable
	CommentsTable.ForeignKeys[1].RefTable = UsersTable
	FriendsTable.ForeignKeys[0].RefTable = UsersTable
	FriendsTable.ForeignKeys[1].RefTable = UsersTable
	FriendsTable.Annotation = &entsql.Annotation{}
	FriendsTable.Annotation.Checks = map[string]string{
		"different_users": "requester_id <> requestee_id",
	}
	MessagesTable.ForeignKeys[0].RefTable = ChatsTable
	MessagesTable.ForeignKeys[1].RefTable = FilesTable
	MessagesTable.ForeignKeys[2].RefTable = UsersTable
	NotificationsTable.ForeignKeys[0].RefTable = CommentsTable
	NotificationsTable.ForeignKeys[1].RefTable = PostsTable
	NotificationsTable.ForeignKeys[2].RefTable = RepliesTable
	NotificationsTable.ForeignKeys[3].RefTable = UsersTable
	OtpsTable.ForeignKeys[0].RefTable = UsersTable
	PostsTable.ForeignKeys[0].RefTable = FilesTable
	PostsTable.ForeignKeys[1].RefTable = UsersTable
	ReactionsTable.ForeignKeys[0].RefTable = CommentsTable
	ReactionsTable.ForeignKeys[1].RefTable = PostsTable
	ReactionsTable.ForeignKeys[2].RefTable = RepliesTable
	ReactionsTable.ForeignKeys[3].RefTable = UsersTable
	RegionsTable.ForeignKeys[0].RefTable = CountriesTable
	RepliesTable.ForeignKeys[0].RefTable = CommentsTable
	RepliesTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = CitiesTable
	UsersTable.ForeignKeys[1].RefTable = FilesTable
	UserNotificationsTable.ForeignKeys[0].RefTable = UsersTable
	UserNotificationsTable.ForeignKeys[1].RefTable = NotificationsTable
	UserNotificationsReadTable.ForeignKeys[0].RefTable = UsersTable
	UserNotificationsReadTable.ForeignKeys[1].RefTable = NotificationsTable
	UserMemberChatsTable.ForeignKeys[0].RefTable = UsersTable
	UserMemberChatsTable.ForeignKeys[1].RefTable = ChatsTable
}
