package user

type User struct {
	Email string
}

func New(email string) User {
	return User{email}
}
