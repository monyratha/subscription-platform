package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func httpServer() {
	r := gin.Default()
	r.POST("/activate", func(c *gin.Context) {
		type req struct {
			UserID string `json:"user_id"`
			PlanID string `json:"plan_id"`
		}
		var request req
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"subscription_id": "sub-123"})
	})
	log.Println("SubscriptionService HTTP server running on :8083")
	r.Run(":8083")
}
