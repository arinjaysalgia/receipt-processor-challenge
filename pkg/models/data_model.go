package models

type Receipt struct {
	ID     string `json:"id"`
	Points int64  `json:"points"`
}

var Receipts = map[string]Receipt{}
