package auth

// private functions are not accessible outside the package
func extractSession() string {
	// Simulate session extraction logic
	return "session123"
}

// public functions are accessible outside the package
func GetSession() string {
	// Simulate a session retrieval
	return extractSession()
}
