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
	updateCompanyById := useCase.UpdateCompanyById{CompanyRepository: generalRepository}
	deleteCompanyById := useCase.DeleteCompanyById{CompanyRepository: generalRepository}

	getMovementsByCompany := useCase.GetMovementsByCompany{MovementRepository: generalRepository}
	createMovement := useCase.CreateMovement{MovementRepository: generalRepository}
	updateMovementById := useCase.UpdateMovementById{MovementRepository: generalRepository}
	deleteMovementById := useCase.DeleteMovementById{MovementRepository: generalRepository}

	getProducts := useCase.GetProducts{ProductRepository: generalRepository}
	createProduct := useCase.CreateProduct{ProductRepository: generalRepository}
	updateProductById := useCase.UpdateProductById{ProductRepository: generalRepository}
	deleteProductById := useCase.DeleteProductById{ProductRepository: generalRepository}

	app.GET("/ping", useCase.PingController)

	app.POST("/bulk-create", bulkCreate.Handle)

	companyRoutes := app.Group("/company")
	companyRoutes.GET("/", getCompanies.Handle)
	companyRoutes.GET("/:companyId/products", getProducts.Handle)
	companyRoutes.POST("/", createCompany.Handle)
	companyRoutes.PUT("/:companyId", updateCompanyById.Handle)
	companyRoutes.DELETE("/:companyId", deleteCompanyById.Handle)

	companyRoutes.GET("/:companyId/movements", getMovementsByCompany.Handle)
	companyRoutes.POST("/:companyId/movements", createMovement.Handle)
	companyRoutes.PUT("/:companyId/movements/:movementId", updateMovementById.Handle)
	companyRoutes.DELETE("/movement/:movementId", deleteMovementById.Handle)

	productRoutes := app.Group("/product")
	productRoutes.POST("/", createProduct.Handle)
	productRoutes.PUT("/:productId", updateProductById.Handle)
	productRoutes.DELETE("/:productId", deleteProductById.Handle)

}
