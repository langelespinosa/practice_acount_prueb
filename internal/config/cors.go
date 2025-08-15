package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsConfig() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "http://localhost:7000, http://127.0.0.1:7000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	})
}
