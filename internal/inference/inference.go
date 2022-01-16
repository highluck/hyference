package inference

import (
	"github.com/hyference/internal/config"
	"github.com/hyference/internal/modules"
	"github.com/hyference/ml"
)

type Inference interface {
	Inference(input interface{}) (interface{}, error)
}

func New(config config.Config, modules modules.Module) Inference {
	switch config.GetMlLibType() {
	case ml.FastText:
		return NewFastText(config.ModelName, config.ModelPath, modules)
	default:
		return nil
	}
}
