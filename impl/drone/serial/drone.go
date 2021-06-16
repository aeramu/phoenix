package drone

import (
	"fmt"
	"github.com/ngapung/phoenix/service"
	"github.com/tarm/serial"
	"io/ioutil"
	"strings"
)

func NewDrone() (service.Drone, error) {
	port, err := serial.OpenPort(&serial.Config{
		Name:        findArduino(),
		Baud:        9600,
	})
	if err != nil {
		return nil, err
	}
	return &drone{
		Port: port,
	}, nil
}

func findArduino() string {
	contents, _ := ioutil.ReadDir("/dev")

	// Look for what is mostly likely the Arduino device
	for _, f := range contents {
		if strings.Contains(f.Name(), "tty.usbserial") ||
			strings.Contains(f.Name(), "ttyUSB") {
			return "/dev/" + f.Name()
		}
	}

	return ""
}

type drone struct {
	Port *serial.Port
}

func (d *drone) SetAllSpeed(speed int) {
	d.SetCCWFrontSpeed(speed)
	d.SetCWFrontSpeed(speed)
	d.SetCCWBackSpeed(speed)
	d.SetCWBackSpeed(speed)
}

func (d *drone) SetCCWFrontSpeed(speed int) {
	d.Port.Write(serialFormat("ccw_front", speed))
}

func (d *drone) SetCWFrontSpeed(speed int) {
	d.Port.Write(serialFormat("cw_front", speed))
}

func (d *drone) SetCCWBackSpeed(speed int) {
	d.Port.Write(serialFormat("ccw_back", speed))
}

func (d *drone) SetCWBackSpeed(speed int) {
	d.Port.Write(serialFormat("cw_back", speed))
}

func serialFormat(motor string, speed int) []byte {
	return []byte(fmt.Sprintf("%s %d", motor, speedStabilization(speed)))
}

func speedStabilization(speed int) int {
	if speed < 0 {
		speed = 0
	}
	if speed > 1000 {
		speed = 1000
	}
	return speed
}