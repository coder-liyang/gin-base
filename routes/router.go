package routes

import (
	"github.com/gin-gonic/gin"
	"liyangweb.com/z-library/middleware/jwt"
	"liyangweb.com/z-library/pkg/setting"
	"liyangweb.com/z-library/routes/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)
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
