package frames

import (
	"github.com/tarm/serial"
	"log"
)

type Info struct {
	port    *serial.Port
	cb      func(string, string)
	state   string
	current string
	key     string
	value   string
	ck      string
	cks     int
}

func (i *Info) init() error {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 4200}
	port, err := serial.OpenPort(c)
	i.port = port
	if err != nil {
		log.Fatal(err)
	}
	i.state = "INIT"
	i.cks = 0
	return err
}

func (i *Info) read() (string, string) {
	return "", ""
}

func (i *Info) SetCB(cb func(string, string)) {
	i.cb = cb
}

func (i *Info) decodeSeparator(b byte) {
	switch i.state {
	case "KEY":
		i.key = i.current
		i.current = ""
		i.state = "VALUE"
		if b == ' ' {
			i.cks = i.cks + int(b)
		}
	case "VALUE":
		i.value = i.current
		i.current = ""
		i.state = "CK"
	case "CK":
		i.ck = i.current
		i.current = ""
		i.state = "KEY"
	}
}

func (i *Info) decodeField(b byte) {
	i.current = i.current + string(b)
	switch i.state {
	case "KEY":
		i.cks = i.cks + int(b)
	case "VALUE":
		i.cks = i.cks + int(b)
	}
}

func (i *Info) Decode(b byte) {
	switch b {
	case '\002':
		i.state = "KEY"
	case '\003':
		if i.ck == string((i.cks&0x3F)+0x20) {
			i.cb(i.key, i.value)
		}
		i.cks = 0
		i.state = "KEY"
	case '\n':
		i.decodeSeparator(b)
		if i.ck == string((i.cks&0x3F)+0x20) {
			i.cb(i.key, i.value)
		}
		i.cks = 0
		i.state = "KEY"
	case ' ':
		i.decodeSeparator(b)
	default:
		i.decodeField(b)
	}
}
