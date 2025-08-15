package api

import (
	"errors"
	"practice_account/pkg/domain"
	"practice_account/pkg/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AliasHandler struct {
	service *service.AliasService
	helpers *Helpers
}

func NewAliasHandler(aliasService *service.AliasService) *AliasHandler {
	return &AliasHandler{service: aliasService}
}

// @Tags    Alias
// @Accept  json
// @Success 200
// @Failure 500
// @Failure 401
// @Failure 400
// @Param   request body api.Swagger.CreateAlias true "Local, Remoto"
// @Router  /alias [post]
func (h *AliasHandler) CreateAlias(ctx *fiber.Ctx) error {

	var alias domain.Aliases

	if err := ctx.BodyParser(&alias); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	validate := validator.New()
	if err := validate.Struct(alias); err != nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, err)
	}

	if err := h.service.ValidateLocal(alias.Local); err == nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, errors.New("local already exists"))
	}

	if err := h.service.ValidateRemoto(alias.Remoto); err == nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, errors.New("remoto already exists"))
	}

	if err := h.service.CreateAlias(&alias); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.SendStatus(200)
}

// @Tags    Alias
// @Produce json
// @Success 200
// @Failure 401
// @Failure 404
// @Param   id path int true "Id"
// @Router  /alias/{id} [get]
func (h *AliasHandler) GetAlias(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	alias, err := h.service.GetAlias(uint(id))
	if err != nil {
		return h.helpers.clientError(ctx, fiber.StatusNotFound, err)
	}

	return ctx.JSON(alias)
}

// @Tags    Alias
// @Produce json
// @Success 200
// @Failure 500
// @Failure 401
// @Router  /aliases [get]
func (h *AliasHandler) GetAllAlias(ctx *fiber.Ctx) error {

	alias, err := h.service.GetAllAlias()
	if err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.JSON(alias)
}

// @Tags    Alias
// @Accept  json
// @Success 200
// @Failure 500
// @Failure 404
// @Failure 401
// @Failure 400
// @Param   id path int true "Id"
// @Param   request body api.Swagger.CreateAlias true "Local, Remoto"
// @Router  /alias/{id} [put]
func (h *AliasHandler) UpdateAlias(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	var alias domain.Aliases

	aliasDb, err := h.service.GetAlias(uint(id))
	if err != nil {
		return h.helpers.clientError(ctx, fiber.StatusNotFound, err)
	}

	if err := ctx.BodyParser(&alias); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	validate := validator.New()
	if err = validate.Struct(alias); err != nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, err)
	}

	if aliasDb.Local != alias.Local {
		if err := h.service.ValidateLocal(alias.Local); err == nil {
			return h.helpers.clientError(ctx, fiber.StatusBadRequest, errors.New("local already exists"))
		}
	}

	if aliasDb.Remoto != alias.Remoto {
		if err := h.service.ValidateRemoto(alias.Remoto); err == nil {
			return h.helpers.clientError(ctx, fiber.StatusBadRequest, errors.New("remoto already exists"))
		}
	}

	if err := h.service.UpdateAlias(uint(id), &alias); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.SendStatus(200)
}

// @Tags    Alias
// @Success 200
// @Failure 500
// @Failure 404
// @Failure 401
// @Param   id path int true "Id"
// @Router  /alias/{id} [delete]
func (h *AliasHandler) DeleteAlias(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	if _, err := h.service.GetAlias(uint(id)); err != nil {
		return h.helpers.clientError(ctx, fiber.StatusNotFound, err)
	}

	if err := h.service.DeleteAlias(uint(id)); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.SendStatus(200)
}
