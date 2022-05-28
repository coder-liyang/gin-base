package routes

import (
	"github.com/gin-gonic/gin"
	"liyangweb.com/gin-base/middleware/jwt"
	"liyangweb.com/gin-base/pkg/e"
	"liyangweb.com/gin-base/pkg/setting"
	"liyangweb.com/gin-base/routes/api"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  e.GetMsg(e.SUCCESS),
			"data": nil,
		})
	})
	r.GET("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取列表
		//apiv1.GET("/contents", v1.GetContents)
	}
	return r
}
