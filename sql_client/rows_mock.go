package sql_client

import (
	"errors"
	"fmt"
)

type rowMock struct {
	Columns []string
	Row     []interface{}
}

func (r *rowMock) Scan(dest ...interface{}) error {
	if len(r.Row) != len(dest) {
		return errors.New("invalid destination interface")
	}

	fmt.Println(dest)

	for key, value := range r.Row {
		// TODO: make the scan functionality
		fmt.Println(key, value)
	}

	return nil
}
