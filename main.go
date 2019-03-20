package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"

	db "go-tourney/persist"
	"go-tourney/utils"
	"github.com/gorilla/mux"
)

func main() {
        db.InitDb()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/players", GetPlayers).Methods("GET")
	router.HandleFunc("/players", AddPlayer).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	rows := db.QueryDatabase("SELECT id, name, knickname FROM players")
	defer rows.Close()

	players := []Player{}
	for rows.Next() {
		var p Player
		err := rows.Scan(&p.ID, &p.Name, &p.Knickname)
		if err != nil {
			panic(err.Error())
		}
		players = append(players, p)
	}

	fmt.Printf("Responding with players\n")
	utils.RespondWithJSON(w, http.StatusOK, players)
}

func AddPlayer(w http.ResponseWriter, r *http.Request) {
	var p Player
	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	if err := decoder.Decode(&p); err != nil {
		utils.RespondWithError(w, 422, "Unprocessable Entity")
	}

	statement := fmt.Sprintf("INSERT INTO players VALUES(%d, '%s', '%s')", p.ID, p.Name, p.Knickname)
	db.InsertDatabase(statement)

	utils.RespondWithJSON(w, http.StatusCreated, p)

	fmt.Printf("Player added\n")
}

type Player struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Knickname string `json:"knickname"`
}

type Players struct {
	Players []Player
}
