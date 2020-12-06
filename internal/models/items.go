package models


type Items struct {
	ID        int    `json:"id"`
	Available int    `json:"available"`
	Model     string `json:"model"`
	Size      string `json:"size"`
}
