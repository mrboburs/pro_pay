package handler_func

import (
	"pro_pay/auth/middleware"
	"pro_pay/config"
	"pro_pay/model"
	"pro_pay/tools/logger"
	"pro_pay/tools/response"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	cfg     = config.Config()
	loggers = logger.GetLogger()
)

const (
	ctxUserID     = middleware.CtxUserID
	ctxRoleID     = middleware.CtxRoleID
	paramInvalid  = "%s param is invalid"
	queryInvalid  = "%s query is invalid"
	ParseDate     = "2006-01-02T05:00:00.000Z"
	FormatDate    = "2006-01-02T05:00:00.000Z"
	orderAsc      = "asc"
	orderDesc     = "desc"
	queryPageSize = "page-size"
)

func GetStringParam(ctx *gin.Context, query string) (param string, err error) {
	param = ctx.Param(query)
	if param == "" {
		err := fmt.Sprintf(" %s param is empty", query)
		return "", errors.New(err)
	}
	return param, nil
}
func GetNullStringParam(ctx *gin.Context, query string) (param string, err error) {
	param = ctx.Param(query)
	return param, nil
}
func GetNullStringQuery(ctx *gin.Context, query string) (param string) {
	param = ctx.Query(query)
	param = strings.Trim(param, " ")
	return param
}
func GetStringQuery(ctx *gin.Context, query string) (param string, err error) {
	param = ctx.Query(query)
	if param == "" {
		err := fmt.Sprintf(" %s param is empty", query)
		return "", errors.New(err)
	}
	return param, nil
}

func GetDateOrderQuery(ctx *gin.Context, query string) (dateOrder string, err error) {
	dateOrder = strings.ToLower(ctx.Query(query))
	if dateOrder == "" {
		err := "date-order query is empty"
		return "", errors.New(err)
	}

	if dateOrder != "asc" && dateOrder != "desc" {
		return "", fmt.Errorf("invalid date-order value")
	}

	return dateOrder, nil
}

func GetInt64Query(ctx *gin.Context,
	query string) (int64,
	error) {
	param := ctx.Query(query)
	if param == "" {
		return 0, nil
	}
	paramInt, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, err
	}
	return paramInt, nil
}
func GetFloat64Query(ctx *gin.Context, query string) (float64,
	error) {
	param := ctx.Query(query)
	if param == "" {
		return 0, nil
	}
	paramInt, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, err
	}
	return paramInt, nil
}
func GetArrayStringQuery(ctx *gin.Context,
	query string) ([]string, error) {
	param := ctx.Query(query)
	if param == "" {
		return []string{}, errors.New("param is empty")
	}
	chunks := strings.Split(param, ",")
	return chunks, nil
}
func JsonUnmarshal(pointData *interface{},
	data []byte) error {
	err := json.Unmarshal(data, pointData)
	if err != nil {
		return err
	}
	return nil
}
func GetBooleanQuery(ctx *gin.Context, query string) (bool, error) {
	param := ctx.Query(query)
	if param == "" {
		err := fmt.Sprintf(paramInvalid, query)
		return false, errors.New(err)
	}
	boolVal, err := strconv.ParseBool(param)
	if err != nil {
		err := fmt.Sprintf(paramInvalid, query)
		return false, errors.New(err)
	}
	return boolVal, nil
}
func GetUUIDQuery(ctx *gin.Context, query string) (uuid.UUID, error) {
	param := ctx.Query(query)
	if param == "" {
		err := fmt.Sprintf(paramInvalid, query)
		return uuid.Nil, errors.New(err)
	}
	paramUUID, err := uuid.Parse(param)
	if err != nil {
		loggers.Error(err)
		err := fmt.Sprintf(paramInvalid, query)
		return uuid.Nil, errors.New(err)
	}
	return paramUUID, nil
}
func GetUserId(ctx *gin.Context) (string, error) {
	id, ok := ctx.Get(middleware.CtxUserID)
	if !ok {
		loggers.Error(response.ErrorUserIDInvalid.Error())
		return "", errors.New(response.ErrorUserIDInvalid.Error())
	}
	userID, ok := id.(string)
	if !ok {
		loggers.Error(response.ErrorUserIDInvalid.Error())
		return "", errors.New(response.ErrorUserIDInvalid.Error())
	}
	_, err := uuid.Parse(userID)
	if err != nil {
		loggers.Error(err)
		return "", err
	}
	return userID, nil
}
func GetCtxUserID(ctx *gin.Context) (uuid.UUID, error) {
	id, ok := ctx.Get(ctxUserID)
	if !ok {
		loggers.Error(response.ErrorUserIDInvalid.Error())
		return uuid.Nil, errors.New(response.ErrorUserIDInvalid.Error())
	}
	userIDString, ok := id.(string)
	if !ok {
		loggers.Error(response.ErrorUserIDInvalid.Error())
		return uuid.Nil, errors.New(response.ErrorUserIDInvalid.Error())
	}
	userID, err := uuid.Parse(userIDString)
	if err != nil {
		loggers.Error(err)
		return uuid.Nil, err
	}
	return userID, nil
}
func GetCtxRoleID(ctx *gin.Context) (uuid.UUID, error) {
	id, ok := ctx.Get(ctxRoleID)
	if !ok {
		loggers.Error(response.ErrorRoleIDInvalid.Error())
		return uuid.Nil, errors.New(response.ErrorRoleIDInvalid.Error())
	}
	roleIDString, ok := id.(string)
	if !ok {
		loggers.Error(response.ErrorRoleIDInvalid.Error())
		return uuid.Nil, errors.New(response.ErrorRoleIDInvalid.Error())
	}
	roleID, err := uuid.Parse(roleIDString)
	if err != nil {
		loggers.Error(err)
		return uuid.Nil, err
	}
	return roleID, nil
}

