package rqlite

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"net/url"
	"strings"
)

func init() {
	sql.Register("rqlite", &RQLiteDriver{})
}

// RQLiteDriver implements driver.Driver.
type RQLiteDriver struct {
}

// Open opens database and returns a new connection.
func (d *RQLiteDriver) Open(p string) (driver.Conn, error) {
	u, err :=  url.Parse(strings.TrimRight(p, "/"))
	if err != nil {
		return nil, err
	}

	node := NewNode(u)
	_, err = node.Status()
	if err != nil {
		return nil, err
	}

	return &RQLiteConn{
		node: node,
	}, nil
}

type RQLiteConn struct {
	node *Node
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

func (c *RQLiteConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return nil, nil
}