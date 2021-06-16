package config

import (
	"errors"
	"gopkg.in/gcfg.v1"
)

var (
	path = []string{
		"/etc/phoenix/",
		"files/config/",
	}
	filename = "main.ini"
)

func ReadConfig() (Config, error) {
	var cfg Config
	for _, p := range path {
		err := gcfg.ReadFileInto(&cfg, p+filename)
		if err != nil {
			continue
		} else {
			return cfg, nil
		}
	}
	return cfg, errors.New("no valid config file in directory")
}

type Config struct {
	Websocket WebsocketConfig
	GPIO      GPIOConfig
	Serial    SerialConfig
}

type WebsocketConfig struct {
	URL string
}

type GPIOConfig struct {
	CWFrontPin  uint8
	CCWFrontPin uint8
	CWBackPin   uint8
	CCWBackPin  uint8
}

type SerialConfig struct {
	PortName string
}