package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/repository"
	"github.com/stock-controller/app/useCase"
)

func GetRouter(app *gin.Engine, db *sql.DB) {
	generalRepository := repository.Repository{Db: db}

	bulkCreate := useCase.BulkCreate{BulkCreateRepository: generalRepository}

	getCompanies := useCase.GetCompanies{CompanyRepository: generalRepository}
	createCompany := useCase.CreateCompany{CompanyRepository: generalRepository}
	getMovementsByCompany := useCase.GetMovementsByCompany{MovementRepository: generalRepository}

	//getProducts := useCase.GetProducts{ProductRepository: generalRepository}
	createProduct := useCase.CreateProduct{ProductRepository: generalRepository}

	app.GET("/ping", useCase.PingController)

	app.POST("/bulk-create", bulkCreate.Handle)

	companyRoutes := app.Group("/company")
	companyRoutes.GET("/", getCompanies.Handle)
	companyRoutes.GET("/:id/movements", getMovementsByCompany.Handle)
	companyRoutes.POST("/", createCompany.Handle)

	productRoutes := app.Group("/product")
	productRoutes.POST("/", createProduct.Handle)

}
