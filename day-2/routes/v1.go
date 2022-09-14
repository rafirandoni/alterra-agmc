package routes

import (
	"day-2/controllers"

	"github.com/labstack/echo/v4"
)

func NewRouteV1(e *echo.Echo) {
	bookController := controllers.NewBookController()
	userController := controllers.NewUserController()

	v1 := e.Group("/v1")

	bookRoute := v1.Group("/books")
	bookRoute.GET("", bookController.GetLists)
	bookRoute.GET("/:id", bookController.GetDetail)
	bookRoute.POST("", bookController.Store)
	bookRoute.PUT("/:id", bookController.Update)
	bookRoute.DELETE("/:id", bookController.Delete)

	userRoute := v1.Group("/users")
	userRoute.GET("", userController.GetLists)
	userRoute.GET("/:id", userController.GetDetail)
	userRoute.POST("", userController.Store)
	userRoute.PUT("/:id", userController.Update)
	userRoute.DELETE("/:id", userController.Delete)
}
