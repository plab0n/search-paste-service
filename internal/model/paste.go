package model

type Paste struct {
	ID    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Text  string `json:"text" db:"text"`
	*Base
}
