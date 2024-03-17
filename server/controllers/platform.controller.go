package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uqyanzie/platform-library-grpc-api/models"
	"github.com/uqyanzie/platform-library-grpc-api/services"
)

type PlatformController struct {
	platformService services.PlatformService
}

func NewPlatformController(platformService services.PlatformService) PlatformController {
	return PlatformController{platformService: platformService}
}

func (pl *PlatformController) CreateNewPlatform(c *gin.Context) {
	var platform *models.NewPlatform

	if err := c.ShouldBindJSON(&platform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPlatform, err := pl.platformService.CreateNewPlatform(platform)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPlatform})
}

