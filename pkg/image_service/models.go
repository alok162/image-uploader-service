package image_service

import (
	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model
	FileName string
	Url      string
	Size     int64
}
