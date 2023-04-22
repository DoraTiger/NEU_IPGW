package data

type LoginMessage struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Redirect string `json:"Redirect"`
	ID       string `json:"ID"`
}
