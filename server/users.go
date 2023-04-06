package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kuruyasin8/ginger/model"
	"github.com/kuruyasin8/ginger/service"
)

func (s *Server) Register(ctx context.Context) *Server {
	s.app.Post("/register", func(c *fiber.Ctx) error {
		user := new(model.User)

		if err := c.BodyParser(user); err != nil {
			return err
		}

		if err := s.service.Register(ctx, user); err != nil {
			return err
		}

		return c.Status(http.StatusCreated).SendString("registration successfully completed")
	})

	return s
}

func (s *Server) Login(ctx context.Context) *Server {
	s.app.Post("/login", func(c *fiber.Ctx) error {
		user := new(model.User)

		if err := c.BodyParser(user); err != nil {
			return err
		}

		res, err := s.service.Login(ctx, user)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	return s
}

func (s *Server) RefreshToken(ctx context.Context) *Server {
	s.app.Post("/refresh", func(c *fiber.Ctx) error {
		token := new(model.Token)

		if err := c.BodyParser(token); err != nil {
			return err
		}

		res, err := s.service.RefreshToken(ctx, token)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	return s
}

func (s *Server) GetSingleUser(ctx context.Context) *Server {
	s.app.Get("/users/:uid", func(c *fiber.Ctx) error {
		query := new(service.UserQuery)

		if uid, err := strconv.Atoi(c.Params("uid")); err != nil {
			return err
		} else {
			query.ID = uint(uid)
		}

		user, err := s.service.GetSingleUser(ctx, query)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(user)
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
			return err
		}

		return c.Status(http.StatusOK).JSON(users)
	})

	return s
}
