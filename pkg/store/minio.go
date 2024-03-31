package store

import (
	"pro_pay/config"
	"pro_pay/tools/logger"
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinioConnection(config *config.MinioStore, loggers *logger.Logger) (*minio.Client, error) {
	ctx := context.Background()
	// Initialize minio client object.
	minioClient, errInit := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKeyID, config.MinioSecretKey, ""),
		Secure: config.MinioUseSSL,
	})
	if errInit != nil {
		loggers.Fatalln(errInit)
	}
	err := minioClient.MakeBucket(ctx, config.MinioBucketName, minio.MakeBucketOptions{Region: config.MinioLocation})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, config.MinioBucketName)
		if errBucketExists != nil && !exists {
			loggers.Fatal(err)
		}
	}
	return minioClient, errInit
}
