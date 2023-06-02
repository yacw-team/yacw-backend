package Game

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strconv"
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
	apiKey, ok := reqBody["apiKey"].(string)

	gameId, ok := reqBody["gameId"].(string)

	modelStr, ok := reqBody["modelId"].(string)

	apiKeyCheck := utils.IsValidApiKey(apiKey)
	if !apiKeyCheck {
		var errCode models.ErrCode
		errCode.ErrCode = "3004"
		c.JSON(http.StatusBadRequest, errCode)
		return
	}

	uid, _ := utils.Encrypt(apiKey)

	//检查一下之前是否有这个用户的历史，有则删除
	err := utils.DB.Exec("delete from gamemessage where uid = ?", uid).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
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

	//获取prompt字段
	systemPrompt := ""

	db_result := utils.DB.Table("game").Select("systemprompt").Where("gameId = ?", gameId).Find(&systemPrompt)

	//数据库没有查到数据
	if db_result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1008"})
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

	var gamemessage models.GameMessage
	gamemessage.Uid = uid
	gamemessage.Story, ok = result["story"].(string)

	jsonData, err := json.Marshal(result["choice"].([]interface{}))

	gamemessage.Chocie = string(jsonData)
	gamemessage.Round = int(result["round"].(float64))
	gamemessage.GameId = gameId

	err = utils.DB.Table("gamemessage").Create(&gamemessage).Error

	c.JSON(http.StatusOK, result)

}
