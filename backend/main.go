package main

import (
	"backend/services"
	"backend/utils"
	"log"
	"net/http"
)

func main() {
	db := utils.ConnectDB()
	playerRepository := repositories.NewPlayerRepository(db)
	playerService := services.NewPlayerService(playerRepository)

	http.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			player, err := playerService.GetPlayer(1)
			if err != nil {
				log.Println(err)
				http.Error(w, "Error fetching player", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(player)
		} else if r.Method == "POST" {
			var player models.Player
			err := json.NewDecoder(r.Body).Decode(&player)
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}
			err = playerService.CreatePlayer(&player)
			if err != nil {
				log.Println(err)
				http.Error(w, "Error creating player", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

