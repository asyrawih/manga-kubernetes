package domain

type Genre string

type Manga struct {
	Title         string `json:"title"`
	Thumb         string `json:"thumb"`
	Genre         Genre  `json:"genre"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	YearPublished string `json:"year_published"`
	Status        string `json:"status"`
}

type Mangas []Manga
