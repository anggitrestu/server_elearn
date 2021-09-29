package reviews

type ReviewFormatter struct {
	ID       int    `json:"id"`
	CourseID int    `json:"course_id"`
	UserID   int    `json:"user_id"`
	Rating   int    `json:"rating"`
	Note     string `json:"note"`
}

func FormatReview(review Review) ReviewFormatter {
	formatter := ReviewFormatter{
		ID:       review.ID,
		CourseID: review.CourseID,
		UserID:   review.UserID,
		Rating:   review.Rating,
		Note:     review.Note,
	}
	return formatter
}