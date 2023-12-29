package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/useCase"
)

func GetRouter(app *gin.Engine, db *sql.DB) {
	generalRepository := repository.Repository{Db: db}

	NewCompanyController := useCase.GetCompanies{CompanyRepository: generalRepository}
	NewMovementsByCompanyController := useCase.GetMovementsByCompany{MovementRepository: generalRepository}

	app.GET("/ping", useCase.PingController)

	companyRoute := app.Group("/company")
	companyRoute.GET("/", NewCompanyController.Handle)
	companyRoute.GET("/:id/movements", NewMovementsByCompanyController.Handle)

}
