package Game

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strconv"
)

// 发送给API所需的结构
type SendStructure struct {
	Story  string              `json:"story"`
	Choice []map[string]string `json:"choice"`
	Round  int                 `json:"round"`
}

// 存在数据库里的结构
type StoreStructure struct {
	Story  string `json:"story"`
	Choice string `json:"choice"`
	Round  int    `json:"round"`
}

func SendGameMessage(c *gin.Context) {
	var reqBody map[string]interface{}
	reqTemp, ok := c.Get("reqBody")
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}
	reqBody = reqTemp.(map[string]interface{})

	apiKey, ok := reqBody["apiKey"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}
	choiceId, ok := reqBody["choiceID"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}
	modelStr, ok := reqBody["modelId"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}

	slice := []string{apiKey, choiceId, modelStr}
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

	uid, err := utils.Encrypt(apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3006"})
		return
	}

	modelId, err := strconv.Atoi(modelStr)
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2005"})
		return
	}

	if modelId < 0 || modelId > 6 {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1005"})
		return
	}

	var gamemessage models.GameMessage
	err = utils.DB.Table("gamemessage").Where("uid = ?", uid).Find(&gamemessage).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	//获取背景字段
	var systemPrompt string
	err = utils.DB.Table("game").Select("systemprompt").Where("gameId = ?", gamemessage.GameId).Find(&systemPrompt).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	jsonData, err := json.MarshalIndent(gamemessage, "", "  ")
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2005"})
		return
	}

	//获取历史信息
	history := string(jsonData)

	history, err = transForm(history)
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2005"})
		return
	}

	//构造请求体
	message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: history,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: choiceId,
		},
	}

	var result map[string]interface{}
	result, err = CheckAndReSend(message, modelId, apiKey)
	if err != nil {
		errCode := utils.GPTRequestErrorCode(err)
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: errCode})
		return
	}

	story, err := summarize(gamemessage.Story, result["story"].(string), modelId, apiKey)
	if err != nil {
		errCode := utils.GPTRequestErrorCode(err)
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: errCode})
		return
	}

	//删除历史记录
	err = utils.DB.Exec("delete from gamemessage where uid = ?", uid).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	gamemessage.Story = story
	jsonData, err = json.Marshal(result["choice"].([]interface{}))
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2005"})
		return
	}
	gamemessage.Chocie = string(jsonData)
	gamemessage.Round = int(result["round"].(float64))

	err = utils.DB.Table("gamemessage").Create(&gamemessage).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// 将存储的数据转为正确的json字符串
func transForm(str string) (string, error) {

	var secondData StoreStructure
	err := json.Unmarshal([]byte(str), &secondData)
	if err != nil {
		return "", err
	}

	var firstChoices []map[string]string
	err = json.Unmarshal([]byte(secondData.Choice), &firstChoices)
	if err != nil {
		return "", err
	}

	firstData := SendStructure{
		Story:  secondData.Story,
		Choice: firstChoices,
		Round:  secondData.Round,
	}

	sendJSON, err := json.Marshal(firstData)
	if err != nil {
		return "", err
	}

	return string(sendJSON), nil
}

// 概要历史信息
func summarize(str1 string, str2 string, modelId int, apiKey string) (string, error) {
	prompt := "我将给你两段json，忽略json中的choice字段，帮我融合并且概括一下它的”story“字段的情节，在附带上story字段的重要信息的基础上，尽量的缩短长度（直接输出概括的结果不要附带其他的内容）。"
	//构造请求体
	message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: prompt,
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: str1,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: str2,
		},
	}
	data, err := SendToOpenAI(message, modelId, apiKey)
	if err != nil {
		return "", err
	}

	return data, nil
}
