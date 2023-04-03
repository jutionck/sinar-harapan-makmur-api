package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
	"net/http"
)

type BrandController struct {
	router  *gin.Engine
	usecase usecase.BrandUseCase
}

func (b *BrandController) createHandler(c *gin.Context) {
	var payload model.Brand
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := b.usecase.SaveData(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}

func (b *BrandController) updateHandler(c *gin.Context) {
	var payload model.Brand
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := b.usecase.SaveData(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payload)
}

func (b *BrandController) listHandler(c *gin.Context) {
	vehicles, err := b.usecase.FindAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve brand data"})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

func (b *BrandController) getByIDHandler(c *gin.Context) {
	id := c.Param("id")
	vehicle, err := b.usecase.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve brand data"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

func (b *BrandController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := b.usecase.DeleteData(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewBrandController(r *gin.Engine, usecase usecase.BrandUseCase) *BrandController {
	controller := BrandController{
		router:  r,
		usecase: usecase,
	}
	r.GET("/brands", controller.listHandler)
	r.GET("/brands/:id", controller.getByIDHandler)
	r.POST("/brands", controller.createHandler)
	r.PUT("/brands", controller.updateHandler)
	r.DELETE("/brands/:id", controller.deleteHandler)
	return &controller
}
