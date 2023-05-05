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
	utils.DB.Table("prompt").Where("designer = ?", 0).Find(&prompts)
	c.JSON(http.StatusOK, prompts)
}
