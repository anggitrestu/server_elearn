package courses

type CreateCourseInput struct {
	Name        string `json:"name" `
	Certificate bool   `json:"certificate"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Level       string `json:"level" `
	Description string `json:"description"`
	MentorID    int    `json:"mentor_id"`
}

type GetCourseInput struct {
	ID int `uri:"id" binding:"required"`
}
