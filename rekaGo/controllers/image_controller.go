package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"rekaGo/services"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	service services.ImageService
}

func NewImageController(service services.ImageService) *ImageController {
	return &ImageController{
		service: service,
	}
}

func (ctrl *ImageController) GetImage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	image, err := ctrl.service.GetImage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", image.Name))

	c.Header("Content-Type", "image/jpeg")
	c.Data(http.StatusOK, "image/jpeg", image.Data)

}
