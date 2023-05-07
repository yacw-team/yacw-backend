package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/controllers"
	"os"
	"path"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//添加中间件
	r.Use(ApiPrefixMiddleware())
	r.GET("/v1/chat/prompts", controllers.GetPromptShop)
	r.GET("/v1/chat/myprompts", controllers.GetMyPrompt)
	r.POST("/v1/chat/prompts", controllers.CreatePrompt)
	r.DELETE("/v1/chat/prompts", controllers.DeletePrompt)
	r.POST("/v1/chat/apiKey", controllers.VerifyApiKey)

	r.POST("/v1/translate/translate", controllers.Translate)
	r.POST("/v1/chat/chat", controllers.SendMessage)
	return r
}

// ApiPrefixMiddleware 中间件添加前缀
func ApiPrefixMiddleware() gin.HandlerFunc {
	apiPath := os.Getenv("API_PATH")
	return func(c *gin.Context) {
		c.Request.URL.Path = path.Join(apiPath, c.Request.URL.Path)
		c.Next()
	}
}

//// AuthMiddleware 中间件验证用户输入了API
//func AuthMiddleware() gin.HandlerFunc {
//
//}
