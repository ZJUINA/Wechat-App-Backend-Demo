package main

import (
	"WangYiYunDemo/music/mDAO"
	v0 "WangYiYunDemo/music/v0"
	"github.com/gin-gonic/gin"
)

var BasicRouter *gin.Engine

func init() {
	mDAO.DBInit()
	BasicRouter = gin.Default()
	v0.RegisterMusic(BasicRouter)
}

func main() {
	//conf.Init() 暂时没用到conf， 先放在这

	BasicRouter.Run(":8080")
}
