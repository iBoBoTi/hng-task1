package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/hng-task1/internal/handler"
)

func main() {
	router := gin.Default()
	router.GET("/my-profile", handler.HandleShowStudentProfile)

	log.Println("starting server at port: 3000")
	http.ListenAndServe(":3000", router)

}
