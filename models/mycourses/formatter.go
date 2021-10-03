package mycourses

type MyCourseFormatter struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	CourseID int `json:"course_id"`
}

type MyCoursesFormatter struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	CourseID int    `json:"course_id"`
	Courses  course `json:"courses"`
}

type course struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Certificate bool   `json:"certificate"`
	Thumbnail   string `json:"thumbnail"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Price       int    `json:"price"`
	Level       string `json:"level"`
	Description string `json:"description"`
}

func FormatMyCourse(mycourse MyCourse) MyCourseFormatter {
	formatter := MyCourseFormatter{
		ID:       mycourse.ID,
		UserID:   mycourse.UserID,
		CourseID: mycourse.UserID,
	}

	return formatter
}

func FormatMyCourses(mycourses MyCourse) MyCoursesFormatter {
	formatter := MyCoursesFormatter{
		ID:       mycourses.ID,
		UserID:   mycourses.UserID,
		CourseID: mycourses.CourseID,
		Courses: course{
			ID:          mycourses.Course.ID,
			Name:        mycourses.Course.Name,
			Certificate: mycourses.Course.Certificate,
			Thumbnail:   mycourses.Course.Thumbnail,
			Type:        mycourses.Course.Type,
			Status:      mycourses.Course.Status,
			Price:       mycourses.Course.Price,
			Level:       mycourses.Course.Level,
			Description: mycourses.Course.Description,
		},
	}

	return formatter
}

func FormatMyAllCourses(mycourses []MyCourse) []MyCoursesFormatter {
	if len(mycourses) == 0 {
		return []MyCoursesFormatter{}
	}

	var mycoursesFormatter []MyCoursesFormatter

	for _, mycourse := range mycourses {
		formatter := FormatMyCourses(mycourse)
		mycoursesFormatter = append(mycoursesFormatter, formatter)
	}

	return mycoursesFormatter

}