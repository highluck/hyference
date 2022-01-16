package server

import (
	"log"
	"strings"
)

type Config struct {
	Port       int    `json:"port"`
	DebugLabel string `json:"debug_label"`
	MlLibType  string `json:"ml_lib_type"`
}

type MlLibType string
type MlLibTypeMap map[string]MlLibType

var libTypes = MlLibTypeMap{
	"fasttext": FastText,
}

const FastText = MlLibType("fasttext")
const UnKnown = MlLibType("unknown")

func (cfg Config) GetMlLibType() MlLibType {
	if v, ok := libTypes[strings.ToLower(cfg.MlLibType)]; ok {
		return v
	}

	log.Fatalf("Not supported Lib %s", cfg.MlLibType)
	return UnKnown
}
