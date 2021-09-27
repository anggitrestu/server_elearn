package imagecourses

type CreateImageCourseInput struct {
	Image    string `json:"image"`
	CourseID int    `json:"course_id"`
}

type GetImageCourseInput struct {
	ID int `uri:"id"`
}