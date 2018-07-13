package activerecord

import (
	"database/sql"
)

// MockableTx is a mock TX for test usage
type MockableTx interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Rollback() error
	Commit() error
}
