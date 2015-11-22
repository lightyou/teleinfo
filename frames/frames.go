package frames

import (
	"github.com/tarm/serial"
	"log"
)

type Info struct {
	port    *serial.Port
	fieldcb func(string, string)
	framecb func(map[string]string)
	frame   map[string]string
	state   string
	current string
	key     string
	value   string
	ck      string
	cks     int
}

func (i *Info) init(device string) error {
	c := &serial.Config{Name: device, Baud: 4200}
	port, err := serial.OpenPort(c)
	if err != nil {
		log.Print(err)
	}
	i.port = port
	i.state = "INIT"
	i.cks = 0
	i.frame = map[string]string{}
	return err
}

func (i *Info) read() (string, string) {
	return "", ""
}

func (i *Info) SetFieldCB(cb func(string, string)) {
	i.fieldcb = cb
}

func (i *Info) SetFrameCB(cb func(map[string]string)) {
	i.framecb = cb
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
		// compare checksum
		i.framecb(i.frame)
		i.frame = map[string]string{}
		i.cks = 0
		i.state = "KEY"
	case '\n':
		i.decodeSeparator(b)
		// compare checksum
		if i.ck == string((i.cks&0x3F)+0x20) {
			i.fieldcb(i.key, i.value)
			i.frame[i.key] = i.value
		}
		i.cks = 0
		i.state = "KEY"
	case ' ':
		i.decodeSeparator(b)
	default:
		i.decodeField(b)
	}
}
