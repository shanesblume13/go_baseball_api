package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// API v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/pitch", getPitches)
		v1.GET("/pitch/:id", getPitchById)
		v1.OPTIONS("/pitch", options)
	}

	router.Run()
}

func getPitches(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pitches called",
	})
}

func getPitchById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "getPitchById " + id + " called",
	})
}

func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "options called",
	})
}
