package mentors

// respone yang ditampikan di api
type MentorFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Profile    string `json:"profile"`
	Profession string `json:"profession"`
	Email      string `json:"email"`
}

func FormatMentor(mentor Mentor) MentorFormatter {
	formatter := MentorFormatter {
		ID:         mentor.ID,
		Name:       mentor.Name,
		Profile: mentor.Profile,
		Profession: mentor.Profession,
		Email:      mentor.Email,
	}

	return formatter

}

func FormatMentors(mentors []Mentor) []MentorFormatter {
	mentorsFormatter := []MentorFormatter{}

	for _, mentor := range mentors {
		mentorFormatter := FormatMentor(mentor)
		mentorsFormatter = append(mentorsFormatter, mentorFormatter)
	}

	return mentorsFormatter
}