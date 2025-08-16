package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/webiva-backend/config"
	"github.com/username/webiva-backend/models"
)

func GetStore(c *gin.Context) {
	var s models.StoreSetting
	config.DB.FirstOrCreate(&s, models.StoreSetting{ID: 1})
	c.JSON(http.StatusOK, s)
}

func UpdateStore(c *gin.Context) {
	var s models.StoreSetting
	config.DB.FirstOrCreate(&s, models.StoreSetting{ID: 1})

	var payload models.StoreSetting
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.StoreName = payload.StoreName
	s.LogoURL = payload.LogoURL
	s.PaymentMethods = payload.PaymentMethods
	config.DB.Save(&s)
	c.JSON(http.StatusOK, s)
}
