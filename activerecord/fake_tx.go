package activerecord

import "database/sql"

// FakeActiveRecord is only for test. You can change MockGetRows to return your own data
type FakeTx struct {
	*sql.Tx
	MockExec     func(query string, args ...interface{}) (sql.Result, error)
	MockRollback func() (error)
	MockCommit   func() (error)
}

// NewFakeTx return a *FakeTx
func NewFakeTx() *FakeTx {
	return &FakeTx{&sql.Tx{}, func(query string, args ...interface{}) (sql.Result, error) {
		return nil, nil
	}, func() (error) {
		return nil
	}, func() (error) {
		return nil
	}}
}

// Exec will call MockExec in FakeTx if it is set
func (m *FakeTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	if m.MockExec != nil {
		return m.MockExec(query, args)
	}
	return nil, nil
}

// Rollback will call Rollback in FakeTx if it is set
func (m *FakeTx) Rollback() (error) {
	if m.MockRollback != nil {
		return m.MockRollback()
	}
	return nil
}

// ExecSQL will call Commit if it is set
func (m *FakeTx) Commit() (error) {
	if m.MockCommit != nil {
		return m.MockCommit()
	}
	return nil
}
