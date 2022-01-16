package https

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/hyference/errors"
	"net/http"
)

var decoder = schema.NewDecoder()

func DecodeFromQueryString(request *http.Request, dataModel interface{}) error {
	op := errors.GetMethodName()
	if err := decoder.Decode(dataModel, request.URL.Query()); err != nil {
		return errors.Wrapper(err, op+"QueryString")
	}
	return nil
}

func DecodeFromQueryRequestBody(request *http.Request, dataModel interface{}) error {
	op := errors.GetMethodName()
	if err := json.NewDecoder(request.Body).Decode(&dataModel); err != nil {
		return errors.Wrapper(err, op+"body")
	}
	return nil
}
