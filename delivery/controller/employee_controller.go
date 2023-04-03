package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
	"net/http"
)

type EmployeeController struct {
	router  *gin.Engine
	usecase usecase.EmployeeUseCase
}

func (e *EmployeeController) createHandler(c *gin.Context) {
	var payload model.Employee
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := e.usecase.SaveData(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}

func (e *EmployeeController) updateHandler(c *gin.Context) {
	var payload model.Employee
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := e.usecase.SaveData(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payload)
}

func (e *EmployeeController) listHandler(c *gin.Context) {
	vehicles, err := e.usecase.FindAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve employee data"})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

func (e *EmployeeController) getByIDHandler(c *gin.Context) {
	id := c.Param("id")
	vehicle, err := e.usecase.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve employee data"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

func (e *EmployeeController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := e.usecase.DeleteData(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewEmployeeController(r *gin.Engine, usecase usecase.EmployeeUseCase) *EmployeeController {
	controller := EmployeeController{
		router:  r,
		usecase: usecase,
	}
	r.GET("/employees", controller.listHandler)
	r.GET("/employees/:id", controller.getByIDHandler)
	r.POST("/employees", controller.createHandler)
	r.PUT("/employees", controller.updateHandler)
	r.DELETE("/employees/:id", controller.deleteHandler)
	return &controller
}
