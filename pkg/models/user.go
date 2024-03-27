package models

type User struct {
	Id        int    `json:"id"`
	Pass_hash string `json:"pass"`
	Username  string `json:"user"`
}
