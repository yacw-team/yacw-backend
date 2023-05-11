package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

type deletePersonalityReqBody struct {
	ApiKey        string `json:"apiKey"`
	PersonalityId string `json:"personalityId"`
}

// DeletePersonality 删除用户创建的Personality
func DeletePersonality(c *gin.Context) {

	var err error
	var reqBody deletePersonalityReqBody
	err = c.BindJSON(&reqBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1010"})
		return
	}
	//检测utf-8编码
	slice := []string{reqBody.ApiKey, reqBody.PersonalityId}
	if !utils.Utf8Check(slice) {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1011"})
		return
	}
	apiKeyCheck := utils.IsValidApiKey(reqBody.ApiKey)
	if apiKeyCheck {
		if apiKeyCheck {
			uid, err := utils.Encrypt(reqBody.ApiKey) //用户id
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3006"})
				return
			}

			id := reqBody.PersonalityId

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
