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
	alamatRepository := repositories.NewAlamatRepository(db)
	productRepository := repositories.NewProductRepository(db)
	// trxRepository := repositories.NewTRXRepository(db)

	// services
	authService := services.NewAuthService(userRepository, tokoRepository, provinceCityRepository)
	userService := services.NewUserService(userRepository, alamatRepository)
	provinceCityService := services.NewProvinceCityService(provinceCityRepository)
	categoryService := services.NewCategoryService(categoryRepository)
	tokoService := services.NewTokoService(tokoRepository)
	productService := services.NewProductService(productRepository, tokoRepository, categoryRepository)
	// trxService := services.NewTRXService(trxRepository)

	// handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	provinceCityHandler := handlers.NewProvinceCityHandler(provinceCityService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	tokoHandler := handlers.NewTokoHandler(tokoService)
	productHandler := handlers.NewProductHandler(productService)
	// trxHandler := handlers.NewTRXHandler(trxService)

	// middleware
	authMiddleware := middleware.AuthMiddleware(userService)

	api := router.Group("/api/v1")
	// auth
	api.Post("/auth/register", authHandler.RegisterUser)
	api.Post("/auth/login", authHandler.LoginUser)

	// user
	api.Get("/user", authMiddleware, userHandler.GetMyProfile)
	api.Put("/user", authMiddleware, userHandler.UpdateProfile)
	api.Get("/user/alamat", authMiddleware, userHandler.GetMyAlamat)
	api.Get("/user/alamat/:id", authMiddleware, userHandler.GetDetailAlamat)
	api.Post("/user/alamat", authMiddleware, userHandler.CreateAlamatUser)
	api.Put("/user/alamat/:id", authMiddleware, userHandler.UpdateAlamatUser)
	api.Delete("/user/alamat/:id", authMiddleware, userHandler.DeleteAlamatUser)

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
	api.Get("/toko/my", authMiddleware, tokoHandler.MyToko)
	api.Get("/toko", authMiddleware, tokoHandler.GetListToko)
	api.Get("/toko/:id", authMiddleware, tokoHandler.GetDetailToko)
	// api.Put("/toko", tokoHandler.UpdateProfileToko)

	// product
	// api.Get("/product", authMiddleware, productHandler.GetListProduct)
	api.Get("/product/:id", authMiddleware, productHandler.GetDetailProduct)
	api.Post("/product", authMiddleware, productHandler.CreateProduct)
	api.Put("/product/:id", authMiddleware, productHandler.UpdateProduct)
	api.Delete("/product/:id", authMiddleware, productHandler.DeleteProduct)

	// trx
	// api.Get("/trx". trxHandler.GetListTRX)
	// api.Get("/trx/:id". trxHandler.GetDetailTRX)
	// api.Post("/trx". trxHandler.CreatelTRX)

	if err := router.Listen(":8000"); err != nil {
		log.Println("Error: ", err)
	}
}