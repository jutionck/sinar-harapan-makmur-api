package api

import "github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"

func SendResponse(data interface{}, status dto.ResponseStatus) *dto.SingleResponse {
	return &dto.SingleResponse{
		Status: status,
		Data:   data,
	}
}

func SendPageResponse(data []interface{}, status dto.ResponseStatus, paging dto.Paging) *dto.PagedResponse {
	return &dto.PagedResponse{
		Status: status,
		Data:   data,
		Paging: paging,
	}
}

func SendErrorResponse(status dto.ResponseStatus) *dto.ResponseStatus {
	return &dto.ResponseStatus{
		Code:        status.Code,
		Description: status.Description,
	}
}
