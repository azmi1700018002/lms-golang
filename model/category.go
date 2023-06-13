package model

type Category struct {
	IDcategory   int         `json:"idcategory" gorm:"primarykey"`
	Categoryname string      `json:"categoryname"`
	Knowledges   []Knowledge `json:"knowledges" gorm:"foreignKey:IDcategory"`
}
