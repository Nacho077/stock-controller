package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stock-controller/app/controller"
)

func GetRouter(app *gin.Engine) {
	app.GET("/ping", controller.PingController)
}
