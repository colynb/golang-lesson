package user

var highestID int

type User struct {
	ID    int
	Email string
	Name  string
}

func New(u User) User {
	highestID++
	u.ID = highestID
	return u
}

