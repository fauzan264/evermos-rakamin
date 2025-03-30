package handlers

import (
	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/helpers"
	"github.com/fauzan264/evermos-rakamin/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type trxHandler struct {
	trxService services.TRXService
}

func NewTRXHandler(trxService services.TRXService) *trxHandler {
	return &trxHandler{trxService}
}

func (h *trxHandler) GetListTRX(c *fiber.Ctx) error {
	var requestUser request.GetByUserIDRequest
	var requestData request.TRXListRequest

	authUser := c.Locals("authUser")
	if authUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: false,
			Message: constants.Unauthorized,
			Errors: []string{constants.ErrUnauthorized.Error()},
			Data: nil,
		})
	}

	user, ok := authUser.(response.UserResponse)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: false,
			Message: constants.Unauthorized,
			Errors: []string{constants.ErrUnauthorized.Error()},
			Data: nil,
		})
	}

	requestUser.ID = user.ID

	err := c.QueryParser(&requestData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}
	
	if requestData.Page <= 0 {
		requestData.Page = 1
	}

	if requestData.Limit <= 0 {
		requestData.Limit = 10
	}

	trxResponse, err := h.trxService.GetListTRX(requestUser, requestData)
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
		Data: trxResponse,
	})
}

func (h *trxHandler) GetDetailTRX(c *fiber.Ctx) error {
	var requestUser request.GetByUserIDRequest
	var requestID request.GetByTRXIDRequest

	authUser := c.Locals("authUser")
	user, ok := authUser.(response.UserResponse)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: false,
			Message: constants.Unauthorized,
			Errors: []string{constants.ErrUnauthorized.Error()},
			Data: nil,
		})
	}

	requestUser.ID = user.ID

	err := c.ParamsParser(&requestID)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	myTRXResponse, err := h.trxService.GetDetailTRX(requestUser, requestID)
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
		Data: myTRXResponse,
	})
}

func (h *trxHandler) CreateTRX(c *fiber.Ctx) error {
	var requestUser request.GetByUserIDRequest
	var requestData request.CreateTrxRequest

	authUser := c.Locals("authUser")
	user, ok := authUser.(response.UserResponse)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: false,
			Message: constants.Unauthorized,
			Errors: []string{constants.ErrUnauthorized.Error()},
			Data: nil,
		})
	}
	
	requestUser.ID = user.ID
	err := c.BodyParser(&requestData)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedInsertData,
			Errors: helpers.FormatValidationError(err),
			Data: nil,
		})
	}

	validate := validator.New()
	err = validate.Struct(requestData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedInsertData,
			Errors: helpers.FormatValidationError(err),
			Data: nil,
		})
	}

	trxResponse, err := h.trxService.CreateTRX(requestUser, requestData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedInsertData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: true,
		Message: constants.SuccessInsertData,
		Errors: nil,
		Data: trxResponse,
	})
}