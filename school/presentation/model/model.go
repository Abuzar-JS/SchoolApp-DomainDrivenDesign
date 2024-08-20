package models

type CreateSchoolRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type UpdateSchoolRequest struct {
	ID   int    `validate:"required" json:"id"`
	Name string `validate:"required,max = 200, min =1" json:"name"`
}

type SchoolResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
