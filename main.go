package main

import (
	// "html/template"
	// "log"
	"net/http"
	// "os"

	"github.com/gin-gonic/gin"
	"shasco.app/baseball/models"
)

func main() {

	// API
	err := models.ConnectDatabase()
	checkError(err)

	router := gin.Default()

	// API v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", getWelcomeMessage)
		v1.GET("/team", getTeams)
		v1.GET("/team/:id", getTeamById)
		v1.GET("/team/:id/player", getPlayersByTeamId)
		v1.GET("/player", getPlayers)
		v1.GET("/player/:id", getPlayerById)
		//v1.GET("/player/:pitcherId/pitch", getPitchesByPlayerId)
		//v1.GET("/player/:pitcherId/pitch/against/:batterId", getPitchesByPlayerIds)
		v1.GET("/pitch/pitcher/:pitcherId", getPitchesByPitcherId)
		v1.GET("/pitch/pitcher/:pitcherId/batter/:batterId", getPitchesByPitcherIdVsBatterId)
	}

	router.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// WLECOME //

func getWelcomeMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Welcoome to the GO Baseball API"})
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

func getPitchesByPitcherId(c *gin.Context) {
	pitcherId := c.Param("pitcherId")

	pitches, err := models.GetPitchesByPitcherId(pitcherId)
	checkError(err)

	if pitches == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Fuond"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": pitches})
	}
}

func getPitchesByPitcherIdVsBatterId(c *gin.Context) {
	pitcherId := c.Param("pitcherId")
	batterId := c.Param("batterId")

	pitches, err := models.GetPitchesByPitcherIdVsBatterId(pitcherId, batterId)
	checkError(err)

	if pitches == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Fuond"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": pitches})
	}
}
