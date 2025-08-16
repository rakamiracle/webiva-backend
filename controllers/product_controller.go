package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/webiva-backend/config"
	"github.com/username/webiva-backend/models"
)

func ListProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Preload("Category").Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	var p models.Product
	if err := config.DB.Preload("Category").First(&p, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func CreateProduct(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func UpdateProduct(c *gin.Context) {
	var p models.Product
	if err := config.DB.First(&p, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	var payload models.Product
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.Name = payload.Name
	p.Description = payload.Description
	p.Price = payload.Price
	p.Stock = payload.Stock
	p.ImageURL = payload.ImageURL
	p.CategoryID = payload.CategoryID
	config.DB.Save(&p)
	c.JSON(http.StatusOK, p)
}

func DeleteProduct(c *gin.Context) {
	if err := config.DB.Delete(&models.Product{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
