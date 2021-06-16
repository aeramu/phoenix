package server

import (
	"github.com/gorilla/websocket"
	"github.com/ngapung/phoenix/infra/config"
	"github.com/ngapung/phoenix/infra/server"
	"github.com/ngapung/phoenix/service"
	log "github.com/sirupsen/logrus"
)

func NewServer(cfg config.WebsocketConfig, svc service.Service) (server.Server, error) {
	c, _, err := websocket.DefaultDialer.Dial(cfg.URL, nil)
	if err != nil {
		log.WithField("err", err).Fatalln("Failed dial websocket server")
	}
	return &handler{
		conn: c,
		svc:  svc,
	}, nil
}

type handler struct {
	conn *websocket.Conn
	svc  service.Service
}

func (h *handler) Run() {
	for {
		var req Request
		err := h.conn.ReadJSON(&req)
		if err != nil {
			log.WithField("err", err).Errorln("Failed read request")
		}
		switch req.Action {
		case "SET_MOTOR_SPEED":
			h.svc.SetMotorSpeed(
				Int(req.Payload["cw_front"]),
				Int(req.Payload["ccw_front"]),
				Int(req.Payload["cw_back"]),
				Int(req.Payload["ccw_back"]),
			)
		case "SET_ALL_MOTOR_SPEED":
			h.svc.SetAllMotorSpeed(Int(req.Payload["speed"]))
		}
	}
}

type Request struct {
	Action string `json:"action"`
	Payload map[string]interface{} `json:"payload"`
}

func Int(i interface{}) int {
	if i == nil {
		return 0
	}
	f, ok := i.(float64)
	if !ok {
		return 0
	}
	return int(f)
}