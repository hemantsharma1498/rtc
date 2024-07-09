package mysqlDb

import (
	"database/sql"
	"members/store"
	_ "github.com/go-sql-driver/mysql"
)

const dsn = "hemant:1@Million@tcp(localhost)/members"

func NewMembersDbConnector() store.Connecter {
	return &MembersDbConnector{}
}

type MembersDbConnector struct {
	db *sql.DB
}

func (m *MembersDbConnector) Connect() (store.Storage, error) {
	if m.db == nil {
		var err error
		m.db, err = initDb()
		if err != nil {
			return nil, err
		}
		return m, nil
	}
	return m, nil
}

func initDb() (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

