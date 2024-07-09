package types

//@API level representation of data stored in Mysql db

type MemberAccount struct {
	Id           int64
	Email        string
	FirstName    string
	LastName     string
	Org          string
	PasswordHash string
	Salt         string
}
