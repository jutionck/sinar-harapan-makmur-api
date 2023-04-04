package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/api"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/utils/common"
)

type VehicleController struct {
	router  *gin.Engine
	usecase usecase.VehicleUseCase
	api.BaseApi
}

func (v *VehicleController) createHandler(c *gin.Context) {
	var payload model.Vehicle
	if err := c.ShouldBindJSON(&payload); err != nil {
		v.NewErrorErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := v.usecase.SaveData(&payload); err != nil {
		v.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	v.NewSuccessSingleResponse(c, payload, "OK")
}

func (v *VehicleController) updateHandler(c *gin.Context) {
	var payload model.Vehicle
	if err := c.ShouldBindJSON(&payload); err != nil {
		v.NewErrorErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := v.usecase.SaveData(&payload); err != nil {
		v.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	v.NewSuccessSingleResponse(c, payload, "OK")
}

func (v *VehicleController) listHandler(c *gin.Context) {
	requestQueryParams, err := common.ValidateRequestQueryParams(c)
	if err != nil {
		v.NewErrorErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	vehicles, paging, err := v.usecase.Paging(requestQueryParams)
	if err != nil {
		v.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	var vehicleInterface []interface{}
	for _, v := range vehicles {
		vehicleInterface = append(vehicleInterface, v)
	}
	v.NewSuccessPageResponse(c, vehicleInterface, "OK", paging)
}

func (v *VehicleController) getByIDHandler(c *gin.Context) {
	id := c.Param("id")
	vehicle, err := v.usecase.FindById(id)
	if err != nil {
		v.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	v.NewSuccessSingleResponse(c, vehicle, "OK")
}

func (v *VehicleController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := v.usecase.DeleteData(id)
	if err != nil {
		v.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
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
