package v1

import (
	"github.com/gin-gonic/gin"
	"liyangweb.com/gin-base/models"
	"liyangweb.com/gin-base/pkg/app"
	"liyangweb.com/gin-base/pkg/e"
	"net/http"
	"strconv"
)

func GetLetters(c *gin.Context) {
	appG := app.Gin{C: c}

	userid, _ := strconv.Atoi(c.Param("userid"))
	var page int64 = 1

	appG.Response(http.StatusOK, e.SUCCESS, models.Letter{}.GetAll(page, userid))
}