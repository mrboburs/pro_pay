package minio

import (
	// "pro_pay/package/service"
	"pro_pay/pkg/service"
	"pro_pay/tools/logger"

	"github.com/gin-gonic/gin"
)

type MinIOHandler struct {
	MinIOEndPoint
}
type MinIOEndPoint interface {
	UploadImages(ctx *gin.Context)
	UploadImage(ctx *gin.Context)
	UploadDoc(ctx *gin.Context)
	DownloadFile(ctx *gin.Context)
	TransferFile(ctx *gin.Context)
}

func NewMinIOHandler(service *service.Service,
	loggers *logger.Logger) *MinIOHandler {
	return &MinIOHandler{
		MinIOEndPoint: NewMinIOEndPointHandler(service, loggers),
	}
}
