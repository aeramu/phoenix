package service

type Adapter struct {
	Drone Drone
}

type Drone interface {
	SetAllSpeed(speed int)
	SetCCWFrontSpeed(speed int)
	SetCWFrontSpeed(speed int)
	SetCCWBackSpeed(speed int)
	SetCWBackSpeed(speed int)
}
