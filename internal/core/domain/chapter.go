package domain

type Content []byte

// Chapter struct
type Chapter struct {
	ID            string  `json:"id,omitempty"`
	MangaID       string  `json:"manga_id,omitempty"`
	ChapterNumber int64   `json:"chapter_number,omitempty"`
	Title         string  `json:"title,omitempty"`
	Images        Content `json:"images,omitempty"`
}

type ChapterResponse struct {
	ID            string `json:"id,omitempty"`
	MangaID       string `json:"manga_id,omitempty"`
	ChapterNumber int64  `json:"chapter_number,omitempty"`
	Title         string `json:"title,omitempty"`
	Images        string `json:"images,omitempty"`
}

type Chapters []ChapterResponse

type CreateChapterRequest struct {
	MangaID       string  `json:"manga_id,omitempty"`
	Title         string  `json:"title,omitempty"`
	ChapterNumber int64   `json:"chapter_number,omitempty"`
	Images        Content `json:"images,omitempty"`
}

type UpdateChapterRequest struct {
	ID            string    `json:"id,omitempty"`
	MangaID       string    `json:"manga_id,omitempty"`
	ChapterNumber int64     `json:"chapter_number,omitempty"`
	Title         string    `json:"title,omitempty"`
	Images        []Content `json:"images,omitempty"`
}
