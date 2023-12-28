package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stock-controller/app/config"
	"github.com/stock-controller/app/router"
)

func main() {
	app := gin.Default()


	app.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*") // Especifica los orígenes permitidos
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		ctx.Next()
	})

	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}

	db := config.GetDB()
	defer db.Close()

	router.GetRouter(app, db)

	//test := repository.CompaniesRepository{Db: db}
	//test.GetCompanies()

	app.Run("localhost:8080")
}
