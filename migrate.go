package migrate

import "strings"

type ColumnKind int

const (
	TextKind ColumnKind = 1
	IntKind    ColumnKind = 2
)

type TableName string

type Column struct {
	table TableName
	name  string
	kind  ColumnKind
}

type Query string

// FIXME: naming
type Backend interface {
	CreateTable(name TableName, columns ...Column) Query
	DropTable(name TableName, columns ...Column) Query
	AddColumn(Column) Query
	RemoveColumn(Column) Query
}

type ChangeMigration struct {
	up      []Action
	down    []Action
}

type Action func(Backend) Query


func (m *ChangeMigration) add(up Action, down Action) {
	m.up = append(m.up, up)
	m.down = append(m.down, down)
}

func (m *ChangeMigration) AddColumn(column Column) {
	m.add(
		func(backend Backend) Query {
			return backend.AddColumn(column)
		},
		func(backend Backend) Query {
			return backend.RemoveColumn(column)
		},
	)
}

func GenerateUp(migration ChangeMigration, backend Backend) string {
	// FIXME: builder
	var result []string
	for _, action := range migration.up {
		result = append(result, string(action(backend)))
	}
	return strings.Join(result, "\n")
}