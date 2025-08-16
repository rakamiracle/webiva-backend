package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/webiva-backend/config"
	"github.com/username/webiva-backend/models"
)

type CreateOrderReq struct {
	Items []struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required,min=1"`
	} `json:"items" binding:"required"`
	ProofURL string `json:"proof_url"`
}

func CreateOrder(c *gin.Context) {
	var req CreateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil || len(req.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	uid := c.GetUint("user_id")

	var total int64
	var items []models.OrderItem
	for _, it := range req.Items {
		var prod models.Product
		if err := config.DB.First(&prod, it.ProductID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
			return
		}
		total += int64(it.Quantity) * prod.Price
		items = append(items, models.OrderItem{
			ProductID: prod.ID,
			Quantity:  it.Quantity,
			Price:     prod.Price,
		})
	}

	order := models.Order{
		UserID:     uid,
		Status:     "pending",
		TotalPrice: total,
		Items:      items,
		ProofURL:   req.ProofURL,
	}
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Preload("Items.Product").First(&order, order.ID)
	c.JSON(http.StatusCreated, order)
}

func MyOrders(c *gin.Context) {
	uid := c.GetUint("user_id")
	var orders []models.Order
	config.DB.Preload("Items.Product").Where("user_id = ?", uid).Order("id DESC").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func AllOrders(c *gin.Context) {
	var orders []models.Order
	config.DB.Preload("User").Preload("Items.Product").Order("id DESC").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func UpdateOrderStatus(c *gin.Context) {
	type Req struct{ Status string `json:"status" binding:"required"` }
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var order models.Order
	if err := config.DB.First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	order.Status = req.Status
	config.DB.Save(&order)
	c.JSON(http.StatusOK, order)
}
