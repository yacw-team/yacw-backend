package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strings"
)

// CreatePersonality 用户创建personality
func CreatePersonality(c *gin.Context) {
	uid := utils.Encrypt(c.PostForm("apiKey"))
	modelName := c.PostForm("name")
	description := c.PostForm("description")
	prompts := c.PostForm("prompts")

	if len(strings.TrimSpace(modelName)) == 0 {
		c.JSON(http.StatusBadRequest, "名称不能为空")
	}

	if len(strings.TrimSpace(prompts)) == 0 {
		c.JSON(http.StatusBadRequest, "prompts不能为空")
	}

	personality := models.Personality{
		Uid:         uid,
		ModelName:   modelName,
		Description: description,
		Prompts:     prompts,
	}

	err := utils.DB.Table("personality").Create(&personality).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "NO")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":          personality.Id,
		"name":        personality.ModelName,
		"description": personality.Description,
		"prompts":     personality.Prompts,
	})
}
