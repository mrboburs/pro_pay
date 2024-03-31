package response

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ServiceError(err error, errorCode codes.Code) error {
	if err == nil {
		return nil
	} else if strings.Contains(err.Error(), "no rows in result set") {
		return status.Error(codes.NotFound, "data is empty")
	} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		//out := strings.TrimLeft(strings.TrimRight(err.Error(), "_"), "_")
		return status.Error(codes.AlreadyExists, "variable value is already exists")
	} else if strings.Contains(err.Error(), "violates foreign key constraint") {
		return status.Error(codes.InvalidArgument, "variable value is not exists")
	} else if strings.Contains(err.Error(), "no rows affected") {
		return status.Error(codes.NotFound, "variable value is not exists")
	} else if errorCode != codes.OK {
		return status.Error(errorCode, err.Error())
	} else {
		return status.Error(codes.Unknown, err.Error())
	}
}
