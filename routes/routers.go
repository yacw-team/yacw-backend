package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", controllers.Hello)
	r.POST("/v1/chat/getchat", controllers.GetChatId)
	return r
}
