package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/api/response"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
)

type BaseApi struct{}

func (b *BaseApi) NewSuccessSingleResponse(c *gin.Context, data interface{}, responseType string) {
	response.SendSingleResponse(c, data, responseType)
}

func (b *BaseApi) NewSuccessPageResponse(c *gin.Context, data []interface{}, responseType string, paging dto.Paging) {
	response.SendPageResponse(c, data, responseType, paging)
}

func (b *BaseApi) NewErrorErrorResponse(c *gin.Context, code int, errorMessage string) {
	response.SendErrorResponse(c, code, errorMessage)
}
