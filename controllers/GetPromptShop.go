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
		c.JSON(http.StatusBadRequest, has)
		return
	}
	err := utils.DB.Table("prompt").Where("designer = ? AND type = ?", 0, promptsType).Find(&prompts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, "数据库查询错误")
		return
	}
	c.JSON(http.StatusOK, prompts)
}
