package modules

import (
	"github.com/hyference/internal/config"
	"github.com/hyference/internal/filesystem"
)

type Module struct {
	Client filesystem.Client
}

func New(cfg config.Config) Module {
	return Module{
		Client: filesystem.New(cfg.FileSystemClientType, cfg.FileSystemClientDetail),
	}
}
