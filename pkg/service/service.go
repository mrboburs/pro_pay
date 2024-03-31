package service

import (
	"pro_pay/config"
	"pro_pay/pkg/repository"

	"pro_pay/pkg/service/atmost"
	"pro_pay/pkg/service/minio"
	// "pro_pay/pkg/service/atmost"

	"pro_pay/pkg/store"
	"pro_pay/tools/logger"
)

type Service struct {
	MinioService *minio.MinioService

	Atmost *atmost.Atmost
}

func NewService(repos *repository.Repository, store *store.Store,
	config *config.Configuration, loggers *logger.Logger) *Service {
	return &Service{
		MinioService: minio.NewMinioService(store, loggers),

		Atmost: atmost.NewAtmostService(repos, store, config, loggers),
	}
}
