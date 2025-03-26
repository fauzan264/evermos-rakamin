package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/services"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) GetDetailProduct(c *fiber.Ctx) error {
	var requestID request.GetByProductIDRequest

	err := c.ParamsParser(&requestID)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Status: false,
			Message: constants.FailedGetData,
			Errors: []string{err.Error()},
			Data: nil,
		})
	}

	myTokoResponse, err := h.productService.GetProductByID(requestID)
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

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	authUser := c.Locals("authUser")
	if authUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: false,
			Message: constants.FailedInsertData,
			Errors: []string{constants.ErrUnauthorized.Error()},
			Data: nil,
		})
	}

	user, ok := authUser.(response.UserResponse)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: false,
			Message: constants.FailedInsertData,
			Errors: []string{constants.ErrUnauthorized.Error()},
			Data: nil,
		})
	}

	requestUser := request.GetByUserIDRequest{ID: user.ID}

	// Parsing form data
	var requestData request.ProductRequest
	requestData.NamaProduk = c.FormValue("nama_produk")

	idCategory, err := strconv.Atoi(c.FormValue("category_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: "Invalid category_id",
			Errors: []string{err.Error()},
			Data: nil,
		})
	}
	requestData.IDCategory = idCategory

	requestData.HargaReseller = c.FormValue("harga_reseller")
	requestData.HargaKonsumen = c.FormValue("harga_konsumen")

	stok, err := strconv.Atoi(c.FormValue("stok"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: "Invalid stok",
			Errors: []string{err.Error()},
			Data: nil,
		})
	}
	requestData.Stok = stok

	requestData.Deskripsi = c.FormValue("deskripsi")

	// Handle file upload
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: false,
			Message: constants.FailedInsertData,
			Errors: []string{"Invalid form-data"},
			Data: nil,
		})
	}

	files := form.File["photos"]
	var photos []request.PhotoProductRequest

	// Pastikan folder "uploads" ada
	uploadDir := "uploads/products"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	for _, file := range files {
		ext := filepath.Ext(file.Filename)
		fileName := fmt.Sprintf("%s%s", uuid.NewString(), ext)
		filePath := filepath.Join(uploadDir, fileName)

		err := c.SaveFile(file, filePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
				Status: false,
				Message: constants.FailedInsertData,
				Errors: []string{"Failed to save file"},
				Data: nil,
			})
		}

		photos = append(photos, request.PhotoProductRequest{URL: filePath})
	}

	requestData.Photos = photos

	productResponse, err := h.productService.CreateProduct(requestUser, requestData)
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
		Data: productResponse,
	})
}