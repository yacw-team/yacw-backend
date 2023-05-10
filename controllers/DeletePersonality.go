package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

// DeletePersonality 删除用户创建的Personality
func DeletePersonality(c *gin.Context) {

	var reqBody deletePromptReqBody
	err := c.BindJSON(&reqBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1010"})
		return
	}
	apiKeyCheck := utils.IsValidApiKey(reqBody.ApiKey)
	if apiKeyCheck {
		if apiKeyCheck {
			uid := utils.Encrypt(reqBody.ApiKey) //用户id

			id := reqBody.PromptsId

			err = utils.DB.Table("personality").Where("id = ? AND uid = ?", id, uid).Delete(models.Personality{}).Error
			if err != nil {
				c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "3009"})
				return
			}
			c.JSON(http.StatusOK, models.ErrCode{ErrCode: "0000"})
		}
	} else {
		var errCode models.ErrCode
		errCode.ErrCode = "3004"
		c.JSON(http.StatusBadRequest, errCode)
	}
}
