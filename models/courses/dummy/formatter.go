package dummy

// respone yang ditampikan di api
type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Profession string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(dummy Dummy, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         dummy.ID,
		Name:       dummy.Name,
		Profession: dummy.Profession,
		Email:      dummy.Email,
		Token:      token,
	}

	return formatter

}