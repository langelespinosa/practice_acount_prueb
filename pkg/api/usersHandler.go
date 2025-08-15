package api

import (
	"errors"
	"practice_account/pkg/domain"
	"practice_account/pkg/service"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UsersHandler struct {
	service *service.UserService
	helpers *Helpers
}

func NewUserHandler(userService *service.UserService) *UsersHandler {
	return &UsersHandler{service: userService}
}

// @Tags    User
// @Accept  json
// @Success 200
// @Failure 500
// @Failure 404
// @Failure 401
// @Param   request body api.Swagger.Login true "Login, Password"
// @Router  /login [post]
func (h *UsersHandler) Login(ctx *fiber.Ctx) error {

	var user domain.Users

	if err := ctx.BodyParser(&user); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, err)
	}

	tokenL, tokenA, err := h.service.ValidateLogin(user.Login, user.Password)
	if err != nil {
		return h.helpers.clientError(ctx, fiber.StatusUnauthorized, err)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh",
		Value:    tokenL,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "access",
		Value:    tokenA,
		Expires:  time.Now().Add(time.Minute * 30),
		HTTPOnly: true,
	})

	return ctx.SendStatus(200)
}

// @Tags    User
// @Accept  json
// @Success 200
// @Failure 500
// @Failure 401
// @Failure 400
// @Param   request body api.Swagger.CreateUser true "Login, Email, Password, Maildir, Identificacion, Grupo, Dominio, Quota"
// @Router  /register [post]
func (h *UsersHandler) CreateUser(ctx *fiber.Ctx) error {

	var user domain.Users

	if err := ctx.BodyParser(&user); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, err)
	}

	if err := h.service.ValidateEmail(user.Email); err == nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, errors.New("email already exists"))
	}

	hashPass, err := h.service.Encrypt(user.Password)
	if err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	user.Password = string(hashPass)
	user.Created = time.Now()

	if err := h.service.CreateUser(&user); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.SendStatus(200)
}

// @Tags    User
// @Produce json
// @Success 200
// @Failure 401
// @Failure 404
// @Param   id path int true "Id"
// @Router  /user/{id} [get]
func (h *UsersHandler) GetUser(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	user, err := h.service.GetUserWithTrans(uint(id))
	if err != nil {
		return h.helpers.clientError(ctx, fiber.StatusNotFound, err)
	}

	return ctx.JSON(h.helpers.dtoUserWithTrans(user))
	//return ctx.JSON(user)
}

// @Tags    User
// @Produce json
// @Success 200
// @Failure 500
// @Failure 401
// @Router  /users [get]
func (h *UsersHandler) GetAllUsers(ctx *fiber.Ctx) error {

	users, err := h.service.GetAllUsersWithTrans()
	if err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.JSON(h.helpers.dtoUsersWithTrans(users))
}

// @Tags    User
// @Accept  json
// @Success 200
// @Failure 500
// @Failure 404
// @Failure 401
// @Failure 400
// @Param   id path int true "Id"
// @Param   request body api.Swagger.UpdateUser true "Login, Email, Maildir, Identificacion, Grupo, Dominio, Quota"
// @Router  /user/{id} [put]
func (h *UsersHandler) UpdateUser(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	var user domain.Users

	userDb, err := h.service.GetUser(uint(id))
	if err != nil {
		return h.helpers.clientError(ctx, fiber.StatusNotFound, err)
	}

	if err := ctx.BodyParser(&user); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	validate := validator.New()
	if err = validate.Struct(user); err != nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, err)
	}

	if userDb.Email != user.Email {
		if err := h.service.ValidateEmail(user.Email); err == nil {
			return h.helpers.clientError(ctx, fiber.StatusBadRequest, errors.New("email already exists"))
		}
	}

	if err := h.service.UpdateUser(uint(id), &user); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.SendStatus(200)
}

// @Tags    User
// @Accept  json
// @Success 200
// @Failure 500
// @Failure 404
// @Failure 401
// @Failure 400
// @Param   id path int true "Id"
// @Param   request body api.Swagger.UpdatePass true "Password, NewPassword, ConfirmPassword"
// @Router  /user/pass/{id} [put]
func (h *UsersHandler) UpdatePass(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	type Password struct {
		Password        string `validate:"min=8,containsany=!¡@#$%&=¿?*-_."`
		NewPassword     string `validate:"min=8,containsany=!¡@#$%&=¿?*-_."`
		ConfirmPassword string `validate:"min=8,eqfield=NewPassword,containsany=!¡@#$%&=¿?*-_."`
	}

	var pass Password

	user, err := h.service.GetUser(uint(id))
	if err != nil {
		return h.helpers.clientError(ctx, fiber.StatusNotFound, err)
	}

	hashPass := user.Password

	if err := ctx.BodyParser(&pass); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	validate := validator.New()
	if err = validate.Struct(pass); err != nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, err)
	}

	if err = h.service.ValidatePass(hashPass, pass.Password); err != nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, err)
	} else {

		hashPassNew, err := h.service.Encrypt(user.Password)
		if err != nil {
			return h.helpers.serverError(ctx, 500, err)
		}

		user.Password = string(hashPassNew)
		pass.NewPassword = ""
		pass.ConfirmPassword = ""
	}

	if err := h.service.UpdateUser(uint(id), user); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.SendStatus(200)
}

// @Tags    User
// @Success 200
// @Failure 500
// @Failure 404
// @Failure 401
// @Param   id path int true "Id"
// @Router  /user/{id} [delete]
func (h *UsersHandler) DeleteUser(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	if _, err := h.service.GetUser(uint(id)); err != nil {
		return h.helpers.clientError(ctx, fiber.StatusNotFound, err)
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.SendStatus(200)
}

// @Tags    User
// @Success 200
// @Failure 401
// @Router  /user/logout [post]
func (h *UsersHandler) Logout(ctx *fiber.Ctx) error {

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24),
		HTTPOnly: true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "access",
		Value:    "",
		Expires:  time.Now().Add(-time.Minute * 30),
		HTTPOnly: true,
	})

	return ctx.SendStatus(200)
}

func (h *UsersHandler) Authentication(ctx *fiber.Ctx) error {

	if err := h.service.AuthRequired(ctx); err != nil {
		return ctx.SendStatus(401)
	}

	return ctx.Next()
}
