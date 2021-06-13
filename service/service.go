package service

type Service interface {
	SetMotorSpeed(cwFront, ccwFront, cwBack, ccwBack int) error
	SetAllMotorSpeed(speed int) error
}

func NewService(adapter Adapter) Service {
	return &service{adapter: adapter}
}

type service struct {
	adapter Adapter
}

func (s *service) SetMotorSpeed(cwFront, ccwFront, cwBack, ccwBack int) error {
	s.adapter.Drone.SetCWFrontSpeed(cwFront)
	s.adapter.Drone.SetCWBackSpeed(cwBack)
	s.adapter.Drone.SetCCWFrontSpeed(ccwFront)
	s.adapter.Drone.SetCCWBackSpeed(ccwBack)
	return nil
}

func (s *service) SetAllMotorSpeed(speed int) error {
	s.adapter.Drone.SetAllSpeed(speed)
	return nil
}
