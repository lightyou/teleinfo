package main

import (
	"github.com/tarm/serial"
	"log"
	"teleinfo/frames"
	"net/http"
	"encoding/json"
)



func onField(name string, value string) {
	print(name, value, "\n")
}

func onFrame(map[string]string data) {
	string json = json.Marshall(data)
	print(json)
}

func main() {

}

