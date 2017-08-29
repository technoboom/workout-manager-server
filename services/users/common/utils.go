package common

import (
	"encoding/json"
	"log"
	"os"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HTTPStatus int    `json:"status"`
	}
	configuration struct {
		Server, ConnectionString, Database string
	}
)

// AppConfig - contains configuration of the application
var AppConfig configuration

func initConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[load config]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[decoding config]: %s\n", err)
	}
}
