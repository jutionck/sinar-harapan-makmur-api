package response

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"net/http"
)

func SendSingleResponse(c *gin.Context, data interface{}, responseType string) {
	c.JSON(http.StatusOK, &dto.SingleResponse{
		Status: dto.ResponseStatus{
			Code:        http.StatusOK,
			Description: responseType,
		},
		Data: data,
	})
}

func SendPageResponse(c *gin.Context, data []interface{}, responseType string, paging dto.Paging) {
	c.JSON(http.StatusOK, &dto.PagedResponse{
		Status: dto.ResponseStatus{
			Code:        http.StatusOK,
			Description: responseType,
		},
		Data:   data,
		Paging: paging,
	})
}

func SendErrorResponse(c *gin.Context, code int, errorMessage string) {
	c.AbortWithStatusJSON(code, &dto.ResponseStatus{
		Code:        code,
		Description: errorMessage,
	})
}
