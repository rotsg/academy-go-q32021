package model

type Song struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
}
