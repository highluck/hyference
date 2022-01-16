package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hyference/config"
	"github.com/hyference/https"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type HttpServer struct {
	server *http.Server
}

func NewHttpServer(config config.Config) (*HttpServer, error) {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Timeout(60 * time.Minute))
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)
	router.Use(middleware.SetHeader("content-type", "application/json"))
	router.Middlewares()
	router.Get("/health", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write(https.SuccessWithResult("success", nil).ToByte())
	})
	router.Handle("/metrics", promhttp.Handler())
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Port),
		Handler: router,
	}
	return &HttpServer{
		server: server,
	}, nil
}

func (s *HttpServer) Start() {
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Http HttpServer StartFail")
		panic(err)
	}
	log.Info().Msgf("server start")
}

func (s *HttpServer) Stop() {
	var (
		stop = make(chan os.Signal, 1)
	)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSEGV)
	defer signal.Stop(stop)
	<-stop
	log.Info().Msg("......... Stopping server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	err := s.server.Shutdown(ctx)
	if err != nil {
		return
	}

	log.Info().Msg("exit Server")
}
