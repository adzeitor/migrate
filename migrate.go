package migrate

type TableName string

type Column struct {
	table TableName
	name  string
	kind  Kind
}

func TextColumn(tableName TableName, name string) Column {
	return Column{
		table: tableName,
		name:  name,
		kind:  TextKind{},
	}
}

func TimeColumn(tableName TableName, name string) Column {
	return Column{
		table: tableName,
		name:  name,
		kind:  Time{},
	}
}

func TimeWithTimezoneColumn(tableName TableName, name string) Column {
	return Column{
		table: tableName,
		name:  name,
		kind:  Time{withTimezone: true},
	}
}

type Query string

type Backend interface {
	CreateTable(name TableName, columns []Column) Query
	DropTable(name TableName) Query
	AddColumn(Column) Query
	RemoveColumn(Column) Query
}

type ChangeMigration struct {
	up   []Action
	down []Action
}

type Action func(Backend) Query

func (m *ChangeMigration) add(up Action, down Action) {
	m.up = append(m.up, up)
	m.down = append([]Action{down}, m.down...)
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

func (m *ChangeMigration) CreateTable(name TableName, columns ...Column) {
	m.add(
		func(backend Backend) Query {
			return backend.CreateTable(name, columns)
		},
		func(backend Backend) Query {
			return backend.DropTable(name)
		},
	)
}

func generate(actions []Action, backend Backend) []string {
	var result []string
	for _, action := range actions {
		result = append(result, string(action(backend)))
	}
	return result
}

func GenerateUp(migration ChangeMigration, backend Backend) []string {
	return generate(migration.up, backend)
}

func GenerateDown(migration ChangeMigration, backend Backend) []string {
	return generate(migration.down, backend)
}
