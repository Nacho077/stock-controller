package use_case

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
