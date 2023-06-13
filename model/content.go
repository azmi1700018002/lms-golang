package model

type Content struct {
	IDcontent    int    `json:"idcontent" gorm:"primaryKey"`
	IDsection    int    `json:"idsection"`
	ContentTitle string `json:"content_title"`
	ContentType  int    `json:"content_type"`
	Image        string `json:"image"`
	Link         string `json:"link"`
}
