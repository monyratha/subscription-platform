package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func httpServer() {
	r := gin.Default()
	r.POST("/site", func(c *gin.Context) {
		type req struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}
		var request req
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"site_id": "site-123"})
	})
	log.Println("SiteService HTTP server running on :8082")
	r.Run(":8082")
}
