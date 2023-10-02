package utils

import (
	"api/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseSocketReq(data []byte, req interface{}) *types.WSResponse {
	if err := json.Unmarshal(data, req); err != nil {
		fmt.Printf("unable to parse data from request err: %s\n data: %s", err.Error(), string(data))
		return &types.WSResponse{
			Status: http.StatusInternalServerError,
			Error:  "unable to parse request",
		}
	}

	if err := types.Validate(req); err != nil {
		return &types.WSResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}
	}

	return nil
}

func ReqFailForbidden(errMsg string) *types.WSResponse {
	return &types.WSResponse{
		Status: http.StatusForbidden,
		Error:  errMsg,
	}
}

func ReqFailInternalError(err error) *types.WSResponse {
	return &types.WSResponse{
		Status: http.StatusBadRequest,
		Error:  err.Error(),
	}
}

func ReqFailBadRequest(errMsg string) *types.WSResponse {
	return &types.WSResponse{
		Status: http.StatusBadRequest,
		Error:  errMsg,
	}
}

func ReqSuccess(data interface{}) *types.WSResponse {
	return &types.WSResponse{
		Status: http.StatusOK,
		Data:   data,
	}
}

func ReqSuccessCustomStatus(status int, data interface{}) *types.WSResponse {
	return &types.WSResponse{
		Status: status,
		Data:   data,
	}
}
