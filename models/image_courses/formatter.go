package imagecourses

type ImageCourseFormatter struct {
	ID       int    `json:"id"`
	Image    string `json:"image"`
	CourseID int    `json:"course_id"`
}

func FormatImageCourse(imageCourse ImageCourse) ImageCourseFormatter {
	formatter := ImageCourseFormatter{
		ID:       imageCourse.ID,
		Image:    imageCourse.Image,
		CourseID: imageCourse.CourseID,
	}
	return formatter
}