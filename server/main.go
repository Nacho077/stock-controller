package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stock-controller/app/config"
	"github.com/stock-controller/app/router"
)

func main() {
	app := gin.Default()

	router.GetRouter(app)

	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}

	db := config.GetDB()
	defer db.Close()

	app.Run("localhost:8080")
}
