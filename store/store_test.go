package store

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type dataType struct {
	MOTDETAT string
	ADCO     string
	OPTARIF  string
	ISOUSC   string
}

var data = map[string]string{
	"MOTDETAT": "000000",
	"ADCO":     "020828337598",
	"OPTARIF":  "BBR(",
	"ISOUSC":   "45",
}

var expected = dataType{
	MOTDETAT: "000000",
	ADCO:     "020828337598",
	OPTARIF:  "BBR(",
	ISOUSC:   "12",
}

func TestMarshal(t *testing.T) {
	var js []byte
	js = Marshal(data)

	var jsonData dataType
	json.Unmarshal(js, &jsonData)

	if jsonData.MOTDETAT != data["MOTDETAT"] {
		t.Error("Wrong value ", jsonData.MOTDETAT, " for MOTDETAT ")
	}
	if jsonData.ADCO != data["ADCO"] {
		t.Error("Wrong value ", jsonData.ADCO, " for ADCO ")
	}
}

func TestStore(t *testing.T) {
	var result dataType
	http.HandleFunc("/teleinfo", func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		dec.Decode(&result)
		w.WriteHeader(200)
	})
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Printf("Can't start HTTP : ", err)
			t.Fatal("Can't start HTTP : ", err)
		}
	}()
	Store("http://localhost:8080/teleinfo", data)
	if result.MOTDETAT != expected.MOTDETAT {
		t.Error("Wrong value ", result.MOTDETAT, " for MOTDETAT ")
	}
	if result.ADCO != expected.ADCO {
		t.Error("Wrong value ", result.ADCO, " for ADCO ")
	}
}
