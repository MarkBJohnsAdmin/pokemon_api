package main

import (
	"log"
	"net/http"
	"server/database/upload"
	"server/servers"
)

func main() {

	db, err := upload.InitDatabase()
	if err != nil {
		log.Fatalf("failed to initiate database")
	}

	pkmnRouter := servers.SetUpPokemonRouter(db)
	pdexRouter := servers.SetUpPokedexRouter(db)
	abltyRouter := servers.SetUpAbilityRouter(db)
	moveRouter := servers.SetUpMoveRouter(db)
	testRouter := servers.SetUpTestRouter(db)

	mux := http.NewServeMux()
	mux.Handle("/pokemon/", pkmnRouter)
	mux.Handle("/pokedex/", pdexRouter)
	mux.Handle("/ability/", abltyRouter)
	mux.Handle("/move/", moveRouter)
	mux.Handle("/test/", testRouter)

	port := ":8080"
	log.Printf("starting server on port %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
