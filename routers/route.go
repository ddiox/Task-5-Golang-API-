package routers

import (
	"github.com/ddiox/evermos_api/handlers"
	"github.com/ddiox/evermos_api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	CategoryRoute(api)
	AuthRoute(api)
	AlamatRoute(api)
	UserRoute(api)
	TokoRoute(api)
	ProdukRoute(api)
	TrxRoute(api)
	ProvCityRoute(api)
}

func CategoryRoute(r fiber.Router) {
	r.Post("/category", middlewares.AuthMiddleware(), handlers.CreateCategory)
	r.Get("/category", middlewares.AuthMiddleware(), handlers.GetAllCategory)
	r.Get("/category/:id", middlewares.AuthMiddleware(), handlers.GetCategoryById)
	r.Put("/category/:id", middlewares.AuthMiddleware(), handlers.UpdateCategoryById)
	r.Delete("/category/:id", middlewares.AuthMiddleware(), handlers.DeleteCategoryById)
}

func AuthRoute(r fiber.Router) {
	r.Post("/auth/register", handlers.Register)
	r.Post("/auth/login", handlers.Login)
}

func AlamatRoute(r fiber.Router) {
	r.Post("/alamat", middlewares.AuthMiddleware(), handlers.CreateAlamat)
	r.Get("/alamat", middlewares.AuthMiddleware(), handlers.GetAllAlamat)
	r.Get("/alamat/:id", middlewares.AuthMiddleware(), handlers.GetAlamatById)
	r.Put("/alamat/:id", middlewares.AuthMiddleware(), handlers.UpdateAlamatById)
	r.Delete("/alamat/:id", middlewares.AuthMiddleware(), handlers.DeleteAlamatById)
}

func UserRoute(r fiber.Router) {
	r.Get("/user", middlewares.AuthMiddleware(), handlers.GetAllUser)
	r.Get("/user/:id", middlewares.AuthMiddleware(), handlers.GetUserById)
	r.Put("/user/:id", middlewares.AuthMiddleware(), handlers.UpdateUserById)
}

func TokoRoute(r fiber.Router) {
	r.Post("/toko", middlewares.AuthMiddleware(), handlers.CreateToko)
	r.Get("/toko", middlewares.AuthMiddleware(), handlers.GetAllToko)
	r.Get("/toko/:id", middlewares.AuthMiddleware(), handlers.GetTokoById)
	r.Put("/toko/:id", middlewares.AuthMiddleware(), handlers.UpdateTokoById)
	r.Delete("/toko/:id", middlewares.AuthMiddleware(), handlers.DeleteTokoById)
}

func ProdukRoute(r fiber.Router) {
	r.Post("/produk", middlewares.AuthMiddleware(), handlers.CreateProduct)
	r.Get("/produk", middlewares.AuthMiddleware(), handlers.GetAllProducts)
	r.Get("/produk/:id", middlewares.AuthMiddleware(), handlers.GetProductById)
	r.Put("/produk/:id", middlewares.AuthMiddleware(), handlers.UpdateProductById)
	r.Delete("/produk/:id", middlewares.AuthMiddleware(), handlers.DeleteProductById)
}

func TrxRoute(r fiber.Router) {
	r.Post("/transaction", middlewares.AuthMiddleware(), handlers.CreateTrx)
	r.Get("/transaction", middlewares.AuthMiddleware(), handlers.GetAllTrx)
	r.Get("/transaction/:id", middlewares.AuthMiddleware(), handlers.GetTrxById)
	r.Put("/transaction/:id", middlewares.AuthMiddleware(), handlers.UpdateTrxById)
	r.Delete("/transaction/:id", middlewares.AuthMiddleware(), handlers.DeleteTrxById)
}

func ProvCityRoute(r fiber.Router) {
	r.Get("/provcity/listprovincies", handlers.GetListProvince)
	r.Get("/provcity/detailprovince/:id", handlers.GetDetailProvince)
	r.Get("/provcity/listcities/:id", handlers.GetListCity)
	r.Get("/provcity/detailcity/:city_id", handlers.GetDetailCity)
}
