package response

import (
	"github.com/gin-gonic/gin"
)

func HandleResponse(ctx *gin.Context, status Status, err error, data interface{}, pagination interface{}) {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	ctx.AbortWithStatusJSON(status.Code, ResponseListModel{
		Status:       status.Status,
		Code:         status.Code,
		Description:  status.Description,
		Pagination:   pagination,
		SnapData:     data,
		ErrorMessage: errorMessage,
	})
}

type ResponseListModel struct {
	Status       string      `json:"status,omitempty"`
	Code         int         `json:"code,omitempty"`
	Description  string      `json:"description,omitempty"`
	SnapData     interface{} `json:"snapData,omitempty"`
	Pagination   interface{} `json:"pagination,omitempty"`
	ErrorMessage string      `json:"error,omitempty"`
}
type ResponseModel struct {
	Status       string      `json:"status,omitempty"`
	Code         int         `json:"code,omitempty"`
	Description  string      `json:"description,omitempty"`
	SnapData     interface{} `json:"snapData,omitempty"`
	ErrorMessage string      `json:"error,omitempty"`
}
