package model

type University struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Country string   `json:"country"`
	WebPage []string `json:"web_pages"`
}
