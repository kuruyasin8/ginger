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

func New(app *fiber.App, service *service.Service) *Server {
	return &Server{
		app,
		service,
	}
}

func (s *Server) Listen() error {
	return s.app.Listen(config.Port)
}
