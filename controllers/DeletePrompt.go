package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

type deletePromptReqBody struct {
	ApiKey    string `json:"apiKey"`
	PromptsId string `json:"promptsId"`
}

// DeletePrompt 删除用户创建的prompt
func DeletePrompt(c *gin.Context) {
	var reqBody deletePromptReqBody
	err := c.Bind(&reqBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1010"})
		return
	}

	apiKeyCheck := utils.IsValidApiKey(reqBody.ApiKey)
	if apiKeyCheck == false {
		var errCode models.ErrCode
		errCode.ErrCode = "3004"
		c.JSON(http.StatusBadRequest, errCode)
		return
	}

	uid := utils.Encrypt(reqBody.ApiKey) //用户id
	id := reqBody.PromptsId

	err = utils.DB.Table("prompt").Where("id = ? AND uid = ?", id, uid).Delete(models.Prompt{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}
	c.JSON(http.StatusOK, models.ErrCode{ErrCode: "0000"})
}
