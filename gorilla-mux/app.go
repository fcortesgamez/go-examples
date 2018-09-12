// app.go

package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbName string) {

	log.Printf("Setting up database connection. User=%s, DB=%s", user, dbName)

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbName)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Setting up router...")

	a.Router = mux.NewRouter()

	log.Printf("Setup finished")
}

func (a *App) Run(addr string) {}
