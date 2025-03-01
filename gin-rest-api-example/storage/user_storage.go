package storage

import (
	data_types "example.com/gin-rest-api-example/data_types" // Import the package
)

var users = []data_types.User{
	{ID: "1", Name: "Alice"},
	{ID: "2", Name: "Bob"},
}

func GetUsers() []data_types.User {
	return users
}

func GetUserByID(id string) (data_types.User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}
	return data_types.User{}, false
}

func AddUser(user data_types.User) {
	users = append(users, user)
}

func UpdateUser(id string, updatedUser data_types.User) bool {
	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			return true
		}
	}
	return false
}

func DeleteUser(id string) bool {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
