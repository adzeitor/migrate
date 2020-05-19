package migrate

import "fmt"

type Postgres struct {
}

func (p Postgres) Kind(kind ColumnKind) string {
	switch kind {
	case StringKind:
		return "TEXT"
	default:
		return string(kind)
	}
}
func (p Postgres) CreateTable(name TableName, columns ...Column) Query {
	panic("implement me")
}

func (p Postgres) DropTable(name TableName, columns ...Column) Query {
	panic("implement me")
}

func (p Postgres) AddColumn(column Column) Query {
	return Query(
		fmt.Sprintf(
			"ALTER TABLE %s ADD %s %s;",
			column.table, p.Kind(column.kind), column.name,
		),
	)
}

func (p Postgres) RemoveColumn(column Column) Query {
	return Query(
		fmt.Sprintf(
			"ALTER TABLE %s DROP COLUMN %s;",
			column.table, column.name,
		),
	)
}
