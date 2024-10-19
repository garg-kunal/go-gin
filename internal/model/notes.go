package model

type Notes struct {
	Id      int 	`gorm:"primaryKey" json:"id"`
	Title   string	`json:"title"`
	Status  bool 	`json:"status"`
}

func (Notes) TableName() string {
	return "notes"
}