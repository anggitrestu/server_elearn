package users

// mewakili apa yang diinputkan oleh user
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Profession string `json:"profession" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email"  binding:"required,email"`
}

// domain