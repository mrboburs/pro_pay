package model

type Files struct {
	Name string `json:"fileName" db:"-"`
	Link string `json:"fileLink" db:"-"`
}
