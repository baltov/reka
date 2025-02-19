package models

// Category represents the categories table in the database
type Category struct {
	ID        int64      `gorm:"primaryKey;column:id" json:"id"`
	Filename  string     `gorm:"type:varchar(255);column:filename" json:"filename"`
	Name      string     `gorm:"type:varchar(255);column:name" json:"name"`
	Row       int        `gorm:"not null;column:row" json:"row"`
	Questions []Question `gorm:"foreignKey:CategoryID" json:"questions"`
}

// Question represents the questions table in the database
type Question struct {
	ID         int64    `gorm:"primaryKey;column:id" json:"id"`
	Questioned int      `gorm:"not null;column:questioned" json:"questioned"`
	Row        int      `gorm:"not null;column:row" json:"row"`
	Text       string   `gorm:"type:varchar(255);column:text" json:"text"`
	CategoryID int64    `gorm:"column:category_id" json:"-"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"-"`
	Answers    []Answer `gorm:"foreignKey:QuestionID" json:"answers"`
	Images     []Image  `gorm:"foreignKey:QuestionID" json:"images"`
}

// Answer represents the answers table in the database
type Answer struct {
	ID         int64    `gorm:"primaryKey;column:id" json:"id"`
	Correct    bool     `gorm:"not null;column:correct" json:"correct"`
	Text       string   `gorm:"type:varchar(255);column:text" json:"text"`
	QuestionID int64    `gorm:"column:question_id" json:"-"`
	Question   Question `gorm:"foreignKey:QuestionID" json:"-"`
}

// Image represents the images table in the database

type Image struct {
	ID         int64    `gorm:"primaryKey;column:id" json:"id"`
	Data       []byte   `gorm:"type:blob;column:data" json:"-"`
	Name       string   `gorm:"type:varchar(255);column:name" json:"-"`
	Row        int      `gorm:"not null;column:row" json:"-"`
	QuestionID int64    `gorm:"column:question_id" json:"-"`
	Question   Question `gorm:"foreignKey:QuestionID" json:"-"`
}
