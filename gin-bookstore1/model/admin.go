package model

type Admin struct {
	AdminID       int    `json:"adminID"`
	AdminName     string `json:"adminName"`
	AdminPassword string `json:"adminPassword"`
	//State int
}
