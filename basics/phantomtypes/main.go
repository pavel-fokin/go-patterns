package main

import "fmt"

// Roles.
type Basic struct{}
type Admin struct{}

// User.
type User[T Basic | Admin] struct {
	name string
}

// createUser is only available for admins.
func createUser[T User[Admin]](user T) {
	fmt.Printf("%s creates a user\n", user)
}

// viewUser is available for all users.
func viewUser[T User[Basic] | User[Admin]](user T) {
	fmt.Printf("%s views a user\n", user)
}

func main() {
	basicUser := User[Basic]{"Bob"}
	adminUser := User[Admin]{"Jonh"}

	viewUser(basicUser)
	viewUser(adminUser)
	createUser(adminUser)
}
