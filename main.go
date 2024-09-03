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
                http.Error(w, "Error getting player", http.StatusInternalServerError)
                return
            }
        }
    }
}
