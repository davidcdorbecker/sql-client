package sql_client

import "database/sql"

type sqlRow struct {
	row *sql.Row
}

type row interface {
	Scan(dest ...interface{}) error
}

func (r *sqlRow) Scan(dest ...interface{}) error {
	return r.row.Scan(dest...)
}
