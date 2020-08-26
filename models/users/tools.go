package users

import "golang.org/x/crypto/bcrypt"

func ifExists(compare string, trueResponse string, falseResponse string) string {
	if len(compare) > 0 {
		return trueResponse
	} else {
		return falseResponse
	}
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func removePasswordOnScope(users []User) []User {
	for index, _ := range users {
		users[index].Password = ""
	}
	return users
}
