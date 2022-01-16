package inference

import (
	errors2 "github.com/hyference/errors"
	"github.com/hyference/internal/filesystem"
	"github.com/hyference/internal/modules"
	"github.com/hyference/ml/fasttext_wrapper"
	"github.com/rs/zerolog/log"
)

var _ Inference = &FastText{}

type FastText struct {
	model     *fasttext_wrapper.Model
	modelName string
	modelPath string
	client    filesystem.Client
}

func NewFastText(modelName string, modelPath string, module modules.Module) *FastText {
	inference := &FastText{
		client:    module.Client,
		modelName: modelName,
		modelPath: modelPath,
	}
	model, err := inference.initializeModel()
	if err != nil {
		log.Panic().Err(err).Send()
	}

	inference.model = model
	return inference
}

func (i *FastText) initializeModel() (*fasttext_wrapper.Model, error) {
	if err := i.client.DownloadModel(i.modelName, i.modelPath); err != nil {
		return nil, err
	}
	return fasttext_wrapper.New(i.modelName)
}

func (i *FastText) Inference(input interface{}) (interface{}, error) {

	op := errors2.GetMethodName()
	content, ok := input.(string)
	if ok == false || content == "" {
		return fasttext_wrapper.Result{}, errors2.Wrapper(errors2.New("contents empty"), op)
	}

	predict, err := i.model.Predict(content)
	if err != nil {
		return fasttext_wrapper.Result{}, errors2.Wrapper(err, op)
	}
	return predict, nil
}
