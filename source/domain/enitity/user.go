package enitity

type User struct {
	// The user identity.
	Id string
	// The username.
	Username string
	// The user password.
	Password string
	// Whether user is deleted or not.
	IsDeleted bool
}
