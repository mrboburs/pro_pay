package store

import (
	"pro_pay/config"
	"pro_pay/model"
	"pro_pay/tools/logger"
	"context"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

type FileLinkMinio struct {
	minio   *minio.Client
	config  *config.MinioStore
	loggers *logger.Logger
}

func NewFileLinkMinio(minio *minio.Client, config *config.MinioStore, loggers *logger.Logger) *FileLinkMinio {
	return &FileLinkMinio{minio: minio, config: config, loggers: loggers}
}
func (store *FileLinkMinio) GetImageUrl(imageName string) (string, error) {
	expiry := time.Second * 24 * 60 * 60 * 7
	loggers := store.loggers
	client := store.minio
	configuration := store.config
	presignedURL, err := client.PresignedGetObject(context.Background(), configuration.MinioBucketName, imageName, expiry, url.Values{})
	if err != nil {
		loggers.Error("Error while getting object URL: ", err.Error())
		return "", err
	}
	return presignedURL.String(), nil
}
func (store *FileLinkMinio) GetFileUrlList(fileNameList []model.Files) (fileList []model.Files, err error) {
	expiry := time.Second * 24 * 60 * 60 * 7
	loggers := store.loggers
	client := store.minio
	configuration := store.config
	for i := range fileNameList {
		if len(fileNameList[i].Name) != 0 {
			presignedURL, err := client.PresignedGetObject(context.Background(), configuration.MinioBucketName, fileNameList[i].Name, expiry, url.Values{})
			if err != nil {
				loggers.Error("Error while getting object URL: ", err.Error())
				return fileList, err
			}
			fileList[i].Link = presignedURL.String()
		}
	}
	return fileList, nil
}
