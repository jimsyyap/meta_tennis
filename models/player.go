package models

type Player struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Experience   string `json:"experience"`
	PlayingStyle string `json:"playing_style"`
}
