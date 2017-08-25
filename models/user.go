package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	// UserList ...
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	UserList["user_11111"] = &u
}

// User ...
type User struct {
	ID       string
	Username string
	Password string
	Profile  Profile
}

// Profile ...
type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

// AddUser ...
func AddUser(u User) string {
	u.ID = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.ID] = &u
	return u.ID
}

// GetUser ...
func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

// GetAllUsers ...
func GetAllUsers() map[string]*User {
	return UserList
}

// UpdateUser ...
func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

// Login ...
func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

// DeleteUser ...
func DeleteUser(uid string) {
	delete(UserList, uid)
}
