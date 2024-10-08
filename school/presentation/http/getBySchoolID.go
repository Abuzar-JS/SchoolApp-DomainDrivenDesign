package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/Go-StudentApp/school/application"

	"github.com/gin-gonic/gin"
)

func NewGetBySchoolID(
	getById application.GetBySchoolId,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		schoolId := ctx.Param("school_id")
		ID, err := strconv.Atoi(schoolId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		school, err := getById(ctx.Request.Context(), ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "school found",
			"data":    school,
		})
	}
}
