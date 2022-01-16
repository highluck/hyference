package https

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"testing"
)

type TestResult struct {
	Blahblah string
}

func Test_responseToByte(t *testing.T) {
	result := TestResult{
		Blahblah: "goood",
	}

	response := SuccessWithResult(result, nil)

	toByte := response.ToByte()
	println("Xxxx")
	marshalString := string(toByte)
	println(marshalString)

	if strings.Contains(marshalString, "200") == false {
		t.Fatal("marshal error")
	}
	println("====================")
	println(string(toByte))
	response2 := &Response{}
	jsoniter.Unmarshal(toByte, response2)

	println("====================")
	println(fmt.Sprintf("code : %d", response2.Code))
	println(fmt.Sprintf("message : %s", response2.Message))
	println(fmt.Sprintf("result : %v", response2.Result))
}
