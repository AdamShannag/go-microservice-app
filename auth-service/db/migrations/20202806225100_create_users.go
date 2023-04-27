package migrations

import (
	"github.com/go-rel/rel"
)

// MigrateCreateUsers definition
func MigrateCreateUsers(schema *rel.Schema) {
	schema.CreateTable("users", func(t *rel.Table) {
		t.ID("id")
		t.String("first_name")
		t.String("last_name")
		t.String("password")
		t.String("email")
		t.String("phone")
		t.String("token")
		t.String("user_type")
		t.String("refresh_token")
		t.DateTime("created_at")
		t.DateTime("updated_at")
	})

	schema.CreateIndex("users", "NAME_IDX", []string{"first_name", "last_name"})
	schema.CreateIndex("users", "FIRST_NAME_IDX", []string{"first_name"})
	schema.CreateIndex("users", "LAST_NAME_IDX", []string{"last_name"})
}

// RollbackCreateUsers definition
func RollbackCreateUsers(schema *rel.Schema) {
	schema.DropTable("users")
}
