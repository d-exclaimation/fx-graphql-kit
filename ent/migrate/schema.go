// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ThoughtsColumns holds the columns for the "thoughts" table.
	ThoughtsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "body", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "image_url", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// ThoughtsTable holds the schema information for the "thoughts" table.
	ThoughtsTable = &schema.Table{
		Name:        "thoughts",
		Columns:     ThoughtsColumns,
		PrimaryKey:  []*schema.Column{ThoughtsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ThoughtsTable,
	}
)

func init() {
}
