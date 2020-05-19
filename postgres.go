package migrate

import (
	"fmt"
	"strings"
)

type Postgres struct{}

func (p Postgres) CreateTable(name TableName, columns []Column) Query {
	var fields []string
	for _, column := range columns {
		kind := column.kind.ToPostgresKind()
		fields = append(fields, column.name+" "+kind)
	}
	query := fmt.Sprintf("CREATE TABLE %s (%s);", name, strings.Join(fields, ", "))

	return Query(query)
}

func (p Postgres) DropTable(name TableName) Query {
	return Query(fmt.Sprintf("DROP TABLE %s", name))
}

func (p Postgres) AddColumn(column Column) Query {
	kind := column.kind.ToPostgresKind()
	return Query(
		fmt.Sprintf(
			"ALTER TABLE %s ADD %s %s;",
			column.table, kind, column.name,
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
