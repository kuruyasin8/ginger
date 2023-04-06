package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kuruyasin8/ginger/service"
	"github.com/kuruyasin8/ginger/stash"
)

func (s *Server) GetSingleUser(ctx context.Context) *Server {
	s.app.Get("/users/:uid", stash.Authenticate(), stash.Auhtorize(stash.Salt, stash.Admin), func(c *fiber.Ctx) error {
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
