package db

import (
	"database/sql"
	"database/sql/driver"
	"golang.org/x/net/context"
)

type SQLInstance interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type SQLDb interface {
	SQLInstance
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Begin() (*sql.Tx, error)
}

type SQLTx interface {
	SQLInstance
	driver.Tx
}

