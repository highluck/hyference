package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/hyference/internal/config"
	"github.com/hyference/internal/https"
	"github.com/hyference/internal/inference"
	"net/http"
)

func Route(router *chi.Mux, config config.Config, inference inference.Inference) {
	uri := config.Uri
	router.Route(uri, func(r chi.Router) {
		r.Post("", func(writer http.ResponseWriter, request *http.Request) {
			req := Request{}
			if err := https.DecodeFromQueryRequestBody(request, &req); err != nil {
				_, _ = writer.Write(https.Fail(err).ToByte())
				return
			}
			_, _ = writer.Write(https.SuccessWithResult(inference.Inference(req.Input)).ToByte())
		})
		r.Get("", func(writer http.ResponseWriter, request *http.Request) {
			req := Request{}
			if err := https.DecodeFromQueryRequestBody(request, &req); err != nil {
				_, _ = writer.Write(https.Fail(err).ToByte())
				return
			}
			_, _ = writer.Write(https.SuccessWithResult(inference.Inference(req.Input)).ToByte())
		})
	})
}

type Request struct {
	Input interface{} `json:"input"`
}
