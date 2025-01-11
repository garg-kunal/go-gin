package model

type UserIdentification struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Nbf   float64  `json:"nbf"`
}

type User struct {
	Id        int     `gorm:"primaryKey" json:"id"`
	Email     string  `json:"email" gorm:"unique;not null" binding:"required"`
	Password  string  `json:"password"`
	// UserNotes []Notes `gorm:"foreignKey:NotesId" json:"user_notes"`
}

func (User) TableName() string {
	return "user"
}
