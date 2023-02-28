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

type Chapters []Chapter

type CreateChapterRequest struct {
	MangaId       string  `json:"manga_id,omitempty"`
	Title         string  `json:"title,omitempty"`
	ChapterNumber int64   `json:"chapter_number,omitempty"`
	Images        Content `json:"images,omitempty"`
}

type UpdateChapterRequest struct {
	Id            string    `json:"id,omitempty"`
	MangaId       string    `json:"manga_id,omitempty"`
	ChapterNumber int64     `json:"chapter_number,omitempty"`
	Title         string    `json:"title,omitempty"`
	Images        []Content `json:"images,omitempty"`
}
