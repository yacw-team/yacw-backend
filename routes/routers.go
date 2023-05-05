package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/controllers"
	"os"
	"path"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(ApiPrefixMiddleware())
	r.GET("/v1/chat/prompts", controllers.GetPromptShop)
	r.GET("/v1/chat/myprompts", controllers.GetMyPrompt)
	r.POST("/v1/chat/prompts", controllers.CreatePrompt)
	r.DELETE("/v1/chat/prompts", controllers.DeletePrompt)
	return r
}

func ApiPrefixMiddleware() gin.HandlerFunc {
	apiPath := os.Getenv("API_PATH")
	if apiPath == "" {
		apiPath = "/ "
	}
	return func(c *gin.Context) {
		c.Request.URL.Path = path.Join(apiPath, c.Request.URL.Path)
		c.Next()
	}
}
