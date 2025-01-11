package model

type Notes struct {
	Id     int    `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
	UserId int    `gorm:"foreignKey:UserId" json:"user_id"`
}

func (Notes) TableName() string {
	return "notes"
}
