package domain

type Student struct {
	ID       int
	Name     string
	Class    string
	SchoolID int `gorm:"foreignKey:SchoolID"`
}
