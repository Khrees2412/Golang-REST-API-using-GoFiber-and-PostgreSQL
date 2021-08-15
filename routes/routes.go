package routes

import (
	"github.com/khrees2412/convas/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// User routes
	app.Post("/api/v1/user/register", controllers.Register)
	app.Post("/api/v1/user/login", controllers.Login)
	app.Post("/api/v1/user/logout", controllers.Logout)
	app.Post("/api/v1/change_password", controllers.ChangePassword)
	app.Post("/api/v1/delete_account", controllers.DeleteAccount)


	//Book routes
	app.Post("/api/v1/book/add", controllers.CreateBook)
	app.Get("api/v1/book/get/:id", controllers.GetBook)
	app.Get("api/v1/book/get_all", controllers.GetBooks)
	app.Post("api/v1/book/update/:id", controllers.UpdateBook)
	app.Post("api/v1/book/delete/:id", controllers.DeleteBook)

}
