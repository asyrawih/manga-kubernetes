package domain

type Content string

type Chapter struct {
	Title  string    `json:"title"`
	Images []Content `json:"images"`
}
