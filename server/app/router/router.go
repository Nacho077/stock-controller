package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/controller"
	"github.com/stock-controller/app/repository"
)

func GetRouter(app *gin.Engine, db *sql.DB) {
	generalRepository := repository.Repository{Db: db}

	NewCompany := controller.GetCompanies{CompanyRepository: generalRepository}
	NewMovementsByCompany := controller.GetMovementsByCompany{MovementRepository: generalRepository}
	NewBulkCreate := controller.BulkCreate{BulkCreateRepository: generalRepository}

	app.GET("/ping", controller.PingController)

	app.POST("/bulk-create", NewBulkCreate.Handle)

	companyRoutes := app.Group("/company")
	companyRoutes.GET("/", NewCompany.Handle)
	companyRoutes.GET("/:id/movements", NewMovementsByCompany.GetAll)
	companyRoutes.POST("/", NewCompany.Handle)

	//productRoutes := app.Group("/product")
	//productRoutes.POST("/", NewCompany.Handle)

}
