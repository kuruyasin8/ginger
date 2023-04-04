package server

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kuruyasin8/ginger/service"
)

func (s *Server) GetSingleUser(ctx context.Context) *Server {
	s.app.Get("/users/:uid", func(c *fiber.Ctx) error {
		query := new(service.UserQuery)

		query.ID = c.Params("uid")

		user, err := s.service.GetSingleUser(ctx, query)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "success",
			"user":    user,
		})
	})

	return s
}

func (s *Server) GetMultipleUsers(ctx context.Context) *Server {
	s.app.Get("/users", func(c *fiber.Ctx) error {
		query := new(service.UserQuery)

		query.Page = uint(c.QueryInt("page", 1))
		query.PerPage = uint(c.QueryInt("per_page", 10))
		query.Filter = c.Query("filter", "")

		users, err := s.service.GetMultipleUsers(ctx, query)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "success",
			"users":   users,
		})
	})

	return s
}
