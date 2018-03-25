package controllers

import (
	"github.com/gin-gonic/gin"
	"NBA/bledar/NBA/functions"
)

func InitRoutes() {
	r := gin.Default()
		r.POST("/login", functions.CheckLogin)
		r.POST("/register",functions.Register)
		r.GET("/getPlayers", functions.GetPlayers)
		r.GET("/getPlayer/:id", functions.GetPlayer)
		r.GET("/leaders/:year/points", functions.MostPoints)
		r.GET("/leaders/:year/assists", functions.MostAssists)
		r.GET("/leaders/:year/rebounds", functions.MostRebounds)
		r.GET("/leaders/:year/blocks", functions.MostBlocks)
		r.GET("/leaders/:year/steals", functions.MostSteals)
		r.POST("/addPlayer", functions.AddPlayer)
		r.PUT("/updatePlayer/:id", functions.UpdatePlayer)
		r.DELETE("deletePlayer/:id", functions.DeletePlayer)
		r.GET("/getTeams", functions.GetTeams)
		r.GET("/getTeam/:id", functions.GetTeam)
		r.GET("teamRanking/:year", functions.TeamRank)
		r.GET("/team/:id/players", functions.GetTeamPlayers)
		r.POST("/addTeam", functions.AddTeam)
		r.DELETE("/deleteTeam/:id", functions.DeleteTeam)
	r.Run(":8080")
}

