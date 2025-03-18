package main

import (
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

	// services
	jwtService := middleware.NewJWTService()
	authService := services.NewAuthService(jwtService, userRepository, provinceCityRepository)
	provinceCityService := services.NewProvinceCityRepository(provinceCityRepository)

	// handleres
	authHandler := handlers.NewAuthHandler(authService)
	provinceCityHandler := handlers.NewProvinceCityHandler(provinceCityService)

	api := router.Group("/api/v1")
	// auth
	api.Post("/auth/register", authHandler.RegisterUser)
	api.Post("/auth/login", authHandler.LoginUser)

	// province city
	api.Get("/provcity/listprovincies", provinceCityHandler.GetListProvince)
	api.Get("/provcity/detailprovince/:prov_id", provinceCityHandler.GetDetailProvince)
	api.Get("/provcity/listcities/:prov_id", provinceCityHandler.GetListCity)
	api.Get("/provcity/detailcity/:city_id", provinceCityHandler.GetDetailCity)

	router.Listen(":8000")
}