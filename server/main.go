package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/router"
)

func main() {
	app := gin.Default()

	router.GetRouter(app)

	app.Run("localhost:8080")
}
