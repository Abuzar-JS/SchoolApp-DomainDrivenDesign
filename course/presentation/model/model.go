package models

type CreateCourseRequest struct {
	Title     string `validate:"required,min=1,max=200" json:"title"`
	StudentID int    `validate:"required,max=200,min=1" json:"student_id"`
	SchoolID  int    `validate:"required,max=200,min=1" json:"school_id"`
}

type UpdateCourseRequest struct {
	Title     *string `json:"title"`
	StudentID *int    `json:"student_id"`
	SchoolID  int     `json:"school_id"`
	CourseID  int     `json:"course_id"`
}

type CourseResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	StudentID int    `json:"student_id"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
