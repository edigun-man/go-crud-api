package models

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Users = []User{}
var LastId = 0

func CreateUser(name string, email string) User {
	LastId++

	newUser := User{
		LastId,
		name,
		email,
	}

	Users = append(Users, newUser)

	return newUser

}

func GetUserById(id int) *User {
	for i := range Users {
		if Users[i].Id == id {
			return &Users[i]
		}
	}

	return nil
}

func UpdateUser(id int, name string, email string) *User {
	for i, v := range Users {
		if v.Id == id {
			Users[i].Id = id
			Users[i].Name = name

			return &Users[i]
		}
	}

	return nil
}

func DeleteUser(id int) bool {
	for i := range Users {
		if Users[i].Id == id {
			Users = append(Users[:i], Users[i+1:]...)
			return true
		}
	}

	return false
}
