package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

type deletePromptReqBody struct {
	apiKey    string
	promptsId string
}

// DeletePrompt 删除用户创建的prompt
func DeletePrompt(c *gin.Context) {
	var reqBody deletePromptReqBody
	err := c.Bind(&reqBody)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "数据绑定失败")
		return
	}

	uid := utils.Encrypt(reqBody.apiKey) //用户id
	id := reqBody.promptsId

	err = utils.DB.Table("prompt").Where("id = ? AND uid = ?", id, uid).Delete(models.Prompt{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "NO")
		return
	}
	c.JSON(http.StatusOK, "OK")
}
