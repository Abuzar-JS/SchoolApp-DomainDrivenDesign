package school

import (
	"context"

	"github.com/Abuzar-JS/Go-StudentApp/school/domain"
	"github.com/Abuzar-JS/Go-StudentApp/school/domain/school"
	"github.com/Abuzar-JS/Go-StudentApp/school/infrastructure/postgres"
	"github.com/Abuzar-JS/Go-StudentApp/school/presentation/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type client struct {
	schoolRepo school.Repository
}

type Client interface {
	GetBySchoolID(ctx context.Context, locationID int) (domain.School, error)
}

func (c *client) GetBySchoolID(ctx context.Context, schoolID int) (domain.School, error) {
	return c.schoolRepo.GetBySchoolID(schoolID)
}

func InitiateAndRegister(router *gin.Engine, database *gorm.DB, validate *validator.Validate) Client {
	// register tariff routes
	http.RegisterRoutes(router, database, validate)

	// return client
	return &client{
		schoolRepo: postgres.NewSchoolPostgres(database),
	}
}
