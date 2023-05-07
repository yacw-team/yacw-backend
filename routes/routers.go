package routes

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(ApiPrefixMiddleware())
	r.GET("/v1/chat/prompts", controllers.GetPrompts)
	r.POST("/v1/chat/prompts", controllers.CreatePrompt)
	r.DELETE("/v1/chat/prompts", controllers.DeletePrompt)
	r.POST("/v1/chat/getChat", controllers.GetChatId)
	r.POST("/v1/chat/getMessage", controllers.GetChatMessage)
	return r
}

func ApiPrefixMiddleware() gin.HandlerFunc {
	apiPath := os.Getenv("API_PATH")
	if apiPath == "" {
		apiPath = "/api"
	}
	return func(c *gin.Context) {
		c.Request.URL.Path = path.Join(apiPath, c.Request.URL.Path)
		c.Next()
	}
}
