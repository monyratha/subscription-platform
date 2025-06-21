package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func httpServer() {
	r := gin.Default()
	r.POST("/view", func(c *gin.Context) {
		type req struct {
			SiteID   string `json:"site_id"`
			Referrer string `json:"referrer"`
		}
		var request req
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	log.Println("StatService HTTP server running on :8085")
	r.Run(":8085")
}
