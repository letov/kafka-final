package domain

type Find struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Find   string `json:"find"`
}
