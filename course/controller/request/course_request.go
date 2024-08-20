package request

type CreateCourseRequest struct {
	Title     string `validate:"required,min=1,max=200" json:"title"`
	StudentID int    `validate:"required,max=200,min=1" json:"student_id"`
}

type UpdateCourseRequest struct {
	Title     *string `json:"title"`
	StudentID *int    `json:"student_id"`
}
