package minio

import (
	"pro_pay/auth/middleware"

	"github.com/gin-gonic/gin"
)

func MinIORouter(api *gin.Engine, handler *MinIOHandler) {
	upload := api.Group("/api/v1/upload", middleware.AuthRequestHandler)
	{
		upload.POST("/upload-images", handler.MinIOEndPoint.UploadImages)
		upload.POST("/upload-image", handler.MinIOEndPoint.UploadImage)
		upload.POST("/upload-doc", handler.MinIOEndPoint.UploadDoc)
	}
	download := api.Group("/api/v1/download")
	{
		download.GET("", handler.MinIOEndPoint.DownloadFile)
	}
	transfer := api.Group("/api/v1/transfer", middleware.AuthRequestHandler)
	{
		transfer.POST("", handler.MinIOEndPoint.TransferFile)
	}
}
