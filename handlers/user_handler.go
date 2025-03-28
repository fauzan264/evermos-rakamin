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

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetMyProfile(c *fiber.Ctx) error {
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

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: true,
		Message: constants.SuccessGetData,
		Errors: nil,
		Data: user,
	})
}

func (h *userHandler) UpdateProfile(c *fiber.Ctx) error {
	var requestID request.GetByUserIDRequest
	var requestData request.UpdateProfileRequest

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

	requestID.ID = user.ID

	err := c.BodyParser(&requestData)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedUpdateData,
			Errors: helpers.FormatValidationError(err),
			Data: nil,
		})
	}

	validate := validator.New()
	err = validate.Struct(requestData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedUpdateData,
			Errors: helpers.FormatValidationError(err),
			Data: nil,
		})
	}

	userResponse, err := h.userService.UpdateUser(requestID, requestData)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedUpdateData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: true,
		Message: constants.SuccessUpdateData,
		Errors: nil,
		Data: userResponse,
	})
}

func (h *userHandler) GetMyAddress(c *fiber.Ctx) error {
	var requestUser request.GetByUserIDRequest

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

	requestUser.ID = user.ID

	myAddressResponse, err := h.userService.GetMyAddress(requestUser)
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
		Data: myAddressResponse,
	})
}

func (h *userHandler) GetDetailAddress(c *fiber.Ctx) error {
	var requestUser request.GetByUserIDRequest
	var requestID request.GetByAddressIDRequest

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

	myAddressResponse, err := h.userService.GetAddressUserByID(requestUser, requestID)
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
		Data: myAddressResponse,
	})
}

func (h *userHandler) CreateAddressUser(c *fiber.Ctx) error {
	var requestUser request.GetByUserIDRequest
	var requestData request.CreateAddressRequest

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
	
	addressResponse, err := h.userService.CreateAddressUser(requestUser, requestData)
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
		Data: addressResponse,
	})
}

func (h *userHandler) UpdateAddressUser(c *fiber.Ctx) error {
	var requestUser request.GetByUserIDRequest
	var requestID request.GetByAddressIDRequest
	var requestData request.UpdateAddressRequest

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

	requestUser.ID = user.ID

	err := c.ParamsParser(&requestID)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedUpdateData,
			Errors: helpers.FormatValidationError(err),
		})
	}

	err = c.BodyParser(&requestData)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedUpdateData,
			Errors: helpers.FormatValidationError(err),
			Data: nil,
		})
	}

	validate := validator.New()
	err = validate.Struct(requestData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedUpdateData,
			Errors: helpers.FormatValidationError(err),
			Data: nil,
		})
	}

	addressResponse, err := h.userService.UpdateAddressUser(requestUser, requestID, requestData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedUpdateData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: true,
		Message: constants.SuccessUpdateData,
		Errors: nil,
		Data: addressResponse,
	})
	
}

func (h *userHandler) DeleteAddressUser(c *fiber.Ctx) error {
	var requestUser request.GetByUserIDRequest
	var requestID request.GetByAddressIDRequest

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

	requestUser.ID = user.ID

	err := c.ParamsParser(&requestID)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedDeleteData,
			Errors: helpers.FormatValidationError(err),
			Data: nil,
		})
	}

	err = h.userService.DeleteAddressUser(requestUser, requestID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedDeleteData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: true,
		Message: constants.SuccessDeleteData,
		Errors: nil,
		Data: true,
	})
}
