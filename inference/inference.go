package inference

type Inference interface {
	Inference(input interface{}) (interface{}, error)
}
