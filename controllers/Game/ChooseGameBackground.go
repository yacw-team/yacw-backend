package Game

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/yacw-team/yacw/controllers"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type gameMessage struct {
	Uid    string `gorm:"uid"`
	Story  string
	Chocie string
	Round  int
}

func ChooseGameBackground(c *gin.Context) {
	var reqBody map[string]interface{}
	reqTemp, ok := c.Get("reqBody")
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}
	reqBody = reqTemp.(map[string]interface{})

	//获取数据
	apiKey := reqBody["apiKey"].(string)
	gameId := reqBody["gameId"].(string)
	modelStr := reqBody["modelId"].(string)

	slice := []string{apiKey, gameId, modelStr}
	if !utils.Utf8Check(slice) {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1011"})
		return
	}

	apiKeyCheck := utils.IsValidApiKey(apiKey)
	if !apiKeyCheck {
		var errCode models.ErrCode
		errCode.ErrCode = "3004"
		c.JSON(http.StatusBadRequest, errCode)
		return
	}

	//检查一下之前是否有这个用户的历史，有则删除
	uid, err := utils.Encrypt(apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3006"})
		return
	}

	err = utils.DB.Exec("delete from gamemessage where uid = ?", uid).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	modelId, err := strconv.Atoi(modelStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2005"})
		return
	}

	//获取prompt字段
	systemPrompt := ""

	err = utils.DB.Table("game").Select("systemprompt").Where("gameId = ?", gameId).Find(&systemPrompt).Error
	if err != nil || strings.EqualFold(systemPrompt, "") {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	//构造请求体
	message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
	}

	var result map[string]interface{}

	for {
		data, err := sendToOpenAI(message, c, systemPrompt, modelId, apiKey)
		if err != nil {
			errCode := utils.GPTRequestErrorCode(err)
			c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: errCode})
			return
		}

		// 将字符数据解析为map[string]interface{}类型
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}

		// 检查解析后的JSON数据是否符合预期格式
		if isValidResult(result) {
			// 返回JSON响应
			break
		}
		time.Sleep(3 * time.Second)
	}

	var gamemessage gameMessage
	gamemessage.Uid = uid
	gamemessage.Story = result["story"].(string)
	jsonData, _ := json.Marshal(result["choice"].([]interface{}))
	gamemessage.Chocie = string(jsonData)
	gamemessage.Round = int(result["round"].(float64))

	err = utils.DB.Table("gamemessage").Create(&gamemessage).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	c.JSON(http.StatusOK, result)

}

func sendToOpenAI(message []openai.ChatCompletionMessage, c *gin.Context, systemPrompt string, modelId int, apiKey string) (string, error) {
	// 创建 OpenAI 客户端
	client := openai.NewClient(apiKey)
	ctx := context.Background()

	//构造请求体
	req := openai.ChatCompletionRequest{
		Model: controllers.Model[modelId],
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
		},
	}

	//获取回复
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	//获取回复
	assistant := resp.Choices[0].Message.Content
	return assistant, nil
}

// 检查生成的格式是否符合预期
func isValidResult(result map[string]interface{}) bool {
	// 检查字段是否存在且不为空
	if result["story"] == nil || result["choice"] == nil || result["round"] == nil {
		return false
	}

	// 进一步检查字段值的类型和合法性
	if _, ok := result["story"].(string); !ok {
		return false
	}

	choices, ok := result["choice"].([]interface{})
	if !ok || len(choices) == 0 {
		return false
	}

	validKeys := map[string]bool{
		"A": true,
		"B": true,
		"C": true,
		"D": true,
	}

	for _, choice := range choices {
		choiceMap, ok := choice.(map[string]interface{})
		if !ok || len(choiceMap) != 1 {
			return false
		}

		for key := range choiceMap {
			if !validKeys[key] {
				return false
			}
		}
	}

	if round, ok := result["round"].(float64); !ok || round <= 0 {
		return false
	}

	// 根据需要添加其他验证规则

	return true
}
