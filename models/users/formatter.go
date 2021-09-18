package users

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Profession string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Profession: user.Profession,
		Email:      user.Email,
		Token:      token,
	}

	return formatter

}