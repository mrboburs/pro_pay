package atmost

import (
	"pro_pay/pkg/service"

	"pro_pay/tools/logger"
)

type AuthEndPointHandler struct {
	service *service.Service
	loggers *logger.Logger
}

func NewAuthEndPointHandler(service *service.Service,
	loggers *logger.Logger) *AuthEndPointHandler {
	return &AuthEndPointHandler{service: service, loggers: loggers}
}
