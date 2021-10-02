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


type chapter  struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CourseID  int            `json:"course_id"`
	Lessons []lesson 			 `json:"lessons"`
} 

type lesson struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Video     string         `son:"video"`
	ChapterID int            `json:"chapter_id"`
}

type mentor struct {
	ID             	int 			`json:"id"`
	Name           	string			`json:"name"`
	Profile        	string			`json:"profile"`
	Email          	string			`json:"email"`
	Profession 		string			`json:"profession"`
}

type imageCourse struct {
	ID        int            `json:"id"`
	CourseID  int            `json:"course_id"`
	Image 	  string	  	 `json:"image"`
}

type DetailCourseFormatter struct {
	ID			int	   `json:"id"`
	Name        string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Certificate bool   `json:"certificate"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Level       string `json:"level" `
	Description string `json:"description"`
	Chapters    []chapter    `json:"chapters"`
	Mentor 	mentor `json:"mentors"`
	ImagesCourses []imageCourse `json:"image_courses"`
}

// func FormatDetailCourse(course Course) DetailCourseFormatter {
// 	formatter := DetailCourseFormatter {
// 		ID: course.ID,
// 		Name: course.Name,
// 		Thumbnail: course.Thumbnail,
// 		Certificate: course.Certificate,
// 		Type: course.Type,
// 		Status: course.Status,
// 		Level: course.Level,
// 		Description: course.Description,
// 		Mentor: mentor{
// 			ID: course.Mentor.ID,
// 			Name: course.Mentor.Name,
// 			Profile: course.Mentor.Profile,
// 			Email: course.Mentor.Email,
// 			Profession: course.Mentor.Profession,
// 		},
// 		Chapters: []chapter{
// 			chapter{
// 				ID: course.Chapters[0].ID,
// 				Name:  course.Chapters[0].Name,
// 				CourseID: course.Chapters[0].CourseID,
// 				Lessons: []lesson{
// 					lesson{
// 						ID: course.Chapters[0].Lessons[0].ID,
// 						Name: course.Chapters[0].Lessons[0].Name,
// 						Video: course.Chapters[0].Lessons[0].Video,
// 						ChapterID: course.Chapters[0].Lessons[0].ChapterID,
// 					},
// 				},
// 			},
// 		},
// 	}

// 	return formatter
// }
