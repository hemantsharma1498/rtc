package mysqlDb

import (
	"github.com/hemantsharma1498/rtc/server/types"
)

func (c *ConnBal) GetAllCommunicationServers() ([]*types.CommunicationServer, error) {
	// @TODO Note: Tentative function, will be removed once login and auth is configured
	rows, err := c.db.Query("SELECT org, address FROM communication_servers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	servers := make([]*types.CommunicationServer, 0)
	for rows.Next() {
		commServer := &types.CommunicationServer{}
		rows.Scan(&commServer.Organisation, &commServer.Address)
		servers = append(servers, commServer)
	}
	return servers, nil
}
