package main

import (
	log "github.com/Sirupsen/logrus"
	rand "crypto/rand"
	hex "encoding/hex"
)

type AuthReq struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LogoutReq struct {
	User string  `json:"user" binding:"required"`
	Token string `json:"token" binding:"required"`
}

var user_lookup map[string]string
var active_tokens map[string]string

func init() {
	user_lookup = make(map[string]string)
	active_tokens = make(map[string]string)
	user_lookup["Foo"] = "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
	user_lookup["Bar"] = "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
}

func VerifyUserToken(user string, token string) bool {
	t, success := active_tokens[user]
	return success && t == token
}

func AuthUser(auth_req *AuthReq) (string, bool) {
	l := log.WithField("User", auth_req.User)
	l.Info("Authing user.")
	hash,user_exists := user_lookup[auth_req.User]
	l.Info("User found: ", user_exists)

	if user_exists && hash == auth_req.Password {
		t := RandomHexToken()
		active_tokens[auth_req.User] = t
		l.Info("Generated token: ", t)
		return t, true
	} else {
		l.Info("Auth failed.")
		return "", false
	}
}

func  GetUsers() []string {
	users := make([]string, len(active_tokens))

	i := 0
	for k,_ := range active_tokens {
		users[i] = k
		i++
	}
	
	return users
}

func LogoutUser(logout_req *LogoutReq) bool {
	_, exists := active_tokens[logout_req.User]
	if !exists {
		log.WithField("User", logout_req.User).Info("Log out failed.")
		return false
	}
	delete(active_tokens, logout_req.User)
	log.WithField("User", logout_req.User).Info("Log out successful.")
	return true
}

func RandomHexToken() string {
	token := RandomByteSlice(8)
	return hex.EncodeToString(token)
}

func RandomByteSlice(size int) []byte {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("Error generating random byte slice:", err)
	}
	return b
}
