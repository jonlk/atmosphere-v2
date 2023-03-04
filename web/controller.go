package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup) {
	group.GET("/dewpoint", dewpointRoute)
	group.GET("/relativehumidity", relativeHumidityRoute)
}
