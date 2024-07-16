package mysqlDb

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hemantsharma1498/rtc/store"
	"github.com/hemantsharma1498/rtc/store/types"
)

const dsn = "hemant:1@Million@tcp(localhost)/members"

func NewMembersDbConnector() store.Connecter {
	return &Messaging{}
}

type Messaging struct {
	db *sql.DB
}

func (m *Messaging) Connect() (store.Storage, error) {
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

func (m *Messaging) SaveMessage(message *types.Message) error {
	//@TODO name gets add when group chats are introduced
	_, err := m.db.Exec(`
	INSERT INTO messages(sender_id, message, channel_id)
	VALUES(?, ?, ?)
	`, message.Sender, message.Payload, message.Channel)
	if err != nil {
		return err
	}
	return nil
}

func (m *Messaging) GetMessages(channelIds []int) ([]*types.Message, error) {
	if len(channelIds) == 0 {
		return nil, errors.New("empty slice sent")
	}
	placeHolders, params, err := getPlaceholders(channelIds)
	query := `
		SELECT 
		message_id, sender_id, name, message, channel_id
		FROM messages
		WHERE channel_id IN 
	` + placeHolders
	fmt.Printf("GetMessages query: %s\n", query)

	rows, err := m.db.Query(query, params...)
	if err != nil {
		log.Printf("error occured in GetMessages: %s\n", err)
		return nil, err
	}
	defer rows.Close()
	return nil, nil
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

func getPlaceholders(params interface{}) (string, []interface{}, error) {
	v := reflect.ValueOf(params)
	if v.Kind() != reflect.Slice {
		return "", nil, fmt.Errorf("expected a slice, got %T", params)
	}
	n := v.Len()
	p := make([]string, n)
	parameters := make([]interface{}, n)
	for i := 0; i < n; i++ {
		p[i] = "?"
		parameters[i] = v.Index(i).Interface()
	}
	pStr := strings.Join(p, ",")
	return pStr, parameters, nil
}
