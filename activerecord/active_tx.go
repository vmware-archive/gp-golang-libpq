package activerecord

import (
	"database/sql"
	"log"
)

type activeTx struct {
	tx *sql.Tx
}

// newActiveTx return a *activeTx
func newActiveTx(db *sql.DB) (MockableTx, error) {
	atx := &activeTx{}
	var err error
	atx.tx, err = db.Begin()
	return atx, err
}

// Exec execute query in a transaction
func (m *activeTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.tx.Exec(query, args...)
}

// Rollback roll back a transaction
func (m *activeTx) Rollback() (error) {
	return m.tx.Rollback()
}

// ExecSQL commit a transaction
func (m *activeTx) Commit() (error) {
	return m.tx.Commit()
}

// GetRow execute query and return one row result
func (m *activeTx) GetRow(query string) (result map[string]interface{}, err error) {
	rows, err := m.GetRows(query)
	if err != nil || len(rows) == 0 {
		return nil, err
	}

	return rows[0], nil
}

// GetRows execute query and return multiple rows result
func (m *activeTx) GetRows(query string) (result []map[string]interface{}, err error) {
	rows, err := m.tx.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	return parseRows(rows)
}

// Exec execute query in a transaction
func (m *activeTx) Prepare(query string) (MockableStmt, error) {
	return newActiveStmt(m.tx, query)
}

type activeStmt struct {
	stmt *sql.Stmt
}

// newActiveStmt return a *activeStmt
func newActiveStmt(tx *sql.Tx, query string) (MockableStmt, error) {
	astmt := &activeStmt{}
	var err error
	astmt.stmt, err = tx.Prepare(query)
	return astmt, err
}

// Exec execute a statement
func (ast *activeStmt) Exec(args ...interface{}) (sql.Result, error) {
	return ast.stmt.Exec(args...)
}

// Close a active statement
func (ast *activeStmt) Close() (error) {
	return ast.stmt.Close()
}
