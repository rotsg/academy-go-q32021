package model

type University struct {
	Domaing       []string `json:"domains"`
	Name          string   `json:"name"`
	Country       string   `json:"country"`
	AlphaTwoCode  string   `json:"alpha_two_code"`
	StateProvince string   `json:"state-province"`
	WebPages      []string `json:"web_pages"`
}
