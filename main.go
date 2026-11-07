package main

import (
	"aio/config"
	"aio/internal/apps/timeline"
	"aio/internal/common"
	"aio/internal/resources"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode) //关闭Debug

	config.Load("config.json")
	common.InitDB()

	r := gin.New()
	r.HandleMethodNotAllowed = false
	r.MaxMultipartMemory = config.Config.Server.Max.Upload << 20 //最大上传文件100M
	r.Use(common.GlobalRecovery())

	//初始化路由
	resources.Init(r)
	timeline.Init(r)

	//	不用 r.Run(":8080")
	server := &http.Server{
		Handler:        r,
		Addr:           ":" + config.Config.Server.Port,
		MaxHeaderBytes: config.Config.Server.Max.Header << 20, //Header最大10M
	}
	server.ListenAndServe()
}
