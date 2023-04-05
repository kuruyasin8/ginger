package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kuruyasin8/ginger/config"
	"github.com/kuruyasin8/ginger/service"
)

type Server struct {
	app     *fiber.App
	service *service.Service
}

var server *Server

func New(app *fiber.App, service *service.Service) *Server {
	if server == nil {
		server = &Server{
			app:     app,
			service: service,
		}
	}

	return server
}

func (s *Server) Listen() error {
	return s.app.Listen(config.Port)
}

func (s *Server) Close() error {
	return s.app.Shutdown()
}
