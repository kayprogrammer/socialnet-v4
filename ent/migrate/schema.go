// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CitiesColumns holds the columns for the "cities" table.
	CitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
	}
	// CitiesTable holds the schema information for the "cities" table.
	CitiesTable = &schema.Table{
		Name:       "cities",
		Columns:    CitiesColumns,
		PrimaryKey: []*schema.Column{CitiesColumns[0]},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString},
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
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
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
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString},
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
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ReactionsColumns holds the columns for the "reactions" table.
	ReactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "rtype", Type: field.TypeString},
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
				OnDelete:   schema.NoAction,
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
	// RepliesColumns holds the columns for the "replies" table.
	RepliesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString},
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
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "replies_users_replies",
				Columns:    []*schema.Column{RepliesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
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
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CitiesTable,
		CommentsTable,
		FilesTable,
		OtpsTable,
		PostsTable,
		ReactionsTable,
		RepliesTable,
		SiteDetailsTable,
		UsersTable,
	}
)

func init() {
	CommentsTable.ForeignKeys[0].RefTable = PostsTable
	CommentsTable.ForeignKeys[1].RefTable = UsersTable
	OtpsTable.ForeignKeys[0].RefTable = UsersTable
	PostsTable.ForeignKeys[0].RefTable = FilesTable
	PostsTable.ForeignKeys[1].RefTable = UsersTable
	ReactionsTable.ForeignKeys[0].RefTable = CommentsTable
	ReactionsTable.ForeignKeys[1].RefTable = PostsTable
	ReactionsTable.ForeignKeys[2].RefTable = RepliesTable
	ReactionsTable.ForeignKeys[3].RefTable = UsersTable
	RepliesTable.ForeignKeys[0].RefTable = CommentsTable
	RepliesTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = CitiesTable
	UsersTable.ForeignKeys[1].RefTable = FilesTable
}
