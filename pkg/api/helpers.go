package api

import (
	"fmt"
	"log/slog"
	"os"
	"practice_account/pkg/domain"

	"github.com/gofiber/fiber/v2"
)

type Helpers struct {
}

func NewHelpers() *Helpers {
	return &Helpers{}
}

func (h *Helpers) clientError(ctx *fiber.Ctx, status int, err error) error {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Error(err.Error(), "method", ctx.Method(), "uri", ctx.Context().URI(), "codigo", status)

	return ctx.Status(status).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func (h *Helpers) serverError(ctx *fiber.Ctx, status int, err error) error {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Error(err.Error(), "method", ctx.Method(), "uri", ctx.Context().URI(), "codigo", status)

	return ctx.SendStatus(status)
}

func (h *Helpers) dtoUser(u *domain.Users) UserResponse {
	return UserResponse{
		ID:             int(u.ID),
		Userid:         u.Userid,
		Login:          u.Login,
		Email:          u.Email,
		Maildir:        u.Maildir,
		Identificacion: u.Identificacion,
		Grupo:          u.Grupo,
		Dominio:        u.Dominio,
		Quota:          u.Quota,
	}
}

func (h *Helpers) dtoUserWithTrans(u *domain.Users) UserResponseWithTrans {

	var trans *TransportResponse
	if (u.Transport != domain.Transport{}) {
		trans = &TransportResponse{
			ID:        int(u.Transport.ID),
			Domain:    u.Transport.Domain,
			Transport: u.Transport.Transport,
		}
	}

	return UserResponseWithTrans{
		ID:             int(u.ID),
		Userid:         u.Userid,
		Login:          u.Login,
		Email:          u.Email,
		Maildir:        u.Maildir,
		Identificacion: u.Identificacion,
		Grupo:          u.Grupo,
		Dominio:        u.Dominio,
		Quota:          u.Quota,
		Transport:      trans,
	}
}

func (h *Helpers) dtoUsers(users []domain.Users) []UserResponse {

	resultado := make([]UserResponse, len(users))
	for index, value := range users {
		resultado[index] = UserResponse{
			ID:             int(value.ID),
			Userid:         value.Userid,
			Login:          value.Login,
			Email:          value.Email,
			Maildir:        value.Maildir,
			Identificacion: value.Identificacion,
			Grupo:          value.Grupo,
			Dominio:        value.Dominio,
			Quota:          value.Quota,
		}
	}
	return resultado
}

func (h *Helpers) dtoUsersWithTrans(users []domain.Users) []UserResponseWithTrans {

	resultado := make([]UserResponseWithTrans, len(users))
	for index, value := range users {
		var trans *TransportResponse
		if (value.Transport != domain.Transport{}) {
			trans = &TransportResponse{
				ID:        int(value.Transport.ID),
				Domain:    value.Transport.Domain,
				Transport: value.Transport.Transport,
			}
		}

		resultado[index] = UserResponseWithTrans{
			ID:             int(value.ID),
			Userid:         value.Userid,
			Login:          value.Login,
			Email:          value.Email,
			Maildir:        value.Maildir,
			Identificacion: value.Identificacion,
			Grupo:          value.Grupo,
			Dominio:        value.Dominio,
			Quota:          value.Quota,
			Transport:      trans,
		}
	}
	return resultado
}

func (h *Helpers) Swagger() {

	type Login struct {
		Login    string
		Password string
	}

	type CreateUser struct {
		Login          string
		Email          string
		Password       string
		Maildir        string
		Identificacion string
		Grupo          string
		Dominio        int
		Quota          int32
	}

	type UpdateUser struct {
		Login          string
		Email          string
		Maildir        string
		Identificacion string
		Grupo          string
		Dominio        int
		Quota          int32
	}

	type UpdatePass struct {
		Password        string
		NewPassword     string
		ConfirmPassword string
	}

	type CreateAlias struct {
		Local  string
		Remoto string
	}

	type CreateTransport struct {
		Domain    string
		Transport string
	}

	var login Login
	var createUser CreateUser
	var updateUser UpdateUser
	var updatePass UpdatePass
	var createAlias CreateAlias
	var createTransport CreateTransport

	fmt.Println(login, createUser, updateUser, updatePass)
	fmt.Println(createAlias)
	fmt.Println(createTransport)
}

type UserResponse struct {
	ID             int
	Userid         int
	Login          string
	Email          string
	Maildir        string
	Identificacion string
	Grupo          string
	Dominio        int
	Quota          int32
}

type UserResponseWithTrans struct {
	ID             int
	Userid         int
	Login          string
	Email          string
	Maildir        string
	Identificacion string
	Grupo          string
	Dominio        int
	Quota          int32
	Transport      *TransportResponse
}

type TransportResponse struct {
	ID        int
	Domain    string
	Transport string
}
