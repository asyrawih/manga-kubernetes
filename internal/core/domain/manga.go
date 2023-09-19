package domain

type Genre string

type Manga struct {
	ID            string `json:"id"`
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
	Title         string `json:"title"          validate:"required"`
	Thumb         string `json:"thumb"          validate:"required"`
	Genre         Genre  `json:"genre"          validate:"required"`
	Author        string `json:"author"         validate:"required"`
	Publisher     string `json:"publisher"      validate:"required"`
	YearPublished string `json:"year_published" validate:"required"`
	Status        string `json:"status"         validate:"required"`
	CreatedBy     string `json:"created_by"     validate:"required"`
}

type Mangas []Manga
