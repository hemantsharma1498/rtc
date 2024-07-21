package mysqlDb

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hemantsharma1498/rtc/store"
	"github.com/pressly/goose/v3"
)

const dsn = "hemant:1@Million@tcp(localhost)/connection_balancer"
const migrationDir = "./store/migrations"

type ConnBal struct {
	db *sql.DB
}

func NewConnBalConnector() store.Connecter {
	return &ConnBal{}
}

func (c *ConnBal) Connect() (store.Storage, error) {
	if c.db == nil {
		var err error
		c.db, err = initDb()
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	return c, nil
}

func initDb() (*sql.DB, error) {

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
