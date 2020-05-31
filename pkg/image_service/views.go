package image_service

import (
	"bytes"
	"fmt"
	"image-uploader-service/database"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

const (
	servingS3Url = "https://alok-testing.s3.us-east-2.amazonaws.com/media/"
)

const ()

func uploadImage(c *gin.Context) {

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, token)

	cfg := aws.NewConfig().WithRegion(awsS3Region).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)
	file, err := c.FormFile("file")

	if err != nil {
		log.Fatal("error-->", err)
	}

	f, err := file.Open()

	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	fileType, err := validateFileContentType(f)
	if err != nil {
		log.Println("error-->", err)
		c.JSON(400, gin.H{
			"message": "file type not supported",
		})
		return
	}
	fmt.Println("file type ###############", fileType)

	size := file.Size

	buffer := make([]byte, size)

	f.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	// fileType := http.DetectContentType(buffer)
	path := "/media/" + file.Filename
	params := &s3.PutObjectInput{
		Bucket:        aws.String(awsS3Bucket),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}

	_, err = svc.PutObject(params)
	if err != nil {
		log.Println("uploading error", err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
	}

	// Update entry into db
	database.DB.AutoMigrate(&Image{})
	reqBody := Image{FileName: file.Filename, Url: servingS3Url + file.Filename, Size: size}
	database.DB.Create(&reqBody)

	log.Println("image uploaded-->", file.Filename)
	c.JSON(200, gin.H{
		"message": "successfully uploaded",
	})
}

func listImage(c *gin.Context) {

	offset := c.Query("offset")

	var images []Image

	c.JSON(200, gin.H{
		// "message": database.DB.Find(&users, id),
		"message": database.DB.Order("created_at desc").Limit(2).Offset(offset).Find(&images),
	})
}

func getImage(c *gin.Context) {
	var images []Image

	id := c.Query("id")
	if result := database.DB.Find(&images, id); result.Error != nil {
		c.JSON(200, gin.H{
			"message": result,
		})
	}
	c.JSON(200, gin.H{
		"message": "Internal server error",
	})
}
