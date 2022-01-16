package fasttext_wrapper

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

func TestFoo(t *testing.T) {
	model, _ := New("cooking.model.bin")

	// Label the sentence with that FastText model
	sentence := "Sentence to predict"
	result, err := model.Predict(sentence)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestReadWriteInference(t *testing.T) {
	file, err := ioutil.ReadFile("cooking.model.bin")
	if err != nil {
		return
	}
	_ = ioutil.WriteFile("model.bin", file, fs.ModePerm)
	model, _ := New("model.bin")

	// Label the sentence with that FastText model
	sentence := "Sentence to predict"
	result, err := model.Predict(sentence)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
