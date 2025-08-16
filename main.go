package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/username/webiva-backend/config"
	"github.com/username/webiva-backend/routes"
)

func main() {
	_ = godotenv.Load()
	config.InitDB()

	r := gin.Default()

	// CORS: izinkan FE Nuxt saat dev (3000)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	routes.Register(r)

	port := os.Getenv("APP_PORT")
	if port == "" { port = "8080" }

	log.Println("server running on :" + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
