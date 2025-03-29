package handlers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/helpers"
	"github.com/fauzan264/evermos-rakamin/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type shopHandler struct {
	shopService services.ShopService
}

func NewShopHandler(shopService services.ShopService) *shopHandler {
	return &shopHandler{shopService}
}

func (h *shopHandler) MyShop(c *fiber.Ctx) error {
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

	myShopResponse, err := h.shopService.GetMyShop(requestUser)
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
		Data: myShopResponse,
	})
}

func (h *shopHandler) GetListShop(c *fiber.Ctx) error {
	var request request.ShopListRequest

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
		request.Limit = 10
	}

	shopResponse, err := h.shopService.GetListShop(request)
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
		Data: shopResponse,
	})
}

func (h *shopHandler) GetDetailShop(c *fiber.Ctx) error {
	var requestID request.GetShopByID

	err := c.ParamsParser(&requestID)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	myShopResponse, err := h.shopService.GetShopByID(requestID)
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
		Data: myShopResponse,
	})
}

func ( h *shopHandler) UpdateProfileShop(c *fiber.Ctx) error {
	var requestUser request.GetByUserIDRequest
	var requestID request.GetShopByID
	var requestData request.UpdateProfileShopRequest

	authUser := c.Locals("authUser")
	user, ok := authUser.(response.UserResponse)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: false,
			Message: constants.FailedUpdateData,
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
			Data: nil,
		})
	}

	// Parsing form-data
	requestData.Nama = c.FormValue("nama_toko")

	// Handle file upload
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status:  false,
			Message: constants.FailedUpdateData,
			Errors:  []string{"Invalid form-data"},
			Data:    nil,
		})
	}

	var photoPath string
	files := form.File["photo"]
	if len(files) > 0 {
		file := files[0]
		ext := filepath.Ext(file.Filename)
		fileName := fmt.Sprintf("uploads/shop_images/%s%s", uuid.NewString(), ext)

		err := os.MkdirAll("uploads/shop_images/", os.ModePerm) // Buat folder jika belum ada
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
				Status:  false,
				Message: constants.FailedUpdateData,
				Errors:  []string{"Failed to create upload directory"},
				Data:    nil,
			})
		}

		err = c.SaveFile(file, fileName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
				Status:  false,
				Message: constants.FailedUpdateData,
				Errors:  []string{"Failed to save file"},
				Data:    nil,
			})
		}

		photoPath = fileName
	}

	requestData.Photo = photoPath

	productResponse, err := h.shopService.UpdateShop(requestUser, requestID, requestData)
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
		Data: productResponse,
	})
}