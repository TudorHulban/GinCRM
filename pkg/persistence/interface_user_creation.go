package persistence

// IUserCreation Methods used in user creation
type IUserCreation interface {
	// Create User, passing pointer to return created ID
	AddUser(*UserData) error
}

// IUserCRUD Full specifications for a user creation process
type IUserCRUD interface {
	IUserCreation
	// Get User Information
	// Update User
	// Delete USer
}
