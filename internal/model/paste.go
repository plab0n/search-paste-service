package model

type Paste struct {
	ID    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Text  string `json:"text" db:"text"`
	*Base
}
type ScrapingInfo struct {
	PasteId int    `json:"paste_id"`
	Url     string `json:"url"`
}

type EmbeddingPayload struct {
	PasteId int
	Text    string
}
