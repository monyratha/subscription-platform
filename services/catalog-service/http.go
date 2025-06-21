package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func httpServer() {
	r := gin.Default()
	r.POST("/product", func(c *gin.Context) {
		type req struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		}
		var request req
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"product_id": "prod-123"})
	})
	log.Println("CatalogService HTTP server running on :8084")
	r.Run(":8084")
}
