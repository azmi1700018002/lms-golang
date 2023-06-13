package model

import (
	"time"

	"github.com/google/uuid"
)

// type Quiz struct {
// 	IDQuiz    int        `gorm:"primaryKey" json:"id_quiz"`
// 	QuizName  string     `json:"quiz_name"`
// 	QuizDesc  string     `json:"quiz_desc"`
// 	QuizType  string     `json:"quiz_type"`
// 	CreatedAt *time.Time `json:"created_at"`
// 	Questions []Question `json:"questions" gorm:"foreignKey:IDQuiz"`
// 	Sections  []Section  `gorm:"many2many:quiz_sections;"`
// }

// type Question struct {
// 	IDQuestion   int               `gorm:"primaryKey" json:"id_question"`
// 	IDQuiz       int               `json:"id_quiz"`
// 	QuestionName string            `json:"question_name"`
// 	Answers      []Answer          `json:"answers" gorm:"foreignKey:IDQuestion"`
// 	Sections     []QuestionSection `gorm:"many2many:question_sections;"`
// }

// type Answer struct {
// 	IDAnswer         int               `gorm:"primaryKey" json:"id_answer"`
// 	IDQuestion       int               `json:"id_question"`
// 	AnswerText       string            `json:"answer_text"`
// 	IsCorrect        bool              `json:"is_correct"`
// 	QuestionSections []QuestionSection `gorm:"many2many:question_sections;"`
// }

// type QuizSection struct {
// 	IDQuiz    int
// 	IDSection int
// 	Score     int
// 	DateStart *time.Time
// 	DateEnd   *time.Time
// }

// type QuestionSection struct {
// 	ID         int
// 	QuestionID int
// 	SectionID  int
// 	IDAnswer   int
// 	IsCorrect  bool
// 	CreatedAt  time.Time
// 	UpdatedAt  time.Time
// }

type Quiz struct {
	IDQuiz    int        `gorm:"primaryKey" json:"id_quiz"`
	IDUser    uuid.UUID  `gorm:"type:uuid; column:id_user;"`
	QuizName  string     `json:"quiz_name"`
	QuizDesc  string     `json:"quiz_desc"`
	QuizType  string     `json:"quiz_type"`
	CreatedAt *time.Time `json:"created_at"`
	Score     int        `json:"score"`
	DateStart *time.Time `json:"date_start"`
	DateEnd   *time.Time `json:"date_end"`
	Questions []Question `json:"questions" gorm:"foreignKey:IDQuiz"`
	Sections  []Section  `gorm:"many2many:quiz_sections;"`
}

type Question struct {
	IDQuestion   int       `gorm:"primaryKey" json:"id_question"`
	IDQuiz       int       `json:"id_quiz"`
	QuestionName string    `json:"question_name"`
	IsCorrect    bool      `JSON:"is_correct"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Answers      []Answer  `json:"answers" gorm:"foreignKey:IDQuestion"`
	Sections     []Section `gorm:"many2many:question_sections;"`
}

type Answer struct {
	IDAnswer   int    `gorm:"primaryKey" json:"id_answer"`
	IDQuestion int    `json:"id_question"`
	AnswerText string `json:"answer_text"`
	IsCorrect  bool   `json:"is_correct"`
}
