package container

import (
	"github.com/hyference/internal/config"
	"github.com/hyference/internal/modules"
)

type Container struct {
	Config  config.Config
	Modules modules.Module
}

func New(config config.Config) Container {
	return Container{
		Config:  config,
		Modules: modules.New(config),
	}
}
