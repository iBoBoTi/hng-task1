package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/hng-task1/internal/domain"
)

func HandleMathOperation(c *gin.Context) {
	var reqMath domain.Math

	if err := c.ShouldBindJSON(&reqMath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	response := domain.SolveMath(&reqMath)
	if response == nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "invalid operation"})
		return
	}
	c.JSON(http.StatusOK, response)
}
