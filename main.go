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
	r.Use(cors.Default())

	routes.Register(r)

	port := os.Getenv("APP_PORT")
	if port == "" { port = "8080" }

	log.Println("server running on :" + port)
	r.Run(":" + port)
}
