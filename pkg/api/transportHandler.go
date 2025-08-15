package api

import (
	"errors"
	"practice_account/pkg/domain"
	"practice_account/pkg/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TransportHandler struct {
	service *service.TransportService
	helpers *Helpers
}

func NewTransportHandler(transService *service.TransportService) *TransportHandler {
	return &TransportHandler{service: transService}
}

// @Tags    Transport
// @Accept  json
// @Success 200
// @Failure 500
// @Failure 401
// @Failure 400
// @Param   request body api.Swagger.CreateTransport true "Domain, Transport"
// @Router  /transport [post]
func (h *TransportHandler) CreateTransport(ctx *fiber.Ctx) error {

	var transport domain.Transport

	if err := ctx.BodyParser(&transport); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	validate := validator.New()
	if err := validate.Struct(transport); err != nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, err)
	}

	if err := h.service.ValidateDomain(transport.Domain); err == nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, errors.New("domain already exists"))
	}

	if err := h.service.CreateTransport(&transport); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.SendStatus(200)
}

// @Tags    Transport
// @Produce json
// @Success 200
// @Failure 401
// @Failure 404
// @Param   id path int true "Id"
// @Router  /transport/{id} [get]
func (h *TransportHandler) GetTransport(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	transport, err := h.service.GetTransport(uint(id))
	if err != nil {
		return h.helpers.clientError(ctx, fiber.StatusNotFound, err)
	}

	return ctx.JSON(transport)
}

// @Tags    Transport
// @Produce json
// @Success 200
// @Failure 500
// @Failure 401
// @Router  /transports [get]
func (h *TransportHandler) GetAllTransports(ctx *fiber.Ctx) error {

	transports, err := h.service.GetAllTransport()
	if err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.JSON(transports)
}

// @Tags    Transport
// @Accept  json
// @Success 200
// @Failure 500
// @Failure 404
// @Failure 401
// @Failure 400
// @Param   id path int true "Id"
// @Param   request body api.Swagger.CreateTransport true "Domain, Transport"
// @Router  /transport/{id} [put]
func (h *TransportHandler) UpdateTransport(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	var transport domain.Transport

	transDb, err := h.service.GetTransport(uint(id))
	if err != nil {
		return h.helpers.clientError(ctx, fiber.StatusNotFound, err)
	}

	if err := ctx.BodyParser(&transport); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	validate := validator.New()
	if err = validate.Struct(transport); err != nil {
		return h.helpers.clientError(ctx, fiber.StatusBadRequest, err)
	}

	if transDb.Domain != transport.Domain {
		if err := h.service.ValidateDomain(transport.Domain); err == nil {
			return h.helpers.clientError(ctx, fiber.StatusBadRequest, errors.New("domain already exists"))
		}
	}

	if err := h.service.UpdateTransport(uint(id), &transport); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.SendStatus(200)
}

// @Tags    Transport
// @Success 200
// @Failure 500
// @Failure 404
// @Failure 401
// @Param   id path int true "Id"
// @Router  /transport/{id} [delete]
func (h *TransportHandler) DeleteTransport(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	if _, err := h.service.GetTransport(uint(id)); err != nil {
		return h.helpers.clientError(ctx, fiber.StatusNotFound, err)
	}

	if err := h.service.DeleteTransport(uint(id)); err != nil {
		return h.helpers.serverError(ctx, 500, err)
	}

	return ctx.SendStatus(200)
}
