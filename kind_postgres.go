package migrate

func (k TextKind) ToPostgresKind() string {
	return "TEXT"
}

func (t Time) ToPostgresKind() string {
	if t.withTimezone {
		return "TIME WITH TIME ZONE"
	} else {
		return "TIME"
	}
}
