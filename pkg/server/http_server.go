package server

import (
	"first_project/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewHttpServer(httpConf config.HTTP, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           httpConf.Host + ":" + httpConf.Port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

// Start - Запуск сервера
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
