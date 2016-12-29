package Data

import "log"

//Users list of users
var Users []User

//InitUsers int users
func InitUsers() []User {
	Users = []User{User{"Martin", "Ladeveze", 4, "martin", "KatchaKatcha"}, User{"Antoine", "Ladeveze", 2, "toinou", "cacaboudin"}}
	return Users
}

//StoreUser store a user (no kidding)
func StoreUser(user User) {
	Users = append(Users, user)
}

//GetUser get the fucking user from storage
func GetUser(userID int) User {
	if userID > len(Users) || userID < 0 {
		log.Panic(0, "Invalid index")
	}
	return Users[userID]
}

//GetUsers get all the fucking users from storage
func GetUsers() []User {
	return Users
}

//CreateUser add a new user and get the list
func CreateUser(user User) User {
	Users = append(Users, user)
	return Users[len(Users)-1]
}

//LoginUser find which user match credentials
func LoginUser(login, pwd string) User {
	for _, user := range Users {
		log.Output(0, login+" "+pwd)
		if user.Login == login && user.Password == pwd {
			return user
		}
	}
	return User{"", "", 1, "", ""} //panic("User non trouvé")
}
