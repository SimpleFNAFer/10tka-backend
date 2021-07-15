package models

type Post struct {
	ID             int +
	Role           string    `gorm:"column:role" json:"role"`
	Email          string    `gorm:"column:email" json:"email"`
	Phone          string    `gorm:"column:phone" json:"phone"`
	Name           string    `gorm:"column:name" json:"name"`
	Surname        string    `gorm:"column:surname" json:"surname"`
	Patronymic     string    `gorm:"column:patronymic" json:"patronymic"`
	PasswordHash   []byte    `gorm:"column:password_hash" json:"-"`
	StudentGroup   string    `gorm:"column:student_group" json:"student_group"`
	Disciplines    []string  `gorm:"-" json:"prof_disciplines"`
	DisciplinesOrm string    `gorm:"column:prof_disciplines" json:"prof_disciplines_str"`
	About          string    `gorm:"column:about" json:"about"`
}
