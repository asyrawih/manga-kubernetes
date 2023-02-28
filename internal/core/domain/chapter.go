package domain

type Content []byte

// Chapter struct
type Chapter struct {
	Id            string  `json:"id,omitempty"`
	MangaId       string  `json:"manga_id,omitempty"`
	ChapterNumber int64   `json:"chapter_number,omitempty"`
	Title         string  `json:"title,omitempty"`
	Images        Content `json:"images,omitempty"`
}

type ChapterResponse struct {
	Id            string `json:"id,omitempty"`
	MangaId       string `json:"manga_id,omitempty"`
	ChapterNumber int64  `json:"chapter_number,omitempty"`
	Title         string `json:"title,omitempty"`
	Images        string `json:"images,omitempty"`
}

type Chapters []ChapterResponse

type CreateChapterRequest struct {
	MangaId       string   `json:"manga_id,omitempty"`
	Title         string   `json:"title,omitempty"`
	ChapterNumber int64    `json:"chapter_number,omitempty"`
	Images        []string `json:"images,omitempty"`
}

type UpdateChapterRequest struct {
	Id            string   `json:"id,omitempty"`
	MangaId       string   `json:"manga_id,omitempty"`
	ChapterNumber int64    `json:"chapter_number,omitempty"`
	Title         string   `json:"title,omitempty"`
	Images        []string `json:"images,omitempty"`
}
