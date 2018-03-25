package services

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"NBA/bledar/NBA/database"
	"NBA/bledar/NBA/structs"
	"regexp"
)

func EmailValidation (email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func SgetPlayer(id int, c *gin.Context){
	db:= database.Connection()
	var player structs.Player
	rows, err := db.Query("Select * FROM player WHERE id=$1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&player.Id, &player.Name, &player.Surname, &player.Teamid, &player.Age)
	if err != nil {
		c.String(http.StatusOK, "E dhene e futur gabim")
		panic(err)
	}
	c.JSON(http.StatusOK, player)
}

func SaddPlayer(name string, surname string,age int,teamid int, c *gin.Context){
	db:= database.Connection()
	Query := `INSERT INTO player (pname, surname, age, team_id)
	VALUES ($1, $2, $3, $4)`
	var player structs.Player
	player.Name = name
	player.Surname = surname
	player.Age = age
	player.Teamid = teamid
	_, err := db.Exec(Query, name, surname, age, teamid)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, player)
	}
}

func SupdatePlayer (id int, teamId int, c *gin.Context){
	db:= database.Connection()
	Query :=`UPDATE player SET team_id=$1 WHERE id=$s2`
	_, err := db.Exec(Query, teamId, id)
	if err != nil {
		panic(err)
	}else {
		c.String(http.StatusOK,"Player Updated")
	}
}

func SdeletePlayer(id int, c *gin.Context){
	db:= database.Connection()
	Query := `DELETE FROM player WHERE id = $1;`
	_, err := db.Exec(Query, id)
	if err != nil {
		panic(err)
	} else {
		c.String(http.StatusOK, "Player Deleted")
	}
}


func SgetTeam(id int, c *gin.Context){
	db:= database.Connection()
	var team structs.Team
	rows, err := db.Query("Select * FROM team WHERE id=$1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&team.Id, &team.Name)
	if err != nil {
		c.String(http.StatusOK, "E dhene e futur gabim")
		panic(err)
	}
	c.JSON(http.StatusOK, team)
}

func SgetTeamPlayers(id int, c *gin.Context){
	db:= database.Connection()
	var teamPlayers []structs.Player
	rows, err := db.Query("Select * FROM player WHERE team_id=$1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var player structs.Player
		err = rows.Scan(&player.Id, &player.Name, &player.Surname, &player.Teamid, &player.Age)
		if err != nil {
			panic(err)
		}
		teamPlayers = append(teamPlayers, player)
	}

	c.JSON(http.StatusOK, teamPlayers)
}

func SaddTeam(id int, Name string, c *gin.Context) {
	db:= database.Connection()
	Query := `INSERT INTO team (id , tname) VALUES ($1, $2)`
	var team structs.Team
	team.Id = id
	team.Name = Name
	_, err := db.Exec(Query, id, Name)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, team)
	}
}

func SdeleteTeam(id int, c *gin.Context){
	db:= database.Connection()
	Query := `DELETE FROM team WHERE id = $1;`
	_, err := db.Exec(Query, id)
	if err != nil {
		panic(err)
	} else {
		c.String(http.StatusOK, "Team Deleted")
	}
}
