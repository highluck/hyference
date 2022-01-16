package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/hyference/errors"
	"github.com/hyference/internal/config"
	"github.com/hyference/internal/https"
	"github.com/hyference/internal/inference"
	"net/http"
)

func Route(router *chi.Mux, config config.Config, inference inference.Inference) {
	uri := config.Uri
	parameterType := config.ParameterType
	router.Route(uri, func(r chi.Router) {
		r.Post("", func(writer http.ResponseWriter, request *http.Request) {
			req := Request{}
			if err := https.DecodeFromQueryRequestBody(request, &req); err != nil {
				_, _ = writer.Write(https.Fail(err).ToByte())
				return
			}
			if err := verifyRequest(parameterType, req); err != nil {
				_, _ = writer.Write(https.Fail(err).ToByte())
				return
			}

			_, _ = writer.Write(https.SuccessWithResult(inference.Inference(req.Input)).ToByte())
		})
		r.Get("", func(writer http.ResponseWriter, request *http.Request) {
			req := Request{}
			if err := https.DecodeFromQueryString(request, &req); err != nil {
				_, _ = writer.Write(https.Fail(err).ToByte())
				return
			}
			if err := verifyRequest(parameterType, req); err != nil {
				_, _ = writer.Write(https.Fail(err).ToByte())
				return
			}

			_, _ = writer.Write(https.SuccessWithResult(inference.Inference(req.Input)).ToByte())
		})
		r.Post("/bulk", func(writer http.ResponseWriter, request *http.Request) {
			req := RequestBulk{}
			if err := https.DecodeFromQueryRequestBody(request, &req); err != nil {
				_, _ = writer.Write(https.Fail(err).ToByte())
				return
			}

			if len(req.Inputs) == 0 {
				_, _ = writer.Write(https.SuccessResponse(nil).ToByte())
				return
			}

			results := make([]interface{}, 0)
			for _, input := range req.Inputs {
				if err := verifyRequest(parameterType, input); err != nil {
					_, _ = writer.Write(https.Fail(err).ToByte())
					return
				}
				i, err := inference.Inference(input)
				if err != nil {
					_, _ = writer.Write(https.Fail(err).ToByte())
					return
				}
				results = append(results, i)
			}
			_, _ = writer.Write(https.SuccessWithResult(results, nil).ToByte())
		})
	})
}

func verifyRequest(parameterType config.ParameterType, input interface{}) error {
	switch parameterType {
	case config.StringType:
		if _, ok := input.(string); ok == false {
			return errors.New("request Input TypeCast error").
				WithStatus(errors.BadRequest)
		}
	case config.IntType:
		if _, ok := input.(int); ok == false {
			return errors.New("request Input TypeCast error").
				WithStatus(errors.BadRequest)
		}
	case config.LongType:
		if _, ok := input.(int64); ok == false {
			return errors.New("request Input TypeCast error").
				WithStatus(errors.BadRequest)
		}
	case config.FloatType:
		if _, ok := input.(float64); ok == false {
			return errors.New("request Input TypeCast error").
				WithStatus(errors.BadRequest)
		}
	case config.StructType:
		if _, ok := input.(map[interface{}]interface{}); ok == false {
			return errors.New("request Input TypeCast error").
				WithStatus(errors.BadRequest)
		}
	default:
		return errors.New("request Input not supported type error").
			WithStatus(errors.BadRequest)
	}
	return nil
}

type Request struct {
	Input interface{} `json:"input"`
}

type RequestBulk struct {
	Inputs []interface{} `json:"inputs"`
}
