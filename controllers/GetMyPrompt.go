package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

// GetMyPrompt 获取我的prompt
func GetMyPrompt(c *gin.Context) {
	var prompts []models.Prompt
	utils.DB.Table("prompt").Where("designer = ?", 1).Find(&prompts)
	c.JSON(http.StatusOK, prompts)
}
