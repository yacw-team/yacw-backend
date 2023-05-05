package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/controllers"
	"github.com/yacw-team/yacw/controllers/v1/chat"
	"os"
	"path"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(ApiPrefixMiddleware())
	r.GET("/v1/chat/prompts", controllers.GetPrompts)
	r.POST("/v1/chat/prompts", controllers.CreatePrompt)
	r.DELETE("/v1/chat/prompts", controllers.DeletePrompt)
	r.POST("/v1/chat/new", chat.NewChat)
	r.DELETE("/v1/chat/chat", chat.DeleteChat)
	r.POST("/v1/chat/chat", chat.SendChat)
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
