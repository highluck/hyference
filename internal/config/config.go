package config

import (
	"github.com/hyference/internal/filesystem"
	"github.com/hyference/ml"
	"log"
	"strings"
)

type Config struct {
	Port                   int                     `json:"port"`
	DebugLabel             string                  `json:"debugLabel"`
	MlLibType              string                  `json:"mlLibType"`
	ModelName              string                  `json:"modelName"`
	ModelPath              string                  `json:"modelPath"`
	FileSystemClientType   string                  `json:"fileSystemType"`
	FileSystemClientDetail filesystem.ClientDetail `json:"fileSystemClientDetail"`
	ParameterType          string                  `json:"parameterType"`
	Uri                    string                  `json:"uri"`
}

type ParameterType string

const StringType = ParameterType("string")
const IntType = ParameterType("int")
const LongType = ParameterType("long")
const FloatType = ParameterType("float")

func (cfg Config) GetMlLibType() ml.LibType {
	if v, ok := ml.LibTypes[strings.ToLower(cfg.MlLibType)]; ok {
		return v
	}

	log.Fatalf("Not supported Lib %s", cfg.MlLibType)
	return ml.UnKnown
}
