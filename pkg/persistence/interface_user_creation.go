package persistence

// IUserCreation Methods used in user creation
type IUserCreation interface {
	// Create User
	AddUser() error
}

// IUserCRUD Full specifications for a user creation process
type IUserCRUD interface {
	IUserCreation
	// Get User Information
	// Update User
	// Delete USer
}
