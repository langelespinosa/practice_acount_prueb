package main

import (
	"log/slog"
	"os"

	"practice_account/internal/config"
	"practice_account/pkg/api"
	"practice_account/pkg/repository"
	"practice_account/pkg/service"

	"github.com/gofiber/fiber/v2"
)

// @title Practice Account Manager
// @host localhost:7000
func main() {

	app := fiber.New()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := config.ConnectDB()
	if err != nil {
		logger.Error(err.Error())
	}

	userRepo := repository.NewUserRepo(db)
	aliasRepo := repository.NewAliasRepo(db)
	transRepo := repository.NewTransportRepo(db)

	userServ := service.NewUserService(userRepo)
	aliasServ := service.NewAliasService(aliasRepo)
	transServ := service.NewTransportService(transRepo)

	userHandler := api.NewUserHandler(userServ)
	aliasHandler := api.NewAliasHandler(aliasServ)
	transHandler := api.NewTransportHandler(transServ)

	app.Use(config.CorsConfig())
	api.Router(app, userHandler, aliasHandler, transHandler)
	app.Listen(":7000")
}
