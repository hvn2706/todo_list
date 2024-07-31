package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"todo_list/config"
	"todo_list/logger"
	"todo_list/server/task"
)

type Server struct {
	Router *chi.Mux
}

type Option func(s *Server)

func NewServer(opts ...Option) *Server {
	s := &Server{
		Router: chi.NewRouter(),
	}
	for _, opt := range opts {
		opt(s)
	}
	s.initRoutes()
	return s
}

func (s *Server) initRoutes() {
	// 1. Health API
	healthRouter := chi.NewRouter()
	s.Router.Mount("/health", healthRouter)
	healthRouter.Get("/ready", ready)

	// 2. Public API
	apiV1Router := chi.NewRouter()
	apiV1Router.Use()
	s.Router.Mount("/api/v1", apiV1Router)

	taskService := task.InitServer()
	apiV1Router.Post("/task", taskService.UpsertTaskAPI())
	apiV1Router.Get("/task", taskService.GetListTasksAPI())
}

func (s *Server) Serve(cfg config.ServerListen) error {
	logger.Infof("Listening on port %v", cfg.Port)
	address := fmt.Sprintf(":%v", cfg.Port)
	return http.ListenAndServe(address, s.Router)
}
