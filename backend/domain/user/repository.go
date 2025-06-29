package user

type Repository interface {
	FindAll() ([]User, error)
	// Create(user *User) error
}
