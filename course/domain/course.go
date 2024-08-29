package domain

type Course struct {
	ID        int
	Title     string
	StudentID int `gorm:"foreignKey:StudentID"`
}
