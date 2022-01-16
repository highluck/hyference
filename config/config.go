package config

import (
	"github.com/hyference/filesystem"
	"github.com/hyference/ml"
	"github.com/hyference/modules"
	"log"
	"strings"
)

type Config struct {
	Port                   int                     `json:"port"`
	DebugLabel             string                  `json:"debug_label"`
	MlLibType              string                  `json:"ml_lib_type"`
	ModelName              string                  `json:"model_name"`
	ModelPath              string                  `json:"model_path"`
	FileSystemClientType   string                  `json:"file_system_type"`
	FileSystemClientDetail filesystem.ClientDetail `json:"file_system_client_detail"`
	modules                *modules.Module
}

func (cfg Config) GetMlLibType() ml.MlLibType {
	if v, ok := ml.LibTypes[strings.ToLower(cfg.MlLibType)]; ok {
		return v
	}

	log.Fatalf("Not supported Lib %s", cfg.MlLibType)
	return ml.UnKnown
}

func (cfg Config) GetModule() *modules.Module {
	if cfg.modules == nil {
		cfg.modules = &modules.Module{}
	}
	return cfg.modules
}
