package response

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
