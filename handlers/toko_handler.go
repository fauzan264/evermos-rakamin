package handlers

import (
	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/services"
	"github.com/gofiber/fiber/v2"
)

type tokoHandler struct {
	tokoService services.TokoService
}

func NewTokoHandler(tokoService services.TokoService) *tokoHandler {
	return &tokoHandler{tokoService}
}

func (h *tokoHandler) MyToko(c *fiber.Ctx) error {
	authUser := c.Locals("authUser")
	user, ok := authUser.(response.UserResponse)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{constants.ErrUnauthorized.Error()},
			Data: nil,
		})
	}

	requestUser := request.GetByUserIDRequest{
		ID: user.ID,
	}

	myTokoResponse, err := h.tokoService.GetMyToko(requestUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: true,
		Message: constants.SuccessGetData,
		Errors: nil,
		Data: myTokoResponse,
	})
}

func (h *tokoHandler) GetListToko(c *fiber.Ctx) error {
	var request request.TokoListRequest

	err := c.QueryParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	if request.Page <= 0 {
		request.Page = 1
	}

	if request.Limit <= 0 {
		request.Limit = 1
	}

	tokoResponse, err := h.tokoService.GetListToko(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: true,
		Message: constants.SuccessGetData,
		Errors: nil,
		Data: tokoResponse,
	})
}

func (h *tokoHandler) GetDetailToko(c *fiber.Ctx) error {
	var requestID request.GetTokoByID

	err := c.ParamsParser(&requestID)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	myTokoResponse, err := h.tokoService.GetTokoByID(requestID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: true,
		Message: constants.SuccessGetData,
		Errors: nil,
		Data: myTokoResponse,
	})
}