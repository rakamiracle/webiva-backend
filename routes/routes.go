package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/username/webiva-backend/controllers"
)

func Register(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/auth/register", controllers.Register)
	api.POST("/auth/login", controllers.Login)

	api.GET("/categories", controllers.ListCategories)
	api.POST("/categories", controllers.CreateCategory)

	api.GET("/products", controllers.ListProducts)
	api.POST("/products", controllers.CreateProduct)
}
