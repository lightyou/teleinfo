package store

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Marshal(data interface{}) []byte {
	js, _ := json.Marshal(data)
	return js
}

func Store(url string, data interface{}) {
	_, err := http.Post(url, "application/json", bytes.NewReader(Marshal(data)))
	if err != nil {
		log.Fatal("HTTP POST Error : ", err)
		return
	}

	// Manage response
}
