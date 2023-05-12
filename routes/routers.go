package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/controllers"
	"net/http"
	"os"
	"path"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//添加中间件
	r.Use(ApiPrefixMiddleware())
	r.GET("/api/v1/chat/prompts", controllers.GetPromptShop)
	r.POST("/api/v1/chat/myprompts", AuthMiddleware(), controllers.GetMyPrompt)
	r.POST("/api/v1/chat/prompts", AuthMiddleware(), controllers.CreatePrompt)
	r.DELETE("/api/v1/chat/prompts", controllers.DeletePrompt)
	r.POST("/api/v1/chat/apiKey", controllers.VerifyApiKey)
	r.GET("/api/v1/chat/personality", controllers.GetPersonalityShop)
	r.POST("/api/v1/chat/mypersonality", AuthMiddleware(), controllers.GetMyPersonality)
	r.POST("/api/v1/chat/personality", AuthMiddleware(), controllers.CreatePersonality)
	r.DELETE("/api/v1/chat/personality", controllers.DeletePersonality)

	r.POST("/api/v1/chat/getmessage", controllers.GetChatMessage)
	r.POST("/api/v1/chat/getchat", controllers.GetChatId)

	r.POST("/api/v1/translate/translate", AuthMiddleware(), controllers.Translate)
	r.POST("/api/v1/chat/chat", AuthMiddleware(), controllers.SendMessage)
	r.POST("/api/v1/chat/new", AuthMiddleware(), controllers.NewChat)
	r.DELETE("/api/v1/chat/chat", AuthMiddleware(), controllers.DeleteChat)
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

// AuthMiddleware 中间件验证用户输入了apiKey
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody map[string]interface{}

		// 使用 Gin 的 ShouldBindJSON 方法将请求参数绑定到 map 中
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//访问 map 中的字段来获取 POST 请求的每个字段
		apiKey := reqBody["apiKey"].(string)

		if apiKey == "" {
			c.Redirect(http.StatusMovedPermanently, "/index?error=nokey")
			return
		}
		c.Set("reqBody", reqBody)
		c.Next()
	}
}
