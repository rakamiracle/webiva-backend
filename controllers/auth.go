package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/webiva-backend/config"
	"github.com/username/webiva-backend/utils"
)

func Register(c *gin.Context) {
    var req struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    hash, err := utils.HashPassword(req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
        return
    }

    _, err = config.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
        req.Name, req.Email, hash)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}
func Login(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    var id int
    var hash, role string
    err := config.DB.QueryRow("SELECT id, password, role FROM users WHERE email=?", req.Email).
        Scan(&id, &hash, &role)
    if err == sql.ErrNoRows || !utils.CheckPassword(hash, req.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
        return
    }

    token, err := utils.GenerateToken(uint(id), role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token, "role": role})
}