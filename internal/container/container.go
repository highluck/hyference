package container

import (
	config2 "github.com/hyference/internal/config"
	modules2 "github.com/hyference/internal/modules"
)

type Container struct {
	Config  config2.Config
	Modules modules2.Module
}

func New() Container {
	config := config2.Config{}
	modules := modules2.New(config)

	return Container{
		Config: config,
		Modules: modules,
	}
}