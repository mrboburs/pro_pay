package store

import (
	"pro_pay/config"
	"pro_pay/tools/logger"
	"context"
	"errors"

	"github.com/minio/minio-go/v7"
)

type ObjectMinio struct {
	minio   *minio.Client
	config  *config.MinioStore
	loggers *logger.Logger
}

func NewObjectMinio(minio *minio.Client, config *config.MinioStore, loggers *logger.Logger) *ObjectMinio {
	return &ObjectMinio{minio: minio, config: config, loggers: loggers}
}
func (om *ObjectMinio) ObjectExists(imageName string) error {
	loggers := om.loggers
	client := om.minio
	configuration := om.config
	_, err := client.StatObject(context.Background(), configuration.MinioBucketName, imageName, minio.GetObjectOptions{})
	if err != nil {
		errResponse := minio.ToErrorResponse(err)
		if errResponse.Code == "AccessDenied" {
			loggers.Error(errResponse)
			return errors.New("access Denied")
		}
		if errResponse.Code == "NoSuchBucket" {
			loggers.Error(errResponse)
			return errors.New("no Exist Bucket Object")
		}
		if errResponse.Code == "InvalidBucketName" {
			loggers.Error(errResponse)
			return errors.New("invalid Bucket Name")
		}
		if errResponse.Code == "NoSuchKey" {
			return errors.New("no Exist File Object")
		}
		return errors.New("unknown Error")
	}
	return nil
}
