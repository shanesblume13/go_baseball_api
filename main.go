package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shasco.app/baseball/models"
)

func main() {
	err := models.ConnectDatabase()
	checkError(err)

	router := gin.Default()

	// API v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/team", getTeams)
		v1.GET("/team/:id", getTeamById)
		v1.OPTIONS("/pitch", options)
	}

	router.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getTeams(c *gin.Context) {
	teams, err := models.GetTeams(50)
	checkError(err)

	if teams == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Fuond"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": teams})
	}
}

func getTeamById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "getTeamById " + id + " called",
	})
}

func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "options called",
	})
}
