package lessons

type LessonFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ChapterID int    `json:"chapter_id"`
}

func FormatLesson(lesson Lesson) LessonFormatter {
	formatter := LessonFormatter{
		ID:        lesson.ID,
		Name:      lesson.Name,
		ChapterID: lesson.ChapterID,
	}
	return formatter
}

func FormatLessons(lessons []Lesson) []LessonFormatter {
	lessonsFormatter := []LessonFormatter{}

	for _, lesson := range lessons {
		lessonFormatter := FormatLesson(lesson)
		lessonsFormatter = append(lessonsFormatter, lessonFormatter)
	}

	return lessonsFormatter
}