package database

import (
	"comm/database/client"
	"context"
	"database/sql"
)

type Database struct {
	conn   *sql.DB
	client *client.Queries
	ctx    context.Context
}

func Connect(connstring string) (*Database, error) {
	db := &Database{}

	conn, err := sql.Open("postgres", connstring)
	if err != nil {
		return &Database{}, err
	}

	if err := conn.Ping(); err != nil {
		return &Database{}, err
	}

	db.conn = conn
	db.client = client.New(conn)
	db.ctx = context.Background()

	return db, nil
}

func (db Database) Ctx() context.Context {
	return db.ctx
}

func (db *Database) Conn() *client.Queries {
	return db.client
}

func (db *Database) Close() {
	db.conn.Close()
}
