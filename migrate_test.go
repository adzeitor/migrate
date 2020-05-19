package migrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeMigration_Postgres(t *testing.T) {
	backend := Postgres{}
	tests := []struct {
		name      string
		migration func(m *ChangeMigration)
		wantUp    []string
		wantDown  []string
	}{
		{
			name: "Postgres add columns",
			migration: func(m *ChangeMigration) {
				m.AddColumn(TextColumn("table", "name"))
				m.AddColumn(TimeColumn("table", "created_at"))
				m.AddColumn(TimeWithTimezoneColumn("table", "updated_at"))
			},
			wantUp: []string{
				`ALTER TABLE table ADD TEXT name;`,
				`ALTER TABLE table ADD TIME created_at;`,
				`ALTER TABLE table ADD TIME WITH TIME ZONE updated_at;`,
			},
			wantDown: []string{
				`ALTER TABLE table DROP COLUMN updated_at;`,
				`ALTER TABLE table DROP COLUMN created_at;`,
				`ALTER TABLE table DROP COLUMN name;`,
			},
		},
		{
			name: "Postgres create table",
			migration: func(m *ChangeMigration) {
				m.CreateTable(
					"user",
					TextColumn("table", "name"),
					TimeColumn("table", "created_at"),
					TimeWithTimezoneColumn("table", "updated_at"),
				)
			},
			wantUp: []string{
				`CREATE TABLE user (name TEXT, created_at TIME, updated_at TIME WITH TIME ZONE);`,
			},
			wantDown: []string{
				`DROP TABLE user`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ChangeMigration{}
			tt.migration(m)
			gotUp := GenerateUp(*m, backend)
			assert.Equal(t, tt.wantUp, gotUp)
			gotDown := GenerateDown(*m, backend)
			assert.Equal(t, tt.wantDown, gotDown)
		})
	}
}
