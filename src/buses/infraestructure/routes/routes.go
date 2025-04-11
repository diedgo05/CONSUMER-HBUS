package routes

import (
	"consumer/src/buses/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/buses")

	updateBus := dependencies.UpdateBusController()


	routes.PUT("/:idBus", updateBus.Run)
}