package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/players", GetPlayers).Methods("GET")
	router.HandleFunc("/players", AddPlayer).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(tourney-mysql:3306)/tourney_db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, name, knickname FROM players")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	players := []Player{}
	for rows.Next() {
		var p Player
		err = rows.Scan(&p.ID, &p.Name, &p.Knickname)
		if err != nil {
			panic(err.Error())
		}
		players = append(players, p)
	}

	fmt.Printf("Responding with players\n")
        respondWithJSON(w, http.StatusOK, players)
}

func AddPlayer(w http.ResponseWriter, r *http.Request) {
	var p Player
	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, 422, "Unprocessable Entity")
	}

	db, err := sql.Open("mysql", "root:password@tcp(tourney-mysql:3306)/tourney_db")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	
	statement := fmt.Sprintf("INSERT INTO players VALUES(%d, '%s', '%s')", p.ID, p.Name, p.Knickname)
	insert, err := db.Query(statement)

	if err != nil {
		panic(err)
	}

	defer insert.Close()
	
        respondWithJSON(w, http.StatusCreated, p)

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

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}
