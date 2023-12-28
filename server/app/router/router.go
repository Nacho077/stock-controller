package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/controller"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/service"
)

func GetRouter(app *gin.Engine, db *sql.DB) {
	NewCompanyRepository := repository.CompaniesRepository{db}
	NewCompanyService := service.CompanyService{NewCompanyRepository}
	NewCompanyController := controller.CompanyController{NewCompanyService}

	app.GET("/ping", controller.PingController)

	companyRoute := app.Group("/company")

	companyRoute.GET("/", NewCompanyController.GetAllCompanies)
}
