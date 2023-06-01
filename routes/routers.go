package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/controllers"
	"github.com/yacw-team/yacw/controllers/Chat"
	"github.com/yacw-team/yacw/controllers/Game"
	"github.com/yacw-team/yacw/controllers/Personality"
	"github.com/yacw-team/yacw/controllers/Prompt"
	"net/http"
	"os"
	"path"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//添加中间件
	apiPath := os.Getenv("API_PATH")
	fmt.Println(apiPath)
	r.Use(ApiPrefixMiddleware(apiPath))
	r.GET(apiPath+"/v1/chat/prompts", Prompt.GetPromptShop)
	r.POST(apiPath+"/v1/chat/myprompts", AuthMiddleware(), Prompt.GetMyPrompt)
	r.POST(apiPath+"/v1/chat/prompts", AuthMiddleware(), Prompt.CreatePrompt)
	r.POST(apiPath+"/v1/chat/deleteprompts", AuthMiddleware(), Prompt.DeletePrompt)
	r.POST(apiPath+"/v1/chat/apiKey", AuthMiddleware(), controllers.VerifyApiKey)
	r.GET(apiPath+"/v1/chat/personality", Personality.GetPersonalityShop)
	r.POST(apiPath+"/v1/chat/mypersonality", AuthMiddleware(), Personality.GetMyPersonality)

	r.POST(apiPath+"/v1/chat/getmessage", Chat.GetChatMessage)
	r.POST(apiPath+"/v1/chat/getchat", Chat.GetChatId)

	r.POST(apiPath+"/v1/translate/translate", AuthMiddleware(), controllers.Translate)
	r.POST(apiPath+"/v1/chat/chat", AuthMiddleware(), Chat.SendMessage)
	r.POST(apiPath+"/v1/chat/new", AuthMiddleware(), Chat.NewChat)
	r.POST(apiPath+"/v1/chat/deletechat", AuthMiddleware(), Chat.DeleteChat)

	r.GET(apiPath+"/v1/game/story", Game.GetGameBackground)
	r.POST(apiPath+"/v1/game/new", AuthMiddleware(), Game.ChooseGameBackground)
	r.POST(apiPath+"/v1/game/chat", AuthMiddleware(), Game.SendGameMessage)
	return r
}

// ApiPrefixMiddleware 中间件添加前缀
func ApiPrefixMiddleware(apiPath string) gin.HandlerFunc {
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
