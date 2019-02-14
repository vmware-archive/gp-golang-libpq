package activerecord

import (
	"database/sql"
)

// MockableTx is a mock TX for test usage
type MockableTx interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Rollback() error
	Commit() error
	GetRow(query string) (result map[string]interface{}, err error)
	GetRows(query string) (result []map[string]interface{}, err error)
	Prepare(query string)(MockableStmt, error)
}

type MockableStmt interface {
	Exec(args ...interface{}) (sql.Result, error)
	Close() error
}
