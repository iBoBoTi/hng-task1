package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/hng-task1/internal/domain"
)

func HandleShowStudentProfile(c *gin.Context) {
	myProfile := &domain.StudentProfile{
		SlackUserName: "BoBoTi",
		Backend:       true,
		Age:           27,
		Bio:           "I am software engineer in the backend stack with proficiency in the Go language...",
	}
	c.JSON(http.StatusOK, myProfile)
}
