package persistence

/*
File contains structures needed for user management.
*/

// UserAuth Consolidates data referring to user authentication.
type UserAuth struct {
	ID              uint64 `gorm:"primaryKey"` // user ID, primary key
	SecurityGroupID uint8  // security group ID
	CreatedAt       uint64 `gorm:"autoCreateTime"` // UNIX time for creation time
	LastLoginAt     uint64 // UNIX time for last login
	LastUpdateAt    uint64 // UNIX time for last update
	UserCode        string // user code, used for login
	PasswordSALT    string // should not be sent in JSON, exported for ORM
	PasswordHASH    string // should not be sent in JSON, exported for ORM

}

// Contact Is used when defining a app user.
// The app user should have at least one contact.
type Contact struct {
	ID             uint64 // contact ID, primary key
	UserID         uint64 `pg:"userid"`
	FirstName      string `json:"firstname"`
	LastName       string `json:"lasttname"`
	OfficePhoneNo  string
	MobilePhoneNo  string
	CompanyEmail   string
	WorkEmail      string
	AddressHQ      string
	AddressOffice  string
	AddressBilling string
}

// UserData Consolidates user information for creating a user in persistance layer.
type UserData struct {
	UserAuth
	Contacts []Contact
}
