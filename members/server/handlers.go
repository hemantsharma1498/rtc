package server

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"members/pkg/proto"
	"members/store/types"
	"net/http"

	"golang.org/x/crypto/argon2"
)

const (
	saltSize int    = 16
	time     uint32 = 6
	memory   uint32 = 32
	keyLen   uint32 = 32
)

func (m *Members) GetCommServerAddress(w http.ResponseWriter, r *http.Request) {
	d := &GetCommServerAddressReq{}
	if err := decodeReqBody(r, d); err != nil {
		writeResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
	}
	res, err := m.ConnectionClient.GetCommServerAddr(r.Context(), &proto.GetCommServerAddrReq{Org: d.Org})
	if err != nil {
		log.Printf("Unable to fetch communication server for %s from connection-balancer\n", d.Org)
		writeResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
	}
	log.Printf("Addres for %s is %s", d.Org, res.Address)
	writeResponse(w, nil, &GetCommServerAddressRes{Address: res.Address}, http.StatusInternalServerError)

}

func (m *Members) Login(w http.ResponseWriter, r *http.Request) {
	d := &LoginReq{}
	if err := decodeReqBody(r, d); err != nil {
		writeResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
	}
	email := []string{d.Email}
	accounts, err := m.store.GetMembersByEmail(r.Context(), email)
	account := accounts[0]
	if err != nil {
		writeResponse(w, nil, err.Error(), http.StatusBadRequest)
	}
	enteredPasswordHash := createHash(d.Password, []byte(account.Salt))
	if enteredPasswordHash != account.PasswordHash {
		writeResponse(w, errors.New("entered pasword and stored password don't match"), nil, http.StatusBadRequest)
	}
	//@TODO
	//Add user details a jwt token, add token to header and return
}

func (m *Members) SignUp(w http.ResponseWriter, r *http.Request) {
	d := &SignUpReq{}
	if err := decodeReqBody(r, d); err != nil {
		writeResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
	}
	email := []string{d.Email}
	accounts, err := m.store.GetMembersByEmail(r.Context(), email)
	if err != nil && err.Error() != "no account found for the given email" {
		writeResponse(w, nil, err.Error(), http.StatusBadRequest)
	}
	if len(accounts) != 0 {
		writeResponse(w, errors.New("account exists for the given email"), nil, http.StatusBadRequest)
	}
	member := &types.MemberAccount{
		FirstName: d.FirstName,
		LastName:  d.LastName,
		Email:     d.Email,
		Org:       d.Organisation,
	}

	member, err = m.store.CreateAccount(r.Context(), member)
	if err != nil {
		writeResponse(w, errors.New("server encountered an error. please try again"), nil, http.StatusInternalServerError)
	}

	writeResponse(w, nil, "account created successfully", http.StatusOK)
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func createHash(password string, salt []byte) string {
	hash := argon2.Key([]byte(password), salt, time, memory, 8, keyLen)
	return base64.RawStdEncoding.EncodeToString(hash)
}

func decodeReqBody(r *http.Request, d any) error {
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return err
	}
	return nil
}

func writeResponse(w http.ResponseWriter, err error, msg any, httpStatus int) error {
	if err != nil {
		log.Printf("Error occured while decoding req json body: %s\n", err)
	}
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(msg)
}
