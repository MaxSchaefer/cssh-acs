package main

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"flag"
	scribble "github.com/nanobox-io/golang-scribble"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type User struct {
	Username 	string 					`json:"username"`
	Password 	string 					`json:"password"`
	Config 		interface{} `json:"config"`
}

type AuthPasswordRequest struct {
	Username 		string `json:"username"`
	RemoteAddress 	string `json:"remoteAddress"`
	SessionID 		string `json:"sessionId"`
	Password 		string `json:"passwordBase64"`
}

type AuthResponse struct {
	Success	bool `json:"success"`
}

type ConfigRequest struct {
	Username 	string `json:"username"`
	SessionId 	string `json:"sessionId"`
}

type ConfigResponse struct {
	Config interface{} `json:"config"` // idk, not working with https://github.com/ContainerSSH/configuration/blob/f1696ce58c9d317bba1eb8afa250f678efdbb487/appconfig.go#L19 as a type
}

var db *scribble.Driver

func main() {
	log.SetOutput(os.Stdout)

	addr := flag.String("addr", "0.0.0.0:80", "Address for the http server")
	dbDir := flag.String("dbDir", "db", "Path to db directory")
	flag.Parse()

	_db, err := scribble.New(*dbDir, nil); db = _db; _db = nil
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/password", handleAuthPassword)
	http.HandleFunc("/config", handleConfig)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		panic(err)
	}
}

func handleAuthPassword(res http.ResponseWriter, req *http.Request) {
	authPasswordRequest := AuthPasswordRequest{}
	err := json.NewDecoder(req.Body).Decode(&authPasswordRequest)
	if err != nil {
		log.Error(err.Error(), authPasswordRequest)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	user := User{}
	if err := db.Read("users", authPasswordRequest.Username, &user); err != nil {
		log.Error(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	plainPassword, err := base64.StdEncoding.DecodeString(authPasswordRequest.Password)
	hashedPassword := sha512.Sum512(plainPassword)

	compare := user.Password == hex.EncodeToString(hashedPassword[:])

	log.WithField("user", user.Username).Info("does a auth request")

	response := AuthResponse{Success: compare}
	res.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(res).Encode(response)
}

func handleConfig(res http.ResponseWriter, req *http.Request) {
	configRequest := ConfigRequest{}
	err := json.NewDecoder(req.Body).Decode(&configRequest)
	if err != nil {
		log.Error(err.Error(), configRequest)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	user := User{}
	if err := db.Read("users", configRequest.Username, &user); err != nil {
		log.Error(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	log.WithField("user", user.Username).Info("does a config request")

	response := ConfigResponse{Config: user.Config}
	res.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(res).Encode(response)
}
