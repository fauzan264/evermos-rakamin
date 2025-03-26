package handlers

import (
	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/helpers"
	"github.com/fauzan264/evermos-rakamin/services"
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
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
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
	authUser := c.Locals("authUser")
	user, ok := authUser.(response.UserResponse)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: false,
			Message: constants.FailedUpdateData,
			Errors: []string{constants.ErrUnauthorized.Error()},
			Data: nil,
		})
	}

	requestID := request.GetByUserIDRequest{
		ID: user.ID,
	}

	var requestData request.UpdateProfileRequest

	err := c.BodyParser(&requestData)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
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