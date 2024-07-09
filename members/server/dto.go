package server

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpReq struct {
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Organisation string `json:"org"`
}

type GetCommServerAddressReq struct {
	Org string `json:"org"`
}

type GetCommServerAddressRes struct {
	Address string `json:"address"`
}
