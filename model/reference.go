package model

type Reference struct {
	IDref1    int    `json:"idref1" gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Desc1     string `json:"desc1"`
	Groupref1 int    `json:"groupref1"`
	IDref2    int    `json:"idref2" gorm:"default:null"`
	Desc2     string `json:"desc2" gorm:"default:null"`
	Groupref2 int    `json:"groupref2" gorm:"default:null"`
}
