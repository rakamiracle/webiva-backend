package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/username/webiva-backend/controllers"
	"github.com/username/webiva-backend/middlewares"
)

func Register(r *gin.Engine) {
	api := r.Group("/api")

	// healthcheck
	api.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	// auth
	api.POST("/auth/register", controllers.Register)
	api.POST("/auth/login", controllers.Login)

	// public
	api.GET("/products", controllers.ListProducts)
	api.GET("/products/:id", controllers.GetProduct)
	api.GET("/categories", controllers.ListCategories)
	api.GET("/store", controllers.GetStore)

	// protected: user
	user := api.Group("/")
	user.Use(middlewares.AuthRequired())
	{
		user.POST("/orders", controllers.CreateOrder)
		user.GET("/orders", controllers.MyOrders)
	}

	// admin only
	admin := api.Group("/")
	admin.Use(middlewares.AuthRequired(), middlewares.AdminOnly())
	{
		admin.POST("/products", controllers.CreateProduct)
		admin.PUT("/products/:id", controllers.UpdateProduct)
		admin.DELETE("/products/:id", controllers.DeleteProduct)

		admin.POST("/categories", controllers.CreateCategory)
		admin.PUT("/categories/:id", controllers.UpdateCategory)
		admin.DELETE("/categories/:id", controllers.DeleteCategory)

		admin.PUT("/store", controllers.UpdateStore)

		admin.GET("/orders/all", controllers.AllOrders)
		admin.PUT("/orders/:id/status", controllers.UpdateOrderStatus)
	}
}
