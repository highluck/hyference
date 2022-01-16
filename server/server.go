package server

import (
	"github.com/hyference/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func Server(config config.Config) {
	var level zerolog.Level
	if config.DebugLabel == "warn" {
		level = zerolog.WarnLevel
	} else {
		level = zerolog.DebugLevel
	}

	log.Logger = log.Logger.Level(level).Output(os.Stdout)
	httpServer, err := NewHttpServer(config)
	if err != nil {
		log.Err(err).Msgf("HttpServer Init")
		panic(err)
	}

	httpServer.Start()
	httpServer.Stop()
}
