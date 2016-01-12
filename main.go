package main

import (
	_ "log"
	"teleinfo/frames"
	"teleinfo/store"
	_ "fmt"
	"time"
	"strconv"
)

func onPTEC(value string) {
	switch value[:2] {
	case "HP":
//		fmt.Println("Heures pleines")
	case "HC":
//		fmt.Println("Heures creuses")
	}
	switch value[2:4] {
	case "JB":
//		fmt.Println("Bleu")
	case "JW":
//		fmt.Println("Blanc")
	case "JR":
//		fmt.Println("Rouge")
	}
}

func onField(name string, value string) {
	switch name {
		case "PTEC":
			onPTEC(value)
	}
	//fmt.Printf("field %s => %s\n", name, value)
}

func setFieldIntValue(data map[string]int, field string, value string) {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		data[field] = intValue
	}
}


func onFrame(data map[string]string) {
	// Add date info
	t := time.Now()
	d := map[string]map[string]int{}
	d["Energy"] = make(map[string]int)
	setFieldIntValue(d["Energy"], "IINST", data["IINST"])
	setFieldIntValue(d["Energy"], "PAPP", data["PAPP"])

	d["Counters"] = make(map[string]int)
	setFieldIntValue(d["Counters"], "BBRHCJB", data["BBRHCJB"])
	setFieldIntValue(d["Counters"], "BBRHPJB", data["BBRHPJB"])
	setFieldIntValue(d["Counters"], "BBRHCJW", data["BBRHCJW"])
	setFieldIntValue(d["Counters"], "BBRHPJW", data["BBRHPJW"])
	setFieldIntValue(d["Counters"], "BBRHCJR", data["BBRHCJR"])
	setFieldIntValue(d["Counters"], "BBRHPJR", data["BBRHPJR"])

	d["Date"] = make(map[string]int)
	d["Date"]["Day"] = t.Day()
	d["Date"]["Month"] = int(t.Month())
	d["Date"]["Year"] = t.Year()
	d["Date"]["DoW"] = int(t.Weekday())
	store.Store("http://192.168.1.4:3133/teleinfo", d)
}

func main() {
	var ti frames.Info
	ti.Init("/dev/ttyUSB0")
	ti.SetFieldCB(onField)
	ti.SetFrameCB(onFrame)
	ti.Run()
}
