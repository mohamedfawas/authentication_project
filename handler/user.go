package handlers

import (
	"fmt"
	"login-gin/models"
)

// Map for Storing Users list
var Users = map[string]models.User{}

// Check if the user exist
func GetUser(key string) (u models.User, err any) {
	u, exists := Users[key]
	if !exists {
		err = "User don't exist"
		return u, err
	}
	err = nil
	return u, err
}

// Create User function
func CreateUser(name, email, pwd string) (err any) {
	//Check if User with same mail exist
	_, exists := Users[email]
	if exists {
		fmt.Println("User already exists:", Users[email])
		return "User already exists"
	}
	//Add user to map(Users list)
	Users[email] = models.User{
		Name:     name,
		Email:    email,
		Password: pwd,
	}
	fmt.Println("User Created:", Users[email])
	fmt.Println(Users)
	return nil
}
