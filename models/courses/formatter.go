package courses

// respons yang di tampilkan di api

type CourseFormatter struct {
	ID			int	   `json:"id"`
	Name        string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Certificate bool   `json:"certificate"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Level       string `json:"level" `
	Description string `json:"description"`
	MentorID    int    `json:"mentor_id"`
}

func FormatCourse(course Course) CourseFormatter {

	formatter := CourseFormatter {
		ID: course.ID,
		Name: course.Name,
		Thumbnail: course.Thumbnail,
		Certificate: course.Certificate,
		Type: course.Type,
		Status: course.Status,
		Level: course.Level,
		Description: course.Description,
		MentorID: course.MentorID,
	}

	return formatter
	 	
}

func FormatCourses(courses []Course) []CourseFormatter {

	coursesFormatter := []CourseFormatter{}

	for _,course := range courses {
		courseFormatter := FormatCourse(course)
		coursesFormatter = append(coursesFormatter, courseFormatter)
	}

	return coursesFormatter
}