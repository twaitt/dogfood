package main

import "fmt"

// effort to add a feature to the product
type effort struct {
	user
	body string
}

// idea to make the product better
type idea struct {
	user
	body string
}

// user of the product
type user struct {
	id       int64
	fullName string
	username string
	email    string
}

// createUser saves a user to the database
func createUser() {
}

func main() {

}