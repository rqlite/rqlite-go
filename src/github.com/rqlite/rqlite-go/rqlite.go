package rqlite

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

func init() {
	sql.Register("rqlite", &RQLiteDriver{})
}

// RQLiteDriver implements driver.Driver.
type RQLiteDriver struct {
}

// Open opens database and returns a new connection.
func (d *RQLiteDriver) Open(url string) (driver.Conn, error) {
	return &RQLiteConn{}, nil
}

type RQLiteConn struct {
}

func (c *RQLiteConn) Prepare(query string) (driver.Stmt, error) {
	return nil, nil
}

func (c *RQLiteConn) Close() error {
	return nil
}

func (c *RQLiteConn) Begin() (driver.Tx, error) {
	return nil, nil
}

func (c *RQLiteConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error)
	return nil, nil
}