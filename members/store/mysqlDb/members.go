package mysqlDb

import (
	"context"
	"database/sql"
	"errors"
	"members/store/types"
)

func (m *MembersDbConnector) CreateAccount(ctx context.Context, member *types.MemberAccount) (*types.MemberAccount, error) {
	rows, err := m.db.QueryContext(ctx, `
        INSERT INTO 
        members_account(email, first_name, last_name, org, password_hash, salt) 
        VALUES(?, ?, ?, ?, ?, ?)
        RETURNING id, email, first_name, last_name, org`,
		member.Email, member.FirstName, member.LastName, member.Org, member.PasswordHash, member.Salt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	createdMember := &types.MemberAccount{}
	for rows.Next() {
		rows.Scan(createdMember.Id, createdMember.Email, createdMember.FirstName, createdMember.LastName, createdMember.Org)
	}

	return createdMember, nil
}

func (m *MembersDbConnector) GetMembersByEmail(ctx context.Context, emailIds []string) ([]*types.MemberAccount, error) {
	if len(emailIds) == 0 {
		return nil, errors.New("no email ids to search against")
	}
	rows, err := m.db.QueryContext(ctx, `
        SELECT id, email, first_name, last_name, password_hash, salt
        FROM members_account
        WHERE email IN ?
        `, emailIds)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no account found for the given email")
		}
		return nil, err
	}
	defer rows.Close()
	accounts := make([]*types.MemberAccount, 0)
	for rows.Next() {
		account := &types.MemberAccount{}
		rows.Scan(&account.Id, &account.Email, account.FirstName, account.LastName, account.PasswordHash, account.Salt)
		accounts = append(accounts, account)
	}
	return accounts, nil
}
