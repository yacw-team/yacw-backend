package Game

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strconv"
	"strings"
)

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

	uid, err := utils.Encrypt(apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3006"})
		return
	}
	//检查一下之前是否有这个用户的历史，有则删除
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

	result, err = CheckAndReSend(message, modelId, apiKey)
	if err != nil {
		errCode := utils.GPTRequestErrorCode(err)
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: errCode})
		return
	}

	var gamemessage models.GameMessage
	gamemessage.Uid = uid
	gamemessage.Story = result["story"].(string)
	jsonData, err := json.Marshal(result["choice"].([]interface{}))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2005"})
		return
	}
	gamemessage.Chocie = string(jsonData)
	gamemessage.Round = int(result["round"].(float64))
	gamemessage.GameId = gameId

	err = utils.DB.Table("gamemessage").Create(&gamemessage).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	c.JSON(http.StatusOK, result)

}
