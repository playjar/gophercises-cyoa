package model

type Arc struct {
	Name string
	Title string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []ArcOption `json:"options"`
}

type ArcOption struct {
	Name string `json:"arc"`
	Description string `json:"text"`
}