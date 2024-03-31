package handler

import (
	"pro_pay/pkg/handler/minio"
	"pro_pay/pkg/handler/atmost"

	"pro_pay/pkg/service"
	"pro_pay/tools/logger"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	MinIO *minio.MinIOHandler

	Payment *atmost.PaymentHandler
}

func NewHandler(service *service.Service, loggers *logger.Logger) *Handler {
	return &Handler{

		MinIO: minio.NewMinIOHandler(service, loggers),

		Payment: atmost.NewPaymentHandler(service, loggers),
	}
}

func (handler *Handler) Routers(route *gin.Engine) {
	minio.MinIORouter(route, handler.MinIO)

	atmost.PaymentRouter(route, handler.Payment)

}
