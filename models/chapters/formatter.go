package chapters

type ChapterFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CourseID int    `json:"course_id"`
}

func FormatChapter(chapter Chapter) ChapterFormatter {
	formatter := ChapterFormatter{
		ID:       chapter.ID,
		Name:     chapter.Name,
		CourseID: chapter.CourseID,
	}
	return formatter
}

func FormatChapters(chapters []Chapter) []ChapterFormatter {
	chaptersFormatter := []ChapterFormatter{}

	for _, chapter := range chapters {
		chapterFormatter := FormatChapter(chapter)
		chaptersFormatter = append(chaptersFormatter, chapterFormatter)
	}

	return chaptersFormatter
}