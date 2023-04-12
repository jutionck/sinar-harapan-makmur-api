package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/middleware"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
	"net/http"
)

type BrandController struct {
	router         *gin.Engine
	usecase        usecase.BrandUseCase
	authMiddleware middleware.AuthTokenMiddleware
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

func NewBrandController(r *gin.Engine, usecase usecase.BrandUseCase, authMiddleware middleware.AuthTokenMiddleware) *BrandController {
	controller := BrandController{
		router:         r,
		usecase:        usecase,
		authMiddleware: authMiddleware,
	}
	r.GET("/brands", authMiddleware.RequireToken(), controller.listHandler)
	r.GET("/brands/:id", authMiddleware.RequireToken(), controller.getByIDHandler)
	r.POST("/brands", authMiddleware.RequireToken(), controller.createHandler)
	r.PUT("/brands", authMiddleware.RequireToken(), controller.updateHandler)
	r.DELETE("/brands/:id", authMiddleware.RequireToken(), controller.deleteHandler)
	return &controller
}
