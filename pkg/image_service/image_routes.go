package image_service

import (
	"github.com/gin-gonic/gin"
)

func SubRoutes(v1 *gin.RouterGroup) {
	v1.GET("/list", listImage)
	v1.POST("/upload", uploadImage)
	v1.GET("/detail", getImage)
}
