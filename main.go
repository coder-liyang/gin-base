package main

import (
	"fmt"
	"liyangweb.com/z-library/pkg/logging"
	"liyangweb.com/z-library/pkg/setting"
	"liyangweb.com/z-library/routes"
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