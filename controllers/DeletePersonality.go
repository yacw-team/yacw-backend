package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

type deletePersonalityReqBody struct {
	apikey        string
	personalityId string
}

// DeletePersonality 删除用户创建的Personality
func DeletePersonality(c *gin.Context) {

	var reqBody deletePromptReqBody
	err := c.BindJSON(reqBody)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "数据绑定错误")
		return
	}

	uid := utils.Encrypt(reqBody.apiKey) //用户id

	id := reqBody.promptsId

	err = utils.DB.Table("personality").Where("id = ? AND uid = ?", id, uid).Delete(models.Personality{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "NO")
		return
	}
	c.JSON(http.StatusOK, "OK")
}
