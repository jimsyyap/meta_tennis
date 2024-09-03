package main

import (
	"backend/repositories"
	"backend/services"
	"backend/utils"
	"log"
	"net/http"
)

func main() {
	db := utils.ConnectDB()
	playerRepository := repositories.NewPlayerRepository(db)
	playerService := services.NewPlayerService(playerRepository)

	courtRepository := repositories.NewCourtRepository(db)
	courtService := services.NewCourtService(courtRepository)

	http.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		// ...
	})

	http.HandleFunc("/courts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			court, err := courtService.GetCourt(1)
			if err != nil {
				log.Println(err)
				http.Error(w, "Error fetching court", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(court)
		} else if r.Method == "POST" {
			var court models.Court
			err := json.NewDecoder(r.Body).Decode(&court)
			if err != nil {
				log.Println(err)
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}
			err = courtService.CreateCourt(&court)
			if err != nil {
				log.Println(err)
				http.Error(w, "Error creating court", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
