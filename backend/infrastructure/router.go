package infrastructure

import (
	"os"
	"log"
	"payment-app/backend/handler"

	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Router : router api server
var Router *gin.Engine

func init() {
	router := gin.Default()

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("CLIENT_CORS_ADDR")}
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string("Origin", "Content-Type")
	}))

	router.GET("/api/v1/items", func(c *gin.Context) { handler.GetLists(c) })
	router.GET("/api/v1/items/:id", func(c *gin.Context) { handler.GetItem(c) })
	router.POST("/api/v1/charge/items/:id", func(c *gin.Context) { handler.Charge(c) })

	Router = router
}