package users

// User is a type for users on our platform
type User struct {
	Name  string
	Email string

	// Can only be used inside of this file
	// FIELD LEVEL ENCAPSULATION!
	password string
}
