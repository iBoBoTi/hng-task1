package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/hng-task1/internal/handler"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("GIN_MODE")
	if env != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("couldn't load env vars: %v", err)
		}
	}

	port:= os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	router := gin.Default()
	router.GET("/my-profile", handler.HandleShowStudentProfile)

	log.Printf("starting server at port: %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s",port), router)

}
