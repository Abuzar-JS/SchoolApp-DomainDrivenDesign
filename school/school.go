package school

import (
	"context"
	"data/school/domain"
	"data/school/domain/school"
	"data/school/infrastructure/postgres"
	"data/school/presentation/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type client struct {
	schoolRepo school.SchoolRepository
}

type Client interface {
	// GetStudentByID returns a list of utilities added by a user for a given location
	GetStudentByIDClient(ctx context.Context, locationID int) (domain.School, error)
}

func (c *client) GetStudentByIDClient(ctx context.Context, schoolID int) (domain.School, error) {
	return c.schoolRepo.GetBySchoolID(schoolID)
}

func InitiateAndRegister(router *gin.Engine, database *gorm.DB, validate *validator.Validate) Client {
	// register tariff routes
	http.RegisterRoutes(router, database, validate)

	// return clien
	return &client{
		schoolRepo: postgres.NewSchoolPostgres(database),
	}
}
