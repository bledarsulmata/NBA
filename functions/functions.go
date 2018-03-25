package functions

import (
	"github.com/gin-gonic/gin"
	"strconv"
	_ "github.com/lib/pq"
	"net/http"
	"NBA/bledar/NBA/structs"
	"NBA/bledar/NBA/database"
	"NBA/bledar/NBA/services"
)

var teams []structs.Team

var players []structs.Player

var PL []structs.PointLeaders

var TR []structs.TeamRank


func CheckLogin(c *gin.Context){
	email := c.PostForm("email")
	password := c.PostForm("password")
	db:= database.Connection()
	Query := "Select username from perdorues where email=$1 and password=$2"
	username, err := db.Query(Query, email, password)
	if err != nil {
		c.String(http.StatusOK,"Error ne query")
	}
	var u string
	defer username.Close()
	username.Next()
	err = username.Scan(&u)
	if err != nil {
		panic(err)
	}
	data, err := db.Exec(Query, email, password)
	if err != nil {
		panic(err)
	}
	rows, err := data.RowsAffected()
	if rows == 0 {
		c.String(http.StatusOK,"Wrong Email or Password")
	}else {
		c.String(http.StatusOK,"Welcome %s", u)
	}

}

func Register(c *gin.Context){
	db:= database.Connection()
	email := c.PostForm("email")
	password := c.PostForm("password")
	username := c.PostForm("username")
	if email == "" || password == "" || username == "" {
		c.String(http.StatusOK, "Plotesoni te gjitha te dhenat!")
	}else {
		checkemail := services.EmailValidation(email)
		if checkemail == false {
			c.String(http.StatusOK, "Put an valid format of email")
		} else {
			data, err := db.Exec("SELECT * FROM perdorues WHERE username=$1", username)
			if err != nil {
				panic(err)
			}

			rows, err := data.RowsAffected()
			if err != nil {
				panic(err)
			}
			if rows == 1 {
				c.String(http.StatusOK, "Na vjen keq ai username eshte i zene")
			} else {
				_, err := db.Exec(`INSERT INTO perdorues (email, password, username) VALUES ($1,$2,$3)`, email, password, username)
				if err != nil {
					panic(err)
				}
				c.String(http.StatusOK, "Ju jeni regjistruar me sukses tani mund te logoheni")
			}
		}
	}
}

	func GetPlayers(c *gin.Context) {
		db:= database.Connection()
		rows, err := db.Query("Select * FROM player")
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
			players = append(players, player)
		}
		err = rows.Err()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, players)
	}

	func GetPlayer(c * gin.Context) {
		ID := c.Param("id")
		I, _ := strconv.Atoi(ID)
		services.SgetPlayer(I, c)
	}

	func MostSteals(c *gin.Context){
		year := c.Param("year")
		db:= database.Connection()
		rows, err := db.Query("Select pname,surname,steals FROM player_history JOIN player ON year=$1 AND player.id=player_history.player_id ORDER BY steals DESC", year)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var ps structs.PointLeaders
			err = rows.Scan(&ps.PlayerName,&ps.PlayerSurname, &ps.Total)
			if err != nil {
				panic(err)
			}
			PL = append(PL, ps)
		}
		err = rows.Err()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, PL)

	}

func MostBlocks(c *gin.Context){
	year := c.Param("year")
	db:= database.Connection()
	rows, err := db.Query("Select pname,surname,blocks FROM player_history JOIN player ON year=$1 AND player.id=player_history.player_id ORDER BY blocks DESC", year)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var ps structs.PointLeaders
		err = rows.Scan(&ps.PlayerName,&ps.PlayerSurname, &ps.Total)
		if err != nil {
			panic(err)
		}
		PL = append(PL, ps)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, PL)

}

func MostRebounds(c *gin.Context){
	year := c.Param("year")
	db:= database.Connection()
	rows, err := db.Query("Select pname,surname,rebounds FROM player_history JOIN player ON year=$1 AND player.id=player_history.player_id ORDER BY rebounds DESC", year)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var ps structs.PointLeaders
		err = rows.Scan(&ps.PlayerName,&ps.PlayerSurname, &ps.Total)
		if err != nil {
			panic(err)
		}
		PL = append(PL, ps)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, PL)

}

func MostAssists(c *gin.Context){
	year := c.Param("year")
	db:= database.Connection()
	rows, err := db.Query("Select pname,surname,assists FROM player_history JOIN player ON year=$1 AND player.id=player_history.player_id ORDER BY assists DESC", year)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var ps structs.PointLeaders
		err = rows.Scan(&ps.PlayerName,&ps.PlayerSurname, &ps.Total)
		if err != nil {
			panic(err)
		}
		PL = append(PL, ps)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, PL)

}

func TeamRank(c *gin.Context){
	year := c.Param("year")
	db:= database.Connection()
	rows, err := db.Query("SELECT tname, rank FROM team_history JOIN team ON year=$1 AND team.id=team_history.team_id ORDER BY rank ASC", year)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var team structs.TeamRank
		err = rows.Scan(&team.Name, &team.Rank)
		if err != nil {
			panic(err)
		}
		TR = append(TR, team)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, TR)
}

func MostPoints(c *gin.Context){
	year := c.Param("year")
	db:= database.Connection()
	rows, err := db.Query("Select pname,surname,points FROM player_history JOIN player ON year=$1 AND player.id=player_history.player_id ORDER BY points DESC", year)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var ps structs.PointLeaders
		err = rows.Scan(&ps.PlayerName,&ps.PlayerSurname, &ps.Total)
		if err != nil {
			panic(err)
		}
		PL = append(PL, ps)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, PL)

}

	func AddPlayer(c * gin.Context) {
		Name := c.PostForm("name")
		Surname := c.PostForm("surname")
		AGE := c.PostForm("age")
		age, _ := strconv.Atoi(AGE)
		TEAMID := c.PostForm("teamid")
		T, _ := strconv.Atoi(TEAMID)
		services.SaddPlayer(Name, Surname, age, T, c)
	}

	func UpdatePlayer(c * gin.Context){
		ID := c.Param("id")
		I, _ := strconv.Atoi(ID)
		teamid := c.PostForm("teamid")
		T, _ := strconv.Atoi(teamid)
		services.SupdatePlayer(I, T, c)

	}

	func DeletePlayer(c * gin.Context) {
		ID := c.Param("id")
		I, _ := strconv.Atoi(ID)
		services.SdeletePlayer(I, c)
	}

	func GetTeams(c * gin.Context) {
		db:= database.Connection()
		rows, err := db.Query("Select * FROM team")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var team structs.Team
			err = rows.Scan(&team.Id, &team.Name)
			if err != nil {
				panic(err)
			}
			teams = append(teams, team)
		}
		err = rows.Err()
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, teams)
	}

	func GetTeam(c * gin.Context) {
		ID := c.Param("id")
		I, _ := strconv.Atoi(ID)
		services.SgetTeam(I, c)
	}

	func  GetTeamPlayers(c * gin.Context) {
		ID := c.Param("id")
		I, _ := strconv.Atoi(ID)
		services.SgetTeamPlayers(I, c)
	}

	func AddTeam(c *gin.Context) {
		ID := c.PostForm("id")
		I, _ := strconv.Atoi(ID)
		Name := c.PostForm("name")
		services.SaddTeam(I, Name, c)
	}

	func DeleteTeam(c *gin.Context) {
		ID := c.Param("id")
		I, _ := strconv.Atoi(ID)
		services.SdeleteTeam(I, c)
	}

