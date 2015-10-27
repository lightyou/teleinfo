package teleinfo

import (
	"github.com/tarm/serial"
)

type Info struct {
	port serial.Port
}

func (i *Info) init() {

}

func (i *Info) read() (string, string) {
	return "", ""
}
