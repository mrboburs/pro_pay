package atmost

import (
	"pro_pay/model"
	"pro_pay/pkg/service"

	"pro_pay/tools/logger"
	"pro_pay/tools/response"

	"github.com/gin-gonic/gin"
)

const (
	TransactionID = "id"
)

type TransactionEndPointHandler struct {
	service *service.Service
	loggers *logger.Logger
}

func NewTransactionEndPointHandler(service *service.Service,
	loggers *logger.Logger) *TransactionEndPointHandler {
	return &TransactionEndPointHandler{service: service, loggers: loggers}
}

// Create Transaction
// @Description Create Transaction
// @Summary Create Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param create body model.CreateTransaction true " Create Transaction"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/transactions [post]
func (h *TransactionEndPointHandler) CreateTransaction(ctx *gin.Context) {

	
	var (
		body model.CreateTransaction
	)

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	out, err := h.service.Atmost.CreateTransaction(body)
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil, out, nil)
}
