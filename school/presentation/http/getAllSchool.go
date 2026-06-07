package http

import (
	"net/http"

	"github.com/Abuzar-JS/Go-StudentApp/school/application"
	models "github.com/Abuzar-JS/Go-StudentApp/school/presentation/model"

	"github.com/gin-gonic/gin"
)

// GetAllSchools godoc
// @Summary Get all schools
// @Tags schools
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /schools [get]
func NewGetAllSchool(
	getAll application.GetAllSchool,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		schoolResponse := getAll()
		webResponse := models.Response{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   schoolResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}
