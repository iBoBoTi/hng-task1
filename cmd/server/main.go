package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/hng-task1/internal/handler"
	"github.com/iBoBoTi/hng-task1/internal/middleware"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("GIN_MODE")
	if env != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("couldn't load env vars: %v", err)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Recovery())
	// setup cors
	router.Use(middleware.CORS())
	router.GET("/my-profile", handler.HandleShowStudentProfile)
	router.POST("/math", handler.HandleMathOperation)

	log.Printf("starting server at port: %s", port)
	router.Run(fmt.Sprintf(":%s", port))

}
