package cmd

import (
	"github.com/hyference/internal/container"
	"github.com/hyference/internal/server"
)

func Start() error {
	server.Server(container.New())
	return nil
}
