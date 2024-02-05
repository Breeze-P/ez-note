package pack

import (
	"ez-note/kitex_gen/userdemo"

	"ez-note/cmd/user/dal/db"
)

// User pack user info
func User(u *db.User) *userdemo.User {
	if u == nil {
		return nil
	}

	return &userdemo.User{UserId: int64(u.ID), UserName: u.UserName, Avatar: "test"}
}

// Users pack list of user info
func Users(us []*db.User) []*userdemo.User {
	users := make([]*userdemo.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
