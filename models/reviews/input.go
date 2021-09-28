package reviews

type CreateReviewInput struct {
	CourseID int    `json:"course_id"`
	Rating   int    `json:"rating"`
	Note     string `json:"note"`
}

type GetReviewInput struct {
	ID int `uri:"id" binding:"required"`
}
