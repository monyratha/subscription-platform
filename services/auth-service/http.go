package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func httpServer() {
	r := gin.Default()
	r.POST("/register", func(c *gin.Context) {
		type req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		var request req
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// return dummy token
		c.JSON(http.StatusOK, gin.H{"user_id": "123", "token": "dummy-token"})
	})
	log.Println("AuthService HTTP server running on :8081")
	r.Run(":8081")
}
