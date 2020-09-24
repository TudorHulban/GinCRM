package setup

/*
File contains security riths and roles.
*/

// SecuRights Temporary storing security rights.
var SecuRights = []string{
	"Login",
	"Add user",
	"View User",
}

// SecuRoles Temporary storing security roles.
var SecuRoles = []string{
	"User Management",
	"User View Only",
}

// RolesDefinition Temporary storing application security roles definition.
var RolesDefinition = [][]uint8{
	[]uint8{0, 1, 2},
	[]uint8{0, 2},
}

// SecuProfiles Temporary storing security profiles.
var SecuProfiles = []string{
	"Administrator",
	"Guest",
}

// ProfilesDefinition Temporary storing application security profiles definition.
var ProfilesDefinition = [][]uint8{
	[]uint8{0},
	[]uint8{1},
}
