package stash

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kuruyasin8/ginger/errors"
	"golang.org/x/exp/slices"
)

type Role string

const (
	Admin  Role = "admin"
	Salt   Role = "salt"
	Pepper Role = "pepper"
	Soy    Role = "soy"
)

func Auhtorize(roles ...Role) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		pocket := c.Locals("role").(Role)

		safe := true

		if !slices.Contains(roles, pocket) {
			safe = false
		}

		if !safe {
			err := errors.NewForbidden("You are not allowed to access this resource")
			return c.Status(http.StatusForbidden).SendString(err.Error())
		}

		return c.Next()
	}
}

func Authenticate() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// get the token from the header
		token := "secret"

		// if the token is not present, return an error
		if token == "" {
			err := errors.NewUnauthorized("Missing token")
			c.Status(http.StatusUnauthorized).SendString(err.Error())
			return nil
		}

		// validate the token
		var role Role = "salt"
		var err error = nil
		if err != nil {
			return c.Status(http.StatusUnauthorized).SendString(err.Error())
		}

		// set the claims in the context
		c.Locals("role", role)

		return c.Next()
	}
}
