package server

import (
	"image-uploader-service/pkg/image_service"

	"github.com/gin-gonic/gin"
)

// Main routing
func Server() {
	router := gin.Default()

	v1 := router.Group("/api/v1/image")
	{
		image_service.SubRoutes(v1)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// curl -i -X POST -H "Content-Type: multipart/form-data" -F "file=@Screenshot 2020-01-24 at 6.15.06 PM.png" http://localhost:8080/api/v1/image/upload

// curl -X GET -H "Content-Type: application/json" http://localhost:8080/api/v1/image/list?offset=2

// curl -X GET -H "Content-Type: application/json" http://localhost:8080/api/v1/image/detail?id=4
