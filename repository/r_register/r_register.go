package rregister

import (
	"fmt"
	"lms/config/db"
	"lms/config/helper"
	"lms/model"

	"github.com/google/uuid"
)

func QRegisterUser(user *model.User) error {
	iduser, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	user.IDUser = iduser

	// Hash password before save
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// Menyimpan user ke dalam database
	if err := db.Server().Create(user).Error; err != nil {
		return err
	}

	// Menambahkan user ke dalam user_roles
	var roles []model.Role
	if err := db.Server().Where("id_role = ?", user.IDRole).Find(&roles).Error; err != nil {
		return err
	}

	if len(roles) == 0 {
		return fmt.Errorf("role not found")
	}

	if err := db.Server().Model(user).Association("Roles").Append(roles); err != nil {
		return err
	}

	return nil
}
