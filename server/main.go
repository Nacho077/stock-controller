package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stock-controller/app/config"
	"github.com/stock-controller/app/router"
)

func main() {
	app := gin.Default()
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
