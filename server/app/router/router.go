package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/useCase"
)

func GetRouter(app *gin.Engine, db *sql.DB) {
	generalRepository := repository.Repository{Db: db}

	NewCompany := useCase.GetCompanies{CompanyRepository: generalRepository}
	postCompany := useCase.CreateCompany{CompanyRepository: generalRepository}
	NewMovementsByCompany := useCase.GetMovementsByCompany{MovementRepository: generalRepository}
	NewBulkCreate := useCase.BulkCreate{BulkCreateRepository: generalRepository}

	app.GET("/ping", useCase.PingController)

	app.POST("/bulk-create", NewBulkCreate.Handle)

	companyRoutes := app.Group("/company")
	companyRoutes.GET("/", NewCompany.Handle)
	companyRoutes.GET("/:id/movements", NewMovementsByCompany.Handle)
	companyRoutes.POST("/", postCompany.Handle)

	//productRoutes := app.Group("/product")
	//productRoutes.POST("/", NewCompany.Handle)

}
