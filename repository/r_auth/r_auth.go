package rauth

import (
	"time"

	"lms/config/db"
	"lms/model"
)

func QAuthUser(email, password string) (user model.User, err error) {

	err = db.Server().Where("email = ?", email).First(&user).Error

	return user, err

}

func UpdateLastLogin(user *model.User) error {
	now := time.Now()

	user.LastLogin = &now
	return db.Server().Save(user).Error
}
