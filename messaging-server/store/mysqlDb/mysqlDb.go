package mysqlDb

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hemantsharma1498/rtc/messaging/server/types"
	"github.com/hemantsharma1498/rtc/messaging/store"
)

const dsn = "hemant:1@Million@tcp(localhost)/communications"

func NewMessagingDbConnector() store.Connecter {
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

func (m *Messaging) CreateChannel(channelId, orgId int) (int, error) {
	row := m.db.QueryRow(`SELECT channel_id FROM channels WHERE channel_id = ?`, channelId)
	if row.Err() != nil {
		return -1, row.Err()
	}
	found := -1
	row.Scan(&found)
	// there's a new dm channel that's opened, save it
	if found != -1 {
		return -1, errors.New("channel exists for the given channel id")
	}
	row = m.db.QueryRow(`INSERT INTO channels(org_id, type) VALUES(?, ?) RETURNING channel_id`, orgId, "dm")
	row.Scan(channelId)
	return channelId, nil
}

func (m *Messaging) SaveMessage(message *types.Message) error {
	//@TODO name gets add when group chats are introduced
	strChan, _ := strconv.Atoi(message.ChannelId)
	strOrg, _ := strconv.Atoi(message.OrgId)
	_, err := m.CreateChannel(strChan, strOrg)
	if err != nil {
		if err.Error() != "channel exists for the given channel id" {
			return err
		}
	}
	_, err = m.db.Exec(`
	INSERT INTO messages(sender_id, message, channel_id, created_at)
	VALUES(?, ?, ?, ?)
	`, message.SenderId, message.Payload, message.ChannelId, message.CreatedAt)
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
	query := fmt.Sprintf(`
		SELECT 
			message_id, sender_id, name, message, channel_id
		FROM messages
		WHERE channel_id IN (%s)
	`, placeHolders)
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
