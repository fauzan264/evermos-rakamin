package handlers

import (
	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/services"
	"github.com/gofiber/fiber/v2"
)

type provinceCityHandler struct {
	provinceCityService services.ProvinceCityService
}

func NewProvinceCityHandler(provinceCityService services.ProvinceCityService) *provinceCityHandler {
	return &provinceCityHandler{provinceCityService}
}

func (h *provinceCityHandler) GetListProvince(c *fiber.Ctx) error {
	getListProvince, err := h.provinceCityService.GetListProvince()
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
		Data: getListProvince,
	})
}

func (h *provinceCityHandler) GetDetailProvince(c *fiber.Ctx) error {
	var request request.GetByProvinceIDRequest

	err := c.ParamsParser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	getProvince, err := h.provinceCityService.GetDetailProvince(request)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.Response{
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
		Data: getProvince,
	})
}

func (h *provinceCityHandler) GetListCity(c *fiber.Ctx) error {
	var request request.GetByProvinceIDRequest

	err := c.ParamsParser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	getListCity, err := h.provinceCityService.GetListCity(request)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.Response{
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
		Data: getListCity,
	})
}

func (h *provinceCityHandler) GetDetailCity(c *fiber.Ctx) error {
	var request request.GetByCityIDRequest

	err := c.ParamsParser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	getCity, err := h.provinceCityService.GetDetailCity(request)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.Response{
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
		Data: getCity,
	})
}
