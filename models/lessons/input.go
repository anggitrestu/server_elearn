package lessons

type CreateLessonInput struct {
	Name      string `json:"name" binding:"required"`
	Video     string `json:"video" binding:"required"`
	ChapterID int    `json:"chapter_id" binding:"required"`
}

type GetLessonInput struct {
	ID int `uri:"id" binding:"required"`
}