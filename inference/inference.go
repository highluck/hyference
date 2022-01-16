package inference

import (
	"github.com/hyference/config"
	"github.com/hyference/ml"
)

type Inference interface {
	Inference(input interface{}) (interface{}, error)
}

func New(config config.Config) Inference {
	switch config.GetMlLibType() {
	case ml.FastText:
		return NewFastText(config.ModelName, config.ModelPath, config.GetModule())
	default:
		return nil
	}
}
