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
	promptsType, err := c.GetQuery("type")
	if err {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	utils.DB.Table("prompt").Where("designer = ? AND type = ?", 0, promptsType).Find(&prompts)
	c.JSON(http.StatusOK, prompts)
}
