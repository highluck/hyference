package https

import (
	"github.com/hyference/errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
)

type Result struct {
	test string
}

type Response struct {
	Code    errors.ResponseStatus `json:"code"`
	Message string                `json:"message"`
	Result  interface{}           `json:"result"`
}

func SuccessWithResult(result interface{}, err error) Response {
	if err != nil {
		return Fail(err)
	}

	return Response{
		Result: result,
		Code:   errors.Success,
	}
}

func SuccessResponse(err error) Response {
	if err != nil {
		return Fail(err)
	}
	return Response{
		Code: errors.Success,
	}
}

func Fail(err error) Response {
	log.Err(err).Send()
	return Response{
		Code:    errors.ParseHttpStatus(err),
		Message: err.Error(),
	}
}

func (response Response) ToByte() []byte {
	marshal, err := jsoniter.Marshal(&response)
	if err != nil {
		log.Err(err).Msg("response marshal error")
	}
	return marshal
}
