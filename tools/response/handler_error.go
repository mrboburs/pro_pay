package response

import (
	"errors"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ServiceErrorConvert(ctx *gin.Context,
	serviceError error) {
	st, _ := status.FromError(serviceError)
	err := st.Message()
	if st.Code() == codes.Internal {
		HandleResponse(ctx,
			InternalServerError, nil, nil, nil)
	} else if st.Code() == codes.NotFound {
		HandleResponse(ctx, NotFound, errors.New(err), nil, nil)
	} else if st.Code() == codes.InvalidArgument {
		HandleResponse(ctx,
			InvalidArgument, errors.New(err), nil, nil)
	} else if st.Code() == codes.Unavailable {
		HandleResponse(ctx,
			Unavailable, errors.New(err), nil, nil)
	} else if st.Code() == codes.OK {
		HandleResponse(ctx, OK, errors.New(err), nil, nil)
	} else if st.Code() == codes.AlreadyExists {
		HandleResponse(ctx,
			AlreadyExists, errors.New(err), nil, nil)
	} else if st.Code() == codes.Canceled {
		HandleResponse(ctx, Canceled, errors.New(err), nil, nil)
	} else if st.Code() == codes.Unknown {
		HandleResponse(ctx, Unknown, errors.New(err), nil, nil)
	} else if st.Code() == codes.DeadlineExceeded {
		HandleResponse(ctx,
			DeadlineExceeded, errors.New(err), nil, nil)
	} else if st.Code() == codes.PermissionDenied {
		HandleResponse(ctx,
			PermissionDenied, errors.New(err), nil, nil)
	} else if st.Code() == codes.ResourceExhausted {
		HandleResponse(ctx,
			ResourceExhausted, errors.New(err), nil, nil)
	} else if st.Code() == codes.FailedPrecondition {
		HandleResponse(ctx,
			FailedPrecondition, errors.New(err), nil, nil)
	} else if st.Code() == codes.Aborted {
		HandleResponse(ctx, Aborted, errors.New(err), nil, nil)
	} else if st.Code() == codes.OutOfRange {
		HandleResponse(ctx,
			OutOfRange, errors.New(err), nil, nil)
	} else if st.Code() == codes.Unimplemented {
		HandleResponse(ctx,
			Unimplemented, errors.New(err), nil, nil)
	} else if st.Code() == codes.DataLoss {
		HandleResponse(ctx, DataLoss, errors.New(err), nil, nil)
	} else if st.Code() == codes.Unauthenticated {
		HandleResponse(ctx,
			Unauthorized, errors.New(err), nil, nil)
	} else {
		HandleResponse(ctx, Unknown, errors.New(err), nil, nil)
	}
}
