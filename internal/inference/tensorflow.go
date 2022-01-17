package inference

import (
	"github.com/hyference/internal/config"
	"github.com/hyference/internal/filesystem"
	"github.com/hyference/internal/modules"
	"github.com/rs/zerolog/log"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
	"io/ioutil"
)

var _ Inference = &TensorFlow{}

type TensorFlow struct {
	modelName string
	modelPath string
	client    filesystem.Client
	graph     *tf.Graph
}

func NewTensorFlow(config config.Config, modules modules.Module) *TensorFlow {
	tf := TensorFlow{
		modelPath: config.ModelPath,
		modelName: config.ModelName,
		client:    modules.Client,
	}

	graph, err := tf.initializeModel()
	if err != nil {
		log.Panic().Err(err).Send()
	}
	tf.graph = graph
	return &tf
}

func (t TensorFlow) initializeModel() (*tf.Graph, error) {
	if err := t.client.DownloadModel(t.modelName, t.modelPath); err != nil {
		return nil, err
	}
	model, err := ioutil.ReadFile(t.modelName)
	if err != nil {
		return nil, err
	}

	graph := tf.NewGraph()
	if err := graph.Import(model, ""); err != nil {
		return nil, err
	}
	return graph, nil
}

func (t TensorFlow) Inference(input interface{}) (interface{}, error) {
	tensor, err := tf.NewTensor(input)
	if err != nil {
		return nil, err
	}
	session, err := tf.NewSession(t.graph, nil)
	if err != nil {
		return nil, err
	}
	defer func(session *tf.Session) {
		_ = session.Close()
	}(session)

	s := op.NewScope()
	input = op.Placeholder(s, tf.String)
	output, err := session.Run(
		map[tf.Output]*tf.Tensor{t.graph.Operation("input").Output(0): tensor},
		[]tf.Output{t.graph.Operation("output").Output(0)},
		nil)
	if err != nil {
		return nil, err
	}
	return output[0], nil
}
