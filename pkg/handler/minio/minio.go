package minio

import (
	// "pro_pay/package/service"
	// "pro_pay/pkg/service"
	"pro_pay/model"
	"pro_pay/pkg/service"
	// "pro_pay/package/service"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"pro_pay/tools/handler_func"
	"pro_pay/tools/logger"
	"pro_pay/tools/response"

	"github.com/gin-gonic/gin"
)

const (
	contentType = "Content-Type"
	filePath    = "file-path"
)

type MinIOEndPointHandler struct {
	service *service.Service
	loggers *logger.Logger
}

func NewMinIOEndPointHandler(service *service.Service,
	loggers *logger.Logger) *MinIOEndPointHandler {
	return &MinIOEndPointHandler{service: service, loggers: loggers}
}

var (
	errFile              = errors.New("error when close file multipart ")
	errInvalidFileFormat = errors.New("invalid file format")
)

const (
	jpgContentType  = "image/jpg"
	pngContentType  = "image/png"
	jpegContentType = "image/jpeg"
	xlsxContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	xlsContentType  = "application/vnd.ms-excel"
	docContentType  = "application/msword"
	pdfContentType  = "application/pdf"
	docxContentType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
)

// UploadImage
// @Description Upload Image
// @Tags Upload File
// @Accept       json
// @Produce application/octet-stream
// @Produce image/png
// @Produce image/jpeg
// @Produce image/jpg
// @Param file formData file true "file"
// @Accept multipart/form-data
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/upload/upload-image [post]
// @Security ApiKeyAuth
func (h *MinIOEndPointHandler) UploadImage(ctx *gin.Context) {
	// ctx.Request.ParseMultipartForm(1 << 25)
	loggers := h.loggers
	file, err := ctx.FormFile("file")
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	imageContentType := file.Header[contentType][0]
	if imageContentType != jpegContentType && imageContentType != jpgContentType && imageContentType != pngContentType {
		response.HandleResponse(ctx, response.BadRequest, errInvalidFileFormat, nil, nil)
		return
	}
	var fileIO io.Reader
	fileMultipart, err := file.Open()
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	fileIO = fileMultipart
	imageFileName, err := h.service.MinioService.UploadImage(fileIO,
		file.Size,
		imageContentType)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	defer func(fileMultipart multipart.File) {
		err := fileMultipart.Close()
		if err != nil {
			loggers.Error(errFile, err)
		}
	}(fileMultipart)
	response.HandleResponse(ctx, response.OK, nil, imageFileName, nil)
}

// UploadImages
// @Description Upload Images
// @Tags Upload File
// @Accept       json
// @Produce application/octet-stream
// @Produce image/png
// @Produce image/jpeg
// @Produce image/jpg
// @Param files formData file true "files"
// @Accept multipart/form-data
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/upload/upload-images [post]
// @Security ApiKeyAuth
func (h *MinIOEndPointHandler) UploadImages(ctx *gin.Context) {
	// ctx.Request.ParseMultipartForm(1 << 25)
	var uploadedFiles []model.Files
	form, err := ctx.MultipartForm()
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	files := form.File["files"]
	for _, file := range files {
		imageContentType := file.Header[contentType][0]
		if imageContentType != jpegContentType && imageContentType != jpgContentType && imageContentType != pngContentType {
			continue
		}
		var fileIO io.Reader
		fileMultipart, err := file.Open()
		if err != nil {
			response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
			return
		}
		fileIO = fileMultipart
		imageFileName, err := h.service.MinioService.UploadImage(fileIO,
			file.Size,
			imageContentType)
		if err != nil {
			response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
			return
		}
		imageLink, err := h.service.MinioService.GetImageLink(
			imageFileName)
		if err != nil {
			response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
			return
		}
		uploadedFiles = append(uploadedFiles, model.Files{Link: imageLink, Name: imageFileName})
		defer func(fileMultipart multipart.File) {
			err := fileMultipart.Close()
			if err != nil {
				h.loggers.Error(errFile,
					err)
			}
		}(fileMultipart)
	}
	response.HandleResponse(ctx, response.OK, nil, uploadedFiles, nil)
}

// UploadDoc
// @Description Upload doc
// @Tags Upload File
// @Accept       json
// @Produce application/octet-stream
// @Produce application/msword
// @Produce application/vnd.openxmlformats-officedocument.wordprocessingml.document
// @Produce image/jpg
// @Param file formData file true "file"
// @Accept multipart/form-data
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/upload/upload-doc [post]
// @Security ApiKeyAuth
func (h *MinIOEndPointHandler) UploadDoc(ctx *gin.Context) {
	// ctx.Request.ParseMultipartForm(1 << 25)
	loggers := h.loggers
	file, err := ctx.FormFile("file")
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	fileContentType := file.Header[contentType][0]
	if fileContentType != docContentType && fileContentType != docxContentType {
		response.HandleResponse(ctx, response.BadRequest, errInvalidFileFormat, nil, nil)
		return
	}
	var fileIO io.Reader
	fileMultipart, err := file.Open()
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	fileIO = fileMultipart
	docFileName, err := h.service.MinioService.UploadDoc(fileIO, file.Size, fileContentType)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	defer func(fileMultipart multipart.File) {
		err := fileMultipart.Close()
		if err != nil {
			loggers.Error(errFile, err)
		}
	}(fileMultipart)
	docLink, err := h.service.MinioService.GetImageLink(
		docFileName)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, model.Files{Link: docLink, Name: docFileName}, nil)
}

// FileTransfer
// @Description File Transfer
// @Tags Transfer File
// @Accept json
// @Accept application/json
// @Param file-path query string true "file path"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/transfer [post]
// @Security ApiKeyAuth
func (h *MinIOEndPointHandler) TransferFile(ctx *gin.Context) {
	filePath, err := handler_func.GetStringQuery(ctx, filePath)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	_, err = os.ReadFile(filePath)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	if path.Ext(ctx.Request.URL.Path) == ".xlsx" {
		ctx.Header(contentType, xlsxContentType)
	}
	err = handler_func.FileTransfer(ctx, filePath, filePath)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
}

// DownloadFile
// @Description Download File
// @Tags Download File
// @Accept json
// @Accept application/json
// @Param file-path query string true "file path"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/download [get]
func (h *MinIOEndPointHandler) DownloadFile(ctx *gin.Context) {
	filePath, err := handler_func.GetStringQuery(ctx, filePath)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	_, err = os.ReadFile(filePath)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	err = handler_func.DownloadFile(ctx, filePath)
	if err != nil {
		response.HandleResponse(ctx, response.Internal, err, nil, nil)
		return
	}
}
