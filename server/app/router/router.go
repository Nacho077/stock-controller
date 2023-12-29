package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/use_case"
)

func GetRouter(app *gin.Engine, db *sql.DB) {
	generalRepository := repository.Repository{Db: db}

	NewCompanyController := use_case.GetCompanies{CompanyRepository: generalRepository}
	NewMovementsByCompanyController := use_case.GetMovementsByCompany{MovementRepository: generalRepository}

	app.GET("/ping", use_case.PingController)

	companyRoute := app.Group("/company")
	companyRoute.GET("/", NewCompanyController.Handle)
	companyRoute.GET("/:id/movements", NewMovementsByCompanyController.Handle)

}
