package main

import (
	"github.com/tarm/serial"
	"log"
	"teleinfo/frames"
)

func onField(name string, value string) {
	print(name, value, "\n")
}

func main() {
	var ti frames.Info
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 4700}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	ti.SetFieldCB(onField)

	b := make([]byte, 1)
	for {
		s.Read(b)
		ti.Decode(b[0])
	}
}
