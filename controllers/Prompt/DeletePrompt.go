package Prompt

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

// DeletePrompt 删除用户创建的prompt
func DeletePrompt(c *gin.Context) {
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
	promptId, ok := reqBody["promptsId"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}

	slice := []string{apiKey, promptId}
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

	uid, err := utils.Encrypt(apiKey) //用户id
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3006"})
		return
	}
	id := promptId

	result := utils.DB.Table("prompt").Where("id = ? AND uid = ?", id, uid).Delete(models.Prompt{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1008"})
		return
	}
	c.JSON(http.StatusOK, models.ErrCode{ErrCode: "0000"})
}
