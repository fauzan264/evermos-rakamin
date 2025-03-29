package main

import (
	"fmt"
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
	cfg := config.LoadConfig()
	db := config.InitDatabase()

	router := fiber.New()
	router.Use(cors.New())
	router.Static("/uploads", "./uploads")

	// API external for data province & city
	provinceCityApiURL := os.Getenv("API_LOCATION")

	// repositories
	provinceCityRepository := repositories.NewProvinceCityRepository(provinceCityApiURL)
	userRepository := repositories.NewUserRepository(db)
	shopRepository := repositories.NewShopRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)
	addressRepository := repositories.NewAddressRepository(db)
	productRepository := repositories.NewProductRepository(db)
	trxRepository := repositories.NewTRXRepository(db)

	// services
	authService := services.NewAuthService(userRepository, shopRepository, provinceCityRepository)
	userService := services.NewUserService(userRepository, addressRepository)
	provinceCityService := services.NewProvinceCityService(provinceCityRepository)
	categoryService := services.NewCategoryService(categoryRepository)
	shopService := services.NewShopService(shopRepository)
	productService := services.NewProductService(productRepository, shopRepository, categoryRepository)
	trxService := services.NewTRXService(trxRepository, productRepository, addressRepository, shopRepository, categoryRepository)

	// handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	provinceCityHandler := handlers.NewProvinceCityHandler(provinceCityService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	shopHandler := handlers.NewShopHandler(shopService)
	productHandler := handlers.NewProductHandler(productService)
	trxHandler := handlers.NewTRXHandler(trxService)

	// middleware
	authMiddleware := middleware.AuthMiddleware(userService)

	api := router.Group("/api/v1")
	// auth
	api.Post("/auth/register", authHandler.RegisterUser)
	api.Post("/auth/login", authHandler.LoginUser)

	// user
	api.Get("/user", authMiddleware, userHandler.GetMyProfile)
	api.Put("/user", authMiddleware, userHandler.UpdateProfile)
	api.Get("/user/alamat", authMiddleware, userHandler.GetMyAddress)
	api.Get("/user/alamat/:id", authMiddleware, userHandler.GetDetailAddress)
	api.Post("/user/alamat", authMiddleware, userHandler.CreateAddressUser)
	api.Put("/user/alamat/:id", authMiddleware, userHandler.UpdateAddressUser)
	api.Delete("/user/alamat/:id", authMiddleware, userHandler.DeleteAddressUser)

	// province city
	api.Get("/provcity/listprovincies", provinceCityHandler.GetListProvince)
	api.Get("/provcity/detailprovince/:prov_id", provinceCityHandler.GetDetailProvince)
	api.Get("/provcity/listcities/:prov_id", provinceCityHandler.GetListCity)
	api.Get("/provcity/detailcity/:city_id", provinceCityHandler.GetDetailCity)

	// category
	api.Get("/category", authMiddleware, categoryHandler.GetListCategory)
	api.Get("/category/:id", authMiddleware, categoryHandler.GetDetailCategory)
	api.Post("/category", authMiddleware, categoryHandler.CreateCategory)
	api.Put("/category/:id", authMiddleware, categoryHandler.UpdateCategory)
	api.Delete("/category/:id", authMiddleware, categoryHandler.DeleteCategory)

	// toko
	api.Get("/toko/my", authMiddleware, shopHandler.MyShop)
	api.Get("/toko", authMiddleware, shopHandler.GetListShop)
	api.Get("/toko/:id_toko", authMiddleware, shopHandler.GetDetailShop)
	api.Put("/toko/:id_toko", authMiddleware, shopHandler.UpdateProfileShop)

	// product
	api.Get("/product", authMiddleware, productHandler.GetListProduct)
	api.Get("/product/:id", authMiddleware, productHandler.GetDetailProduct)
	api.Post("/product", authMiddleware, productHandler.CreateProduct)
	api.Put("/product/:id", authMiddleware, productHandler.UpdateProduct)
	api.Delete("/product/:id", authMiddleware, productHandler.DeleteProduct)

	// trx
	api.Get("/trx", authMiddleware, trxHandler.GetListTRX)
	api.Get("/trx/:id", authMiddleware, trxHandler.GetDetailTRX)
	api.Post("/trx", authMiddleware, trxHandler.CreateTRX)

	if err := router.Listen(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)); err != nil {
		log.Println("Error: ", err)
	}
}