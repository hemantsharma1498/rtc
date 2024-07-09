package store

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

const dsn = "hemant:1@Million@tcp(localhost)/connection_balancer"
const migrationDir = "./store/migrations"

type ConnBal struct {
	Db   *sql.DB
	host string
	port string
}

func NewConnBalConnector() *ConnBal {
	return &ConnBal{}
}

func (c *ConnBal) Connect(ctx context.Context) (*ConnBal, error) {
	if c.Db == nil {
		var err error
		c.Db, err = initDb(ctx)
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	return c, nil
}

func initDb(ctx context.Context) (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	//  db, err = runMigrations(ctx, db, migrationDir)
	return db, nil
}

func runMigrations(ctx context.Context, db *sql.DB, migrationDir string) (*sql.DB, error) {
	if err := goose.RunContext(ctx, "status", db, migrationDir); err != nil {
		return nil, fmt.Errorf("failed to get goose status: %v", err)
	}

	if err := goose.RunContext(ctx, "up", db, migrationDir); err != nil {
		return nil, fmt.Errorf("failed to get goose up: %v", err)
	}
	return db, nil
}
