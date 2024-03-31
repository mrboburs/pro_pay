package store

import (
	"pro_pay/config"
	"pro_pay/model"
	"pro_pay/tools/logger"
	"io"

	"github.com/minio/minio-go/v7"
)

type Store struct {
	UploadStore
	ObjectStore
	FileLinkStore
	DeleteStore
}
type UploadStore interface {
	UploadImage(imageFile io.Reader, imageSize int64, contextType string) (string, error)
	UploadDoc(docFile io.Reader, docSize int64, contextType string) (string, error)
	UploadExcel(filePath string) (string, error)
}
type ObjectStore interface {
	ObjectExists(name string) error
}
type FileLinkStore interface {
	GetImageUrl(imageName string) (string, error)
	GetFileUrlList(imageNameList []model.Files) ([]model.Files, error)
}
type DeleteStore interface {
	DeleteFile(fileName string) error
}

func NewStore(minio *minio.Client, config *config.MinioStore, loggers *logger.Logger) *Store {
	return &Store{
		UploadStore:   NewUploadMinio(minio, config, loggers),
		ObjectStore:   NewObjectMinio(minio, config, loggers),
		FileLinkStore: NewFileLinkMinio(minio, config, loggers),
		DeleteStore:   NewDeleteMinio(minio, config, loggers),
	}
}
