package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strings"
)

// CreatePrompt 用户创建prompt
func CreatePrompt(c *gin.Context) {
	uid := utils.Encrypt(c.PostForm("apiKey"))
	modelName := c.PostForm("name")
	description := c.PostForm("description")
	prompts := c.PostForm("prompts")

	if len(strings.TrimSpace(modelName)) == 0 {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1007"})
		return
	}

	if len(strings.TrimSpace(prompts)) == 0 {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1007"})
		return
	}

	prompt := models.Prompt{
		Uid:         uid,
		PromptName:  modelName,
		Description: description,
		Prompts:     prompts,
		Designer:    1,
	}

	err := utils.DB.Table("prompt").Create(&prompt).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "3009"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":          prompt.Id,
		"name":        prompt.PromptName,
		"description": prompt.Description,
		"prompts":     prompt.Prompts,
	})
}
