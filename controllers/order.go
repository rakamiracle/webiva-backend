package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/webiva-backend/config"
)

func CreateOrder(c *gin.Context) {
	var req struct {
		UserID  int `json:"user_id"`
		Items []struct {
			ProductID int `json:"product_id"`
			Quantity  int `json:"quantity"`
			Price     int `json:"price"`
		} `json:"items"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// mulai transaksi
	tx, err := config.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "tx begin fail"})
		return
	}

	var total int64 = 0
	for _, item := range req.Items {
		total += int64(item.Price * item.Quantity)
	}

	res, err := tx.Exec("INSERT INTO orders (user_id,total_price) VALUES (?,?)", req.UserID, total)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	orderID, _ := res.LastInsertId()

	for _, item := range req.Items {
		_, err := tx.Exec("INSERT INTO order_items (order_id,product_id,quantity,price) VALUES (?,?,?,?)",
			orderID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{"message": "order created", "order_id": orderID})
}

func GetOrders(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id,user_id,total_price,status,created_at FROM orders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var orders []map[string]interface{}
	for rows.Next() {
		var id, userID int
		var total int64
		var status string
		var createdAt string

		rows.Scan(&id, &userID, &total, &status, &createdAt)

		orders = append(orders, gin.H{
			"id":         id,
			"user_id":    userID,
			"total":      total,
			"status":     status,
			"created_at": createdAt,
		})
	}
	c.JSON(http.StatusOK, orders)
}
