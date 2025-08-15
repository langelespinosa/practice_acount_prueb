package api

import (
	_ "practice_account/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Router(app *fiber.App, userH *UsersHandler, aliasH *AliasHandler, transH *TransportHandler) {

	//http://localhost:7000/swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	//http://localhost:7000/user
	app.Post("/login", userH.Login)
	app.Post("/register", userH.CreateUser)
	app.Use("/user", userH.Authentication)
	app.Post("/user/logout", userH.Logout)
	app.Get("/user/:id", userH.GetUser)
	app.Get("/users", userH.GetAllUsers)
	app.Put("/user/:id", userH.UpdateUser)
	app.Put("/user/pass/:id", userH.UpdatePass)
	app.Delete("/user/:id", userH.DeleteUser)

	//http://localhost:7000/alias
	app.Use("/alia", userH.Authentication)
	app.Post("/alias", aliasH.CreateAlias)
	app.Get("/alias/:id", aliasH.GetAlias)
	app.Get("/aliases", aliasH.GetAllAlias)
	app.Put("/alias/:id", aliasH.UpdateAlias)
	app.Delete("/alias/:id", aliasH.DeleteAlias)

	//http://localhost:7000/transport
	app.Use("/transpor", userH.Authentication)
	app.Post("/transport", transH.CreateTransport)
	app.Get("/transport/:id", transH.GetTransport)
	app.Get("/transports", transH.GetAllTransports)
	app.Put("/transport/:id", transH.UpdateTransport)
	app.Delete("transport/:id", transH.DeleteTransport)

	//http://localhost:7000/buscar?query=
	app.Get("/login/buscar", SemanticSearch)
}
