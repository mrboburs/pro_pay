package atmost

import (
	// "pro_pay/pkg/service"
	"pro_pay/pkg/service"
	"pro_pay/tools/logger"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	TransactionEndPoint
	AuthEndPoint
}
type AuthEndPoint interface {
	// GetToken(ctx *gin.Context)
}
type TransactionEndPoint interface {
	CreateTransaction(ctx *gin.Context)
}

func NewPaymentHandler(service *service.Service, loggers *logger.Logger) *PaymentHandler {
	return &PaymentHandler{
AuthEndPoint: NewAuthEndPointHandler(service,loggers),
		TransactionEndPoint: NewTransactionEndPointHandler(service, loggers),
	}
}
