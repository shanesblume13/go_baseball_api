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
		v1.GET("team/:id/player", getPlayersByTeamId)
		v1.GET("/player", getPlayers)
		v1.GET("player/:id", getPlayerById)
	}

	router.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// TEAMS //

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

	team, err := models.GetTeamById(id)
	checkError(err)

	if team.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": team})
	}
}

// PLAYERS //

func getPlayers(c *gin.Context) {
	players, err := models.GetPlayers(5000)
	checkError(err)

	if players == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Fuond"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": players})
	}
}

func getPlayerById(c *gin.Context) {
	id := c.Param("id")

	player, err := models.GetPlayerById(id)
	checkError(err)

	if player.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": player})
	}
}

func getPlayersByTeamId(c *gin.Context) {
	teamId := c.Param("id")

	players, err := models.GetPlayersByTeamId(teamId)
	checkError(err)

	if players == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Fuond"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": players})
	}
}
