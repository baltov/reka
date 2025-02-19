package main

import (
	"rekaGo/controllers"
	"rekaGo/database"
	"rekaGo/services"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	questionService := services.NewQuestionService(database.DB)
	imageService := services.NewImageService(database.DB)

	questionController := controllers.NewQuestionController(questionService)
	imageController := controllers.NewImageController(imageService)

	router := gin.Default()

	// Define the endpoints
	router.GET("/api/questions", questionController.GetQuestions)
	router.GET("/api/images/:id", imageController.GetImage)

	// Start the server
	err := router.Run(":9090")
	if err != nil {
		panic("Failed to run the server: " + err.Error())
	}
}
