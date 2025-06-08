package auth

// should create function which start with captial letter
func LoginWithCredentials(username, password string) bool {
	// Simulate a login check
	if username == "admin" && password == "password123" {
		return true
	}
	return false
}
