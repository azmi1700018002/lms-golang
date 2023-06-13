package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	IDUser         uuid.UUID      `gorm:"primary_key; unique; type:uuid; column:id_user; default:uuid_generate_v4()"`
	IDRole         int            `json:"id_role"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	Password       string         `json:"password"`
	ProfilePicture string         `json:"profile_picture"`
	LastLogin      *time.Time     `json:"last_login"`
	Roles          []Role         `gorm:"many2many:user_roles;"`
	Course         []Course       `gorm:"foreignKey:IDUser;"`
	Quizzes        []Quiz         `gorm:"foreignKey:IDUser;"`
}

type Role struct {
	IDRole   int    `json:"id_role" gorm:"primaryKey"`
	Rolename string `json:"rolename"`
	User     []User `gorm:"foreignKey:IDRole;"`
	Users    []User `gorm:"many2many:user_roles;"`
	Menus    []Menu `json:"menus" gorm:"foreignKey:IDRole"`
}
