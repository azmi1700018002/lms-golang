package model

type Menu struct {
	IDMenu int `json:"id_menu" gorm:"primaryKey"`
	IDRole int `json:"id_role"`
}
