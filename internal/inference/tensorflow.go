package inference
var _ Inference = &TensorFlow{}

type TensorFlow struct {

}

func (t TensorFlow) Inference(input interface{}) (interface{}, error) {
	panic("implement me")
}
