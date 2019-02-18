package activerecord

import "database/sql"

// FakeActiveRecord is only for test. You can change MockGetRows to return your own data
type FakeTx struct {
	*sql.Tx
	MockExec     func(query string, args ...interface{}) (sql.Result, error)
	MockRollback func() (error)
	MockCommit   func() (error)
	MockGetRow   func(query string) (result map[string]interface{}, err error)
	MockGetRows func(query string) (result []map[string]interface{}, err error)
	MockPrepare func(query string) (MockableStmt, error)
}

// NewFakeTx return a *FakeTx
func NewFakeTx() *FakeTx {
	return &FakeTx{
		&sql.Tx{},
		func(query string, args ...interface{}) (sql.Result, error) {
			return nil, nil
		},
		func() (error) {
			return nil
		},
		func() (error) {
			return nil
		},
		func(query string) (result map[string]interface{}, err error) {
			return nil, nil
		},
		func(query string) (result []map[string]interface{}, err error) {
			return nil, nil
		},
		func(query string) (MockableStmt, error) {
			return nil, nil
		},
	}
}

// Exec will call MockExec in FakeTx if it is set
func (m *FakeTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	if m.MockExec != nil {
		return m.MockExec(query, args...)
	}
	return nil, nil
}

// Rollback will call MockRollback in FakeTx if it is set
func (m *FakeTx) Rollback() (error) {
	if m.MockRollback != nil {
		return m.MockRollback()
	}
	return nil
}

// Commit will call MockCommit in FakeTx if it is set
func (m *FakeTx) Commit() (error) {
	if m.MockCommit != nil {
		return m.MockCommit()
	}
	return nil
}

// GetRow will call MockGetRow in FakeTx if it is set
func (m *FakeTx) GetRow(query string) (result map[string]interface{}, err error) {
	if m.MockGetRow != nil {
		return m.MockGetRow(query)
	}
	return nil, nil
}

// GetRows will call MockGetRows in FakeTx if it is set
func (m *FakeTx) GetRows(query string) (result []map[string]interface{}, err error) {
	if m.MockGetRows != nil {
		return m.MockGetRows(query)
	}
	return nil, nil
}

// Prepare will call MockPrepare in FakeTx if it is set
func (m *FakeTx) Prepare(query string) (MockableStmt, error) {
	if m.MockPrepare != nil {
		return m.MockPrepare(query)
	}
	return nil, nil
}

// FakeStmt is only for test.
type FakeStmt struct {
	MockExec  func(args ...interface{}) (sql.Result, error)
	MockClose func() (error)
}

// Exec will call MockExec in FakeStmt if it is set
func (m *FakeStmt) Exec(args ...interface{}) (sql.Result, error) {
	if m.MockExec != nil {
		return m.MockExec(args)
	}
	return nil, nil
}

// Close will call MockClose in FakeStmt if it is set
func (m *FakeStmt) Close() (error) {
	if m.MockClose != nil {
		return m.MockClose()
	}
	return nil
}
