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
	err := utils.DB.Table("prompt").Where("designer = ?", 1).Find(&prompts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "数据库查询错误")
		return
	}
	c.JSON(http.StatusOK, prompts)
}
