// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SiteDetailsColumns holds the columns for the "site_details" table.
	SiteDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
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
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SiteDetailsTable,
	}
)

func init() {
}