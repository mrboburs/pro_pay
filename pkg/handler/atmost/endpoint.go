package atmost

import (
	"github.com/gin-gonic/gin"
)

func PaymentRouter(api *gin.Engine, handler *PaymentHandler) {
	baseUrlPrivate := api.Group("/api/v1")
	// baseUrlPublic := api.Group("/api/v1")
	{
		payment := baseUrlPrivate.Group("/transactions")
		{
			payment.POST("", handler.TransactionEndPoint.CreateTransaction)

		}

	}

}
