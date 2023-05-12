package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

func GetMyPersonality(c *gin.Context) {
	var personality []models.Personality
	var reqBody map[string]interface{}
	var err error
	apiKey := reqBody["apiKey"].(string)
	apiKey, err = utils.Encrypt(apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3006"})
		return
	}
	err = utils.DB.Table("personality").Where("designer = ? AND uid = ?", 1, apiKey).Find(&personality).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}
	c.JSON(http.StatusOK, personality)
}
