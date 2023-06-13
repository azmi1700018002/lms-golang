package ruser

import (
	"fmt"
	"lms/config/db"
	"lms/config/helper"
	"lms/model"
)

type UserUpdateRepository interface {
	UpdateUser(user *model.User) error
}

type userUpdateRepository struct{}

func UpdateUserRepository() UserUpdateRepository {
	return &userUpdateRepository{}
}

func (r *userUpdateRepository) UpdateUser(user *model.User) error {
	// Check if the user exists
	var existingUser model.User
	if err := db.Server().Where("id_user = ?", user.IDUser).First(&existingUser).Error; err != nil {
		return err
	}

	// Preserve existing password if the field is not provided
	if user.Password == "" {
		user.Password = existingUser.Password
	} else {
		// Hash new password
		hashedPassword, err := helper.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	// Update user data in the database
	if err := db.Server().Model(&model.User{}).
		Where("id_user = ?", user.IDUser).
		Updates(map[string]interface{}{
			"username":        user.Username,
			"email":           user.Email,
			"password":        user.Password,
			"id_role":         user.IDRole,
			"profile_picture": user.ProfilePicture,
			// Exclude IDUser from updates
		}).Error; err != nil {
		return err
	}

	// Update user_roles
	var roles []model.Role
	if err := db.Server().Where("id_role = ?", user.IDRole).Find(&roles).Error; err != nil {
		return err
	}

	if len(roles) == 0 {
		return fmt.Errorf("role not found")
	}

	if err := db.Server().Model(&existingUser).Association("Roles").Replace(roles); err != nil {
		return err
	}

	return nil
}
