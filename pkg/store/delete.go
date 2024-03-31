package store

import (
	"pro_pay/config"
	"pro_pay/tools/logger"
	"context"

	"github.com/minio/minio-go/v7"
)

type DeleteMinio struct {
	minio   *minio.Client
	config  *config.MinioStore
	loggers *logger.Logger
}

func NewDeleteMinio(minio *minio.Client, config *config.MinioStore, loggers *logger.Logger) *DeleteMinio {
	return &DeleteMinio{minio: minio, config: config, loggers: loggers}
}
func (m *DeleteMinio) DeleteFile(fileName string) error {
	err := m.minio.RemoveObject(context.Background(), m.config.MinioBucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
