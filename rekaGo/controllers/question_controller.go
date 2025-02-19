package controllers

import (
	"net/http"
	"strconv"

	"rekaGo/services"

	"github.com/gin-gonic/gin"
)

// QuestionController handles HTTP requests for questions
type QuestionController struct {
	service services.QuestionService
}

// NewQuestionController creates a new QuestionController
func NewQuestionController(service services.QuestionService) *QuestionController {
	return &QuestionController{
		service: service,
	}
}

// GetQuestions handles GET /questions?limit=60
func (ctrl *QuestionController) GetQuestions(c *gin.Context) {
	// Define the limit; default is 60
	limitStr := c.DefaultQuery("limit", "60")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	questions, err := ctrl.service.GetQuestions(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}
