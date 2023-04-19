package router

import (
	"chal9/controllers"
	"chal9/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/", controllers.GetProducts)
		productRouter.GET("/:productId", controllers.GetProductById)
		productRouter.PUT("/:productId", controllers.UpdateProduct)
		productRouter.DELETE("/:productId", controllers.DeleteProduct)

	}
	return r
}
