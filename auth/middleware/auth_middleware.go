package middleware

import (
	"pro_pay/auth/jwt"
	"pro_pay/tools/response"
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	CtxUserID           = "userID"
	CtxRoleID           = "roleID"
)

var (
	errUnAuth        = errors.New("unauthorized. header is empty")
	errUnAuthPayload = errors.New("unauthorized. payload is invalid")
)

func AuthRequestHandler(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		HandleResponse(ctx, response.Unauthorized, errUnAuth, nil)
		return
	}
	user, err := jwt.ExtractTokenMetadata(ctx)
	if err != nil {
		HandleResponse(ctx, response.Unauthorized, err, nil)
		return
	}
	if user == nil {
		// Abort the request with the appropriate error code
		HandleResponse(ctx, response.Unauthorized,
			errUnAuthPayload, nil)
		return
	}
	ctx.Set(CtxUserID, user.UserID)
	ctx.Set(CtxRoleID, user.RoleID)
	// err = rbac.HasPermission(user.RoleID, ctx.FullPath(), ctx.Request.Method)
	// if err != nil {
	// 	HandleResponse(ctx, response.PermissionDenied, nil, nil)
	// 	return
	// }
	// Continue down the chain to handler etc
	ctx.Next()
}

func AuthRefreshTokenRequestHandler(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		HandleResponse(ctx, response.Unauthorized,
			errUnAuth, nil)
		return
	}
	user, err := jwt.ExtractRefreshTokenMetadata(ctx)
	if err != nil {
		HandleResponse(ctx, response.Unauthorized,
			errUnAuth, nil)
		return
	}
	if user == nil {
		// Abort the request with the appropriate error code
		HandleResponse(ctx, response.Unauthorized,
			errUnAuthPayload, nil)
		return
	}
	ctx.Set(CtxUserID, user.UserID)
	//Continue down the chain to handler etc
	ctx.Next()
}
func HandleResponse(ctx *gin.Context, status response.Status,
	err error, data interface{}) {
	ctx.AbortWithStatusJSON(status.Code, response.ResponseModel{
		Status:       status.Status,
		Code:         status.Code,
		Description:  status.Description,
		SnapData:     data,
		ErrorMessage: err.Error(),
	})
}
