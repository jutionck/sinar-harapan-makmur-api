package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/api"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/utils/common"
)

type VehicleController struct {
	router  *gin.Engine
	usecase usecase.VehicleUseCase
}

func (v *VehicleController) createHandler(c *gin.Context) {
	var payload model.Vehicle
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.SendErrorResponse(
			dto.ResponseStatus{
				Code:        http.StatusInternalServerError,
				Description: err.Error(),
			}))
		return
	}
	if err := v.usecase.SaveData(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.SendErrorResponse(
			dto.ResponseStatus{
				Code:        http.StatusInternalServerError,
				Description: err.Error(),
			}))
		return
	}
	c.JSON(http.StatusCreated, api.SendResponse(
		payload, dto.ResponseStatus{
			Code:        c.Writer.Status(),
			Description: "OK",
		}))
}

func (v *VehicleController) updateHandler(c *gin.Context) {
	var payload model.Vehicle
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.SendErrorResponse(
			dto.ResponseStatus{
				Code:        http.StatusInternalServerError,
				Description: err.Error(),
			}))
		return
	}
	if err := v.usecase.SaveData(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, api.SendResponse(
		payload, dto.ResponseStatus{
			Code:        c.Writer.Status(),
			Description: "OK",
		}))
}

func (v *VehicleController) listHandler(c *gin.Context) {
	requestQueryParams, err := common.ValidateRequestQueryParams(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.SendErrorResponse(
			dto.ResponseStatus{
				Code:        http.StatusInternalServerError,
				Description: err.Error(),
			}))
		return
	}

	vehicles, paging, err := v.usecase.Paging(requestQueryParams)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.SendErrorResponse(
			dto.ResponseStatus{
				Code:        http.StatusInternalServerError,
				Description: err.Error(),
			}))
		return
	}

	var vehicleInterface []interface{}
	for _, v := range vehicles {
		vehicleInterface = append(vehicleInterface, v)
	}
	c.JSON(http.StatusOK, api.SendPageResponse(
		vehicleInterface, dto.ResponseStatus{
			Code:        c.Writer.Status(),
			Description: "OK",
		}, paging))
}

func (v *VehicleController) getByIDHandler(c *gin.Context) {
	id := c.Param("id")
	vehicle, err := v.usecase.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.SendErrorResponse(
			dto.ResponseStatus{
				Code:        http.StatusInternalServerError,
				Description: err.Error(),
			}))
		return
	}
	c.JSON(http.StatusOK, api.SendResponse(
		vehicle, dto.ResponseStatus{
			Code:        c.Writer.Status(),
			Description: "OK",
		}))
}

func (v *VehicleController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := v.usecase.DeleteData(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.SendErrorResponse(
			dto.ResponseStatus{
				Code:        http.StatusInternalServerError,
				Description: err.Error(),
			}))
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewVehicleController(r *gin.Engine, usecase usecase.VehicleUseCase) *VehicleController {
	controller := VehicleController{
		router:  r,
		usecase: usecase,
	}
	r.GET("/vehicles", controller.listHandler)
	r.GET("/vehicles/:id", controller.getByIDHandler)
	r.POST("/vehicles", controller.createHandler)
	r.PUT("/vehicles", controller.updateHandler)
	r.DELETE("/vehicles/:id", controller.deleteHandler)
	return &controller
}
