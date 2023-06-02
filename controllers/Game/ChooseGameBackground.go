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

	apiKey, ok := reqBody["apiKey"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}
	gameId, ok := reqBody["gameId"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}
	modelStr, ok := reqBody["modelId"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}

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

	if modelId < 0 || modelId > 6 {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1005"})
		return
	}

	systemPrompt := ""

	db_result := utils.DB.Table("game").Select("systemprompt").Where("gameId = ?", gameId).Find(&systemPrompt)
	if db_result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	if db_result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1008"})
		return
	}
	
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
	gamemessage.Story, ok = result["story"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}
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
