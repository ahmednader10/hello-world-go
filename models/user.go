package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User //creates a slice with pointers for objects of type user
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user should not have an id")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("user with id `%v` not found", id)
}

func UpdateUser(user User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == user.ID {
			users[i] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user with id `%v` not found", user.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with id `%v` not found", id)
}
