package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

// CreatePrompt 用户创建prompt
func CreatePrompt(c *gin.Context) {
	uid := utils.EncryptPassword(c.PostForm("apiKey"))
	modelName := c.PostForm("name")
	description := c.PostForm("description")
	prompts := c.PostForm("prompts")

	prompt := models.Prompt{
		Uid:         uid,
		ModelName:   modelName,
		Description: description,
		Prompts:     prompts,
	}

	err := utils.DB.Table("prompt").Create(&prompt).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "NO")
		return
	}
	c.JSON(http.StatusOK, "OK")
}
