package v1

import (
	"meetup/utils/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, "up")

}
