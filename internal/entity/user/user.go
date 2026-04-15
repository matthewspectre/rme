package user

// Domain entity for application users.

type User struct {
	ID       int
	Username string
	Password string
	FullName string
	Role     int
}
