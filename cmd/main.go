package main

import (
	"image-uploader-service/database"
	"image-uploader-service/server"
)

func init() {
	database.Connect()
}

func main() {

	server.Server()
	// r := gin.Default()
	// baseroutes.Routes(r)
	// fmt.Println(config.X)
	// config.Demo()
	// base.base(r)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong reloaded",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// curl -i -X POST -H "Content-Type: multipart/form-data" -F "data=@test.mp3" http://api/v1/image/upload/
