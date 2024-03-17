package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uqyanzie/platform-library-grpc-api/controllers"
)

type PlatformRouteController struct {
	platformController controllers.PlatformController
}

func NewPlatformRouteController(platformController controllers.PlatformController) PlatformRouteController {
	return PlatformRouteController{platformController: platformController}
}

func (pl *PlatformRouteController) PlatformRoutes(route *gin.RouterGroup) {

	route = route.Group("/platform")

	route.POST("/", pl.platformController.CreateNewPlatform)
}
