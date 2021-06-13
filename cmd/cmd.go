package cmd

import (
	"fmt"
	drone "github.com/ngapung/phoenix/impl/drone/rpio"
	server "github.com/ngapung/phoenix/impl/server/websocket"
	"github.com/ngapung/phoenix/infra/config"
	"github.com/ngapung/phoenix/service"
	log "github.com/sirupsen/logrus"
)

func Run() {
	cfg := config.ReadConfig()

	fmt.Println(cfg.Websocket)
	fmt.Println(cfg.Drone)

	d, err := drone.NewDrone(cfg.Drone)
	if err != nil {
		log.WithField("err", err).Fatalln("Failed initiate drone")
	}

	svc := service.NewService(service.Adapter{
		Drone: d,
	})

	srv, err := server.NewServer(cfg.Websocket, svc)
	if err != nil {
		log.WithField("err", err).Fatalln("Failed initiate server")
	}

	srv.Run()
}
