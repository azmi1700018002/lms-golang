package model

type Knowledge struct {
	IdKnowledge   int      `json:"id_knowledge" gorm:"primarykey"`
	IDcategory    int      `json:"idcategory"`
	KnowledgeName string   `json:"knowledge_name"`
	Description   string   `json:"description"`
	Status        int      `json:"status"`
	Courses       []Course `json:"courses" gorm:"foreignKey:IdKnowledge"`
}
