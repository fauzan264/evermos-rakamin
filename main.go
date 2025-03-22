package main

import (
	"log"
	"os"

	"github.com/fauzan264/evermos-rakamin/config"
	"github.com/fauzan264/evermos-rakamin/handlers"
	"github.com/fauzan264/evermos-rakamin/middleware"
	"github.com/fauzan264/evermos-rakamin/repositories"
	"github.com/fauzan264/evermos-rakamin/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db := config.InitDatabase()

	router := fiber.New()
	router.Use(cors.New())

	// API external for data province & city
	provinceCityApiURL := os.Getenv("API_LOCATION")

	// repositories
	provinceCityRepository := repositories.NewProvinceCityRepository(provinceCityApiURL)
	userRepository := repositories.NewUserRepository(db)
	tokoRepository := repositories.NewTokoRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)

	// services
	jwtService := middleware.NewJWTService()
	authService := services.NewAuthService(jwtService, userRepository, tokoRepository, provinceCityRepository)
	provinceCityService := services.NewProvinceCityService(provinceCityRepository)
	categoryService := services.NewCategoryService(categoryRepository)
	tokoService := services.NewTokoService(tokoRepository)

	// handleres
	authHandler := handlers.NewAuthHandler(authService)
	provinceCityHandler := handlers.NewProvinceCityHandler(provinceCityService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	tokoHandler := handlers.NewTokoHandler(tokoService)

	api := router.Group("/api/v1")
	// auth
	api.Post("/auth/register", authHandler.RegisterUser)
	api.Post("/auth/login", authHandler.LoginUser)

	// province city
	api.Get("/provcity/listprovincies", provinceCityHandler.GetListProvince)
	api.Get("/provcity/detailprovince/:prov_id", provinceCityHandler.GetDetailProvince)
	api.Get("/provcity/listcities/:prov_id", provinceCityHandler.GetListCity)
	api.Get("/provcity/detailcity/:city_id", provinceCityHandler.GetDetailCity)

	// category
	api.Get("/category", categoryHandler.GetListCategory)
	api.Get("/category/:id", categoryHandler.GetDetailCategory)
	api.Post("/category", categoryHandler.CreateCategory)
	api.Put("/category/:id", categoryHandler.UpdateCategory)
	api.Delete("/category/:id", categoryHandler.DeleteCategory)

	// toko
	api.Get("/toko", tokoHandler.GetListToko)

	if err := router.Listen(":8000"); err != nil {
		log.Println("Error: ", err)
	}
}