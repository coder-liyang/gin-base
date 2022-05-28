package main

import (
	"fmt"
	"liyangweb.com/gin-base/pkg/logging"
	"liyangweb.com/gin-base/pkg/setting"
	"liyangweb.com/gin-base/routes"
	"net/http"
)

func main() {
	setting.Setup()
	//models.Setup()
	logging.Setup()

	router := routes.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}