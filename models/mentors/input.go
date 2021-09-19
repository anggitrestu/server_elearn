package mentors

// mewakili apa yang diinputkan oleh user
type AddMentorInput struct {
	Name       	string `json:"name" binding:"required"`
	Profile		string `json:"profile" binding:"required"`
	Profession 	string `json:"profession" binding:"required"`
	Email      	string `json:"email" binding:"required,email"`
}

type GetMentorInput struct {
	ID int `uri:"id" binding:"required"`
}
