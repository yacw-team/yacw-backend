package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

// GetPromptShop 获取prompt商店的内容
func GetPromptShop(c *gin.Context) {
	var prompts []models.Prompt
	//获取类型
	promptsType, has := c.GetQuery("type")
	if !has {
		slice := []string{promptsType}
		if !utils.Utf8Check(slice) {
			c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1011"})
			return
		}
		err := utils.DB.Table("prompt").Where("designer = ?", 0).Find(&prompts).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
			return
		}
		c.JSON(http.StatusOK, prompts)
	}
	err := utils.DB.Table("prompt").Where("designer = ? AND type = ?", 0, promptsType).Find(&prompts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}
	c.JSON(http.StatusOK, prompts)
}
