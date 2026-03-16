package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Client struct {
	db   *sql.DB
	Path string
}

func Open(dbPath string) (*Client, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error while open database connection")
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if _, err = db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		return nil, fmt.Errorf("error while enable foreign keys")
	}

	return &Client{db: db, Path: dbPath}, nil
}

func (c *Client) Close() error {
	return c.db.Close()
}
