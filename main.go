package main

import (
	_ "fmt"
	_ "log"
	"strconv"
	"teleinfo/frames"
	"teleinfo/store"
	"time"
)

type energy struct {
	IINST int
	PAPP  int
	PTEC  string
}
type counters struct {
	BBRHCJB int
	BBRHPJB int
	BBRHCJW int
	BBRHPJW int
	BBRHCJR int
	BBRHPJR int
}
type date struct {
	Day   int
	Month int
	Year  int
	DoW   int
}
type payload struct {
	Energy energy
	Counters counters
	Date date
}

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

func onFrame(data map[string]string) {
	d := payload{}

	intValue, err := strconv.Atoi(data["IINST"])
	if err == nil {
		d.Energy.IINST = intValue
	}
	intValue, err = strconv.Atoi(data["PAPP"])
	if err == nil {
		d.Energy.PAPP = intValue
	}
	d.Energy.PTEC = data["PTEC"]

	intValue, err = strconv.Atoi(data["BBRHCJB"])
	if err == nil {
		d.Counters.BBRHCJB = intValue
	}
	intValue, err = strconv.Atoi(data["BBRHPJB"])
	if err == nil {
		d.Counters.BBRHPJB = intValue
	}
	intValue, err = strconv.Atoi(data["BBRHCJW"])
	if err == nil {
		d.Counters.BBRHCJW = intValue
	}
	intValue, err = strconv.Atoi(data["BBRHPJW"])
	if err == nil {
		d.Counters.BBRHPJW = intValue
	}
	intValue, err = strconv.Atoi(data["BBRHCJR"])
	if err == nil {
		d.Counters.BBRHCJR = intValue
	}
	intValue, err = strconv.Atoi(data["BBRHPJR"])
	if err == nil {
		d.Counters.BBRHPJR = intValue
	}

	t := time.Now()
	d.Date.Day = t.Day()
	d.Date.Month = int(t.Month())
	d.Date.Year = t.Year()
	d.Date.DoW = int(t.Weekday())

	store.Store("http://192.168.1.4:3133/teleinfo", d)
}

func main() {
	var ti frames.Info
	ti.Init("/dev/ttyUSB0")
	ti.SetFieldCB(onField)
	ti.SetFrameCB(onFrame)
	ti.Run()
}
