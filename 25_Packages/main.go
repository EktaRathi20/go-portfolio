package main

import (
	"fmt"
	"golang/25_Packages/auth"
	"golang/25_Packages/user"
)

/*
* Help to re-use code by importing packages
 */
func main() {
	auth.LoginWithCredentials("admin", "password123")
	session := auth.GetSession()
	println("Session:", session)

	user := user.User{
		Email: "one@example.com",
	}

	fmt.Println("User Email:", user.Email)
}
