package config

import "gopkg.in/gcfg.v1"

var (
	path = []string{
		"/etc/phoenix/",
		"files/config/",
	}
	filename = "main.ini"
)

func ReadConfig() Config {
	var cfg Config
	for _, p := range path {
		err := gcfg.ReadFileInto(&cfg, p+filename)
		if err == nil {
			break
		}
	}
	return cfg
}

type Config struct {
	Websocket WebsocketConfig
	Drone     DroneConfig
}

type WebsocketConfig struct {
	URL string
}

type DroneConfig struct {
	CWFrontPin  uint8
	CCWFrontPin uint8
	CWBackPin   uint8
	CCWBackPin  uint8
}