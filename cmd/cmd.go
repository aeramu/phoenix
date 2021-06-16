package cmd

import (
	drone "github.com/ngapung/phoenix/impl/drone/serial"
	server "github.com/ngapung/phoenix/impl/server/websocket"
	"github.com/ngapung/phoenix/infra/config"
	"github.com/ngapung/phoenix/service"
	log "github.com/sirupsen/logrus"
)

func Run() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.WithField("err", err).Fatalln("Failed read config")
	}

	d, err := drone.NewDrone(cfg.Serial)
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

	log.Println("Program started.")
	srv.Run()
}
