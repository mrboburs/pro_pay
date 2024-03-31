package store

import (
	"pro_pay/config"
	"pro_pay/tools/logger"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type UploadMinio struct {
	minio   *minio.Client
	config  *config.MinioStore
	loggers *logger.Logger
}

var (
	docContentType  = "msword"
	docxContentType = "vnd.openxmlformats-officedocument.wordprocessingml.document"
)
var (
	xlsxContentType = "vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	xlsContentType  = "vnd.ms-excel"
)

func NewUploadMinio(minio *minio.Client, config *config.MinioStore, loggers *logger.Logger) *UploadMinio {
	return &UploadMinio{minio: minio, config: config, loggers: loggers}
}
func (um *UploadMinio) UploadImage(file io.Reader, imageSize int64, contextType string) (string, error) {
	loggers := um.loggers
	fileName := uuid.New()
	fileExtension := strings.Split(contextType, "/")[1]
	imageName := fileName.String() + "." + fileExtension
	_, err := um.minio.PutObject(context.Background(), um.config.MinioBucketName, imageName, file, imageSize, minio.PutObjectOptions{ContentType: contextType})
	if err != nil {
		loggers.Error("Internal Server Error: ", err.Error())
		return "", err
	}
	return imageName, nil
}
func (um *UploadMinio) UploadDoc(file io.Reader, docSize int64,
	contentType string) (string, error) {
	loggers := um.loggers
	fileName := uuid.New()
	fileExtension := "docx"
	if strings.Contains(contentType, docContentType) {
		fileExtension = "doc"
	}
	if strings.Contains(contentType, docxContentType) {
		fileExtension = "docx"
	}
	docFileName := fileName.String() + "." + fileExtension
	_, err := um.minio.PutObject(context.Background(),
		um.config.MinioBucketName, docFileName, file, docSize,
		minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		loggers.Error("Internal Server Error: ", err.Error())
		return "", err
	}
	return docFileName, nil
}
func (um *UploadMinio) UploadExcel(filePath string) (string, error) {
	loggers := um.loggers
	fileName := uuid.New()
	fileExtension := "xlsx"
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fileSize, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}
	docFileName := strings.Split(filepath.Base(filePath), ".")[0] + " " + fileName.String() + "." + fileExtension
	_, err = um.minio.PutObject(context.Background(),
		um.config.MinioBucketName, docFileName, file, fileSize.Size(),
		minio.PutObjectOptions{ContentType: xlsxContentType})
	if err != nil {
		loggers.Error("Internal Server Error: ", err.Error())
		return "", err
	}
	return docFileName, nil
}
