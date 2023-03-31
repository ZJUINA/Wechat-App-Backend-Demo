package v0

import (
	"WangYiYunDemo/music/mApi"
	"github.com/gin-gonic/gin"
)

func RegisterMusic(BasicRouter *gin.Engine) {
	BasicRouter.GET("/Hello", mApi.Hello)
	Music := BasicRouter.Group("/music")
	Music.POST("/ListSongs", mApi.ListSongs)
	Music.POST("/ListComments", mApi.ListComments)
}