func GetPageQuery(ctx *gin.Context) (offset int64,
	err error) {
	offsetStr := ctx.DefaultQuery("page", cfg.APIProperty.DefaultPage)
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return 0, response.ErrorNotANumberPage
	}
	if offset < 0 {
		return 0, response.ErrorOffsetNotAUnsignedInt
	}
	return offset, nil
}
func GetPageSizeQuery(ctx *gin.Context) (limit int64, err error) {
	limitStr := ctx.DefaultQuery(queryPageSize, cfg.APIProperty.DefaultPageSize)
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return 0, response.ErrorNotANumberPageSize
	}
	if limit < 0 {
		return 0, response.ErrorLimitNotAUnsignedInt
	}
	return limit, nil
}
func CalculationPagination(page, pageSize int64) (offset, limit int64) {
	if page < 0 {
		page = 1
	}
	offset = (page - 1) * pageSize
	limit = pageSize
	return offset, limit
}
func ListPagination(ctx *gin.Context) (pagination model.Pagination, err error) {
	page, err := GetPageQuery(ctx)
	if err != nil {
		loggers.Error(err)
		return pagination, err
	}
	pageSize, err := GetPageSizeQuery(ctx)
	if err != nil {
		loggers.Error(err)
		return pagination, err
	}
	offset, limit := CalculationPagination(page, pageSize)
	pagination.Limit = limit
	pagination.Offset = offset
	pagination.Page = page
	pagination.PageSize = pageSize
	return pagination, nil
}
func GetUUIDParam(ctx *gin.Context, query string) (uuid.UUID, error) {
	queryData := ctx.Param(query)
	if queryData == "" {
		err := fmt.Sprintf(queryInvalid, queryData)
		return uuid.Nil, errors.New(err)
	}
	queryUUID, err := uuid.Parse(queryData)
	if err != nil {
		err := fmt.Sprintf(queryInvalid, queryData)
		return uuid.Nil, errors.New(err)
	}
	return queryUUID, nil
}
func CheckTime(query string) (time.Time, error) {
	if query != "" {
		checkTime, err := time.Parse(ParseDate, query)
		if err != nil {
			err := fmt.Sprintf(queryInvalid, ParseDate)
			return time.Time{}, errors.New(err)
		}
		return checkTime, nil
	}
	return time.Time{}, nil
}
func GetNullDateQuery(ctx *gin.Context, query string) (time.Time, error) {
	queryDate := GetNullStringQuery(ctx, query)
	date, err := CheckTime(queryDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
func GetNullUUIDQuery(ctx *gin.Context, query string) (uuid.UUID, error) {
	queryData := ctx.Query(query)
	if queryData != "" {
		queryUUID, err := uuid.Parse(queryData)
		if err != nil {
			err := fmt.Sprintf(queryInvalid, queryData)
			return uuid.Nil, errors.New(err)
		}
		return queryUUID, nil
	}
	return uuid.Nil, nil
}
func GetNullUUIDParam(ctx *gin.Context, query string) (uuid.UUID, error) {
	queryData := ctx.Param(query)
	if queryData != "" {
		queryUUID, err := uuid.Parse(queryData)
		if err != nil {
			err := fmt.Sprintf(queryInvalid, queryData)
			return uuid.Nil, errors.New(err)
		}
		return queryUUID, nil
	}
	return uuid.Nil, nil
}
func GetNullBooleanQuery(ctx *gin.Context, query string) (bool, error) {
	param := ctx.Query(query)
	if param != "" {
		boolVal, err := strconv.ParseBool(param)
		if err != nil {
			err := fmt.Sprintf(paramInvalid, query)
			return false, errors.New(err)
		}
		return boolVal, nil
	}
	return true, nil
}
func GetNullInt64Param(ctx *gin.Context, query string) (int64, error) {
	queryData := ctx.Query(query)
	if queryData != "" {
		queryInt, err := strconv.ParseInt(queryData, 10, 64)
		if err != nil {
			err := fmt.Sprintf(queryInvalid, queryData)
			return 0, errors.New(err)
		}
		return queryInt, nil
	}
	return 0, nil
}
func GetNullIntParam(ctx *gin.Context, query string) (int, error) {
	queryData := ctx.Query(query)
	if queryData != "" {
		queryInt, err := strconv.Atoi(queryData)
		if err != nil {
			err := fmt.Sprintf(queryInvalid, queryData)
			return 0, errors.New(err)
		}
		return queryInt, nil
	}
	return 0, nil
}
func GetNullFloat64Param(ctx *gin.Context, query string) (float64, error) {
	queryData := ctx.Query(query)
	if queryData != "" {
		queryFloat, err := strconv.ParseFloat(queryData, 64)
		if err != nil {
			err := fmt.Sprintf(queryInvalid, queryData)
			return 0, errors.New(err)
		}
		return queryFloat, nil
	}
	return 0, nil
}
func GetNullArrayStringQuery(ctx *gin.Context, query string) ([]string, error) {
	queryData := ctx.Query(query)
	if queryData != "" {
		chunks := strings.Split(queryData, ",")
		return chunks, nil
	}
	return []string{}, nil
}
func ResponseHeaderXTotalCountWrite(ctx *gin.Context, total int64) {
	ctx.Writer.Header().Set("X-Total-Count", strconv.Itoa(int(total)))
}
func FileTransfer(ctx *gin.Context, filePath string, contentType string) (err error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", "attachment; filename="+path.Base(filePath))
	ctx.Data(http.StatusOK, "application/octet-stream", bytes)
	ctx.Writer.Header().Set("Content-Type", contentType)
	return nil
}

func DownloadFile(ctx *gin.Context, filePath string) error {
	_, fileName := filepath.Split(filePath)
	bytes, err := os.ReadFile("./" + filePath)
	if err != nil {
		return err
	}
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Writer.Write(bytes)
	return nil
}
func GetAcceptLanguage(ctx *gin.Context) string {
	acceptedLanguage := ctx.GetHeader("Accept-Language")
	if len(acceptedLanguage) == 2 {
		return acceptedLanguage
	}
	if len(acceptedLanguage) > 2 {
		langData := strings.Split(acceptedLanguage, ";")
		if len(langData) >= 2 {
			langParts := strings.Split(langData[0], ",")
			if len(langParts) >= 2 {
				lang := langParts[1]
				if len(lang) == 2 {
					return lang
				}
			}
			return ""
		}
		return ""
	}
	return acceptedLanguage
}
