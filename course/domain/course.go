package domain

type Courses []Course

// func (cs Courses) FilterByX() Courses {
// 	return Courses
// }

type Course struct {
	ID        int
	Title     string
	StudentID int
}
