package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kuruyasin8/ginger/repository"
	"github.com/kuruyasin8/ginger/server"
	"github.com/kuruyasin8/ginger/service"
)

func main() {
	ctx := context.Background()

	repo := repository.New(ctx)
	if err := repo.Connect(ctx); err != nil {
		panic(err)
	}
	defer repo.Close(ctx)

	app := fiber.New()

	app.Use(logger.New())

	service := service.New(ctx, repo)
	server := server.New(app, service)

	server.Register(ctx)
	server.GetMultipleUsers(ctx)
	server.GetSingleUser(ctx)

	log.Fatal(server.Listen())
}
