package security

// User is the authenticated principal extracted from the JWT token
type User struct {
	Username      string
	Roles         []string
	Email         string
	UserID        string
	DisplayName   string
	Authenticated bool
}

// Claim defines the authorization requiremenets
type Claim struct {
	// Name of the applicatiion
	Name string
	// URL of the application
	URL string
	// Roles possible roles
	Roles []string
}
