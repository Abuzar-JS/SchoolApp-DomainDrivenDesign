package controller

import (
	"data/course/controller/request"
	"data/course/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CourseController struct
type CourseController struct {
	CourseService service.CourseService
}

func NewCourseController(service service.CourseService) *CourseController {
	return &CourseController{
		CourseService: service,
	}

}

// Create Controller
func (controller *CourseController) Create(ctx *gin.Context) {
	createCourseRequest := request.CreateCourseRequest{}
	if err := ctx.ShouldBindJSON(&createCourseRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	course, err := controller.CourseService.Create(createCourseRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Course created successfully",
		"course":  course,
	})

}

// Update Controller
func (controller *CourseController) Update(ctx *gin.Context) {
	updateCourseRequest := request.UpdateCourseRequest{}
	err := ctx.ShouldBindJSON(&updateCourseRequest)

	courseId := ctx.Param("course_id")
	id, err := strconv.Atoi(courseId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if updateCourseRequest.Title == nil && updateCourseRequest.StudentID == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "atleast one field is required to update course",
		})
		return
	}

	err = controller.CourseService.Update(id, updateCourseRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Course updated successfully",
	})
}

func (controller *CourseController) Delete(ctx *gin.Context) {
	courseId := ctx.Param("course_id")
	id, err := strconv.Atoi(courseId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = controller.CourseService.Delete(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "course deleted successfully",
	})
}

// FindById Controller
func (controller *CourseController) FindById(ctx *gin.Context) {
	schoolIDParam := ctx.Param("school_id")
	schoolID, err := strconv.Atoi(schoolIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	studentIDParam := ctx.Param("student_id")

	studentID, err := strconv.Atoi(studentIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	courseIdParam := ctx.Param("course_id")

	courseID, err := strconv.Atoi(courseIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	request := service.GetCourseRequest{
		SchoolID:  schoolID,
		StudentID: studentID,
		CourseID:  courseID,
	}

	course, err := controller.CourseService.FindById(request)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "course found",
		"data":    course,
	})
}

// FindByAll Controller
func (controller *CourseController) FindByStudentID(ctx *gin.Context) {
	schoolIDParam := ctx.Param("school_id")
	studentIDParam := ctx.Param("student_id")
	schoolID, err := strconv.Atoi(schoolIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	studentID, err := strconv.Atoi(studentIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	request := service.GetCourseRequest{
		SchoolID:  schoolID,
		StudentID: studentID,
	}

	course, err := controller.CourseService.FindByStudentID(request)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "course found",
		"data":    course,
	})
}
