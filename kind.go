package migrate

type Kind interface {
	PostgresKind
}

type PostgresKind interface {
	ToPostgresKind() string
}

type TextKind struct{}

type Time struct {
	withTimezone bool
}
