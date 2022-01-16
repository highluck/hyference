package config

import (
	"github.com/hyference/internal/filesystem"
	"github.com/hyference/ml"
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
	ParameterType          string                  `json:"parameter_type"`
	Uri                    string                  `json:"uri"`
}

type ParameterType string

const StringType = ParameterType("string")
const IntType = ParameterType("int")
const LongType = ParameterType("long")
const FloatType = ParameterType("float")

func (cfg Config) GetMlLibType() ml.MlLibType {
	if v, ok := ml.LibTypes[strings.ToLower(cfg.MlLibType)]; ok {
		return v
	}

	log.Fatalf("Not supported Lib %s", cfg.MlLibType)
	return ml.UnKnown
}
