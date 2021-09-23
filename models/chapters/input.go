package chapters

type CreateChapterInput struct {
	Name     string `json:"name"  binding:"required"`
	CourseID int    `json:"course_id" binding:"required"`
}

type UpdateChapterInput struct {
	Name     string `json:"name"  binding:"required"`
	CourseID int    `json:"course_id"`
}

type GetChapterInput struct {
	ID int `uri:"id" binding:"required"`
}
