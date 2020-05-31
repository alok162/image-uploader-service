package image_service

import (
	"errors"
	"mime/multipart"
	"net/http"
)

func validateFileContentType(out multipart.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	switch contentType {
	case "image/jpeg", "image/jpg":
		return contentType, nil

	case "image/gif":
		return contentType, nil

	case "image/png":
		return contentType, nil

	default:
		return "", errors.New("file not supported")
	}
}
