package server

import (
	"github.com/hyference/internal/container"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func Server(container container.Container) {
	var level zerolog.Level
	if container.Config.DebugLabel == "warn" {
		level = zerolog.WarnLevel
	} else {
		level = zerolog.DebugLevel
	}

	log.Logger = log.Logger.Level(level).Output(os.Stdout)
	httpServer, err := NewHttpServer(container)
	if err != nil {
		log.Err(err).Msgf("HttpServer Init")
		panic(err)
	}

	httpServer.Start()
	httpServer.Stop()
}
