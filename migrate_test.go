package migrate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMigration_CreateTable(t *testing.T) {
	backend := Postgres{}
	tests := []struct {
		name   string
		migration func(m *ChangeMigration)
		wantUp string
		wantDown string
	}{
		{
			name: "Postgres TEXT column",
			migration: func(m *ChangeMigration) {
				m.AddColumn(Column{
					table: "table",
					name:  "column",
					kind:  StringKind,
				})
			},
			want: `ALTER TABLE table ADD TEXT column;`,
			want: `ALTER TABLE table ADD TEXT column;`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ChangeMigration{
			}
			tt.migration(m)
			got := GenerateUp(*m, backend)
			assert.Equal(t, tt.want, got)
		})
	}
}