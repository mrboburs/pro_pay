package minio

import "io"

type Minio struct {
	MinioMethod
}
type MinioMethod interface {
	UploadImage(imageFile io.Reader, imageSize int64, contextType string) (string, error)
	GetImageLink(imageName string) (string, error)
	UploadDoc(docFile io.Reader, docSize int64, contextType string) (string, error)
	DeleteFile(fileName string) error
}

func NewMinio(minioWorker MinioMethod) *Minio {
	return &Minio{MinioMethod: minioWorker}
}
