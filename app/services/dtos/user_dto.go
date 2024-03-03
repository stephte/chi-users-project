package dtos


type UserInDTO struct {
	FirstName 	string		`json:"firstName"`
	LastName	string		`json:"lastName"`
	Email		string		`json:"email"`
	Role		int64		`json:"role" enum:"UserRole"`
}

type UserOutDTO struct {
	BaseDTO
	UserInDTO
}

// How to get password so we can retrieve it, but not send it??
type CreateUserDTO struct {
	UserInDTO
	Password	string
}
