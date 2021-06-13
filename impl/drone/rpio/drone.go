package drone

import (
	"github.com/ngapung/phoenix/infra/config"
	"github.com/ngapung/phoenix/service"
	"github.com/stianeikeland/go-rpio"
)

func NewDrone(cfg config.DroneConfig) (service.Drone, error) {
	err := rpio.Open()
	if err != nil {
		return nil, err
	}
	return &drone{
		CWFront:  rpio.Pin(cfg.CWFrontPin),
		CCWFront: rpio.Pin(cfg.CCWFrontPin),
		CWBack:   rpio.Pin(cfg.CWBackPin),
		CCWBack:  rpio.Pin(cfg.CCWBackPin),
	}, nil
}

type drone struct {
	CWFront  rpio.Pin
	CCWFront rpio.Pin
	CWBack   rpio.Pin
	CCWBack  rpio.Pin
}

// SetAllSpeed set speed of all motor (0 - 1000)
func (d *drone) SetAllSpeed(speed int) {
	speed = speedStabilization(speed)
	d.CWFront.DutyCycle(uint32(speed), 1000)
	d.CCWFront.DutyCycle(uint32(speed), 1000)
	d.CWBack.DutyCycle(uint32(speed), 1000)
	d.CCWBack.DutyCycle(uint32(speed), 1000)
}

func (d *drone) SetCCWFrontSpeed(speed int) {
	speed = speedStabilization(speed)
	d.CCWFront.DutyCycle(uint32(speed), 1000)
}

func (d *drone) SetCWFrontSpeed(speed int) {
	speed = speedStabilization(speed)
	d.CWFront.DutyCycle(uint32(speed), 1000)
}

func (d *drone) SetCCWBackSpeed(speed int) {
	speed = speedStabilization(speed)
	d.CCWBack.DutyCycle(uint32(speed), 1000)
}

func (d *drone) SetCWBackSpeed(speed int) {
	speed = speedStabilization(speed)
	d.CWFront.DutyCycle(uint32(speed), 1000)
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