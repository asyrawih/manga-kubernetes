package domain

type Genre string

type Manga struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Thumb         string `json:"thumb"`
	Genre         Genre  `json:"genre"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	YearPublished string `json:"year_published"`
	Status        string `json:"status"`
	CreatedBy     string `json:"created_by"`
}

type UpdateRequest struct {
	Title         string `json:"title"`
	Thumb         string `json:"thumb"`
	Genre         Genre  `json:"genre"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	YearPublished string `json:"year_published"`
	Status        string `json:"status"`
	CreatedBy     string `json:"created_by"`
}

type CreateRequest struct {
	Title         string `json:"title"`
	Thumb         string `json:"thumb"`
	Genre         Genre  `json:"genre"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	YearPublished string `json:"year_published"`
	Status        string `json:"status"`
	CreatedBy     string `json:"created_by"`
}

type Mangas []Manga
