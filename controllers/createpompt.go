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
		c.JSON(http.StatusBadRequest, "名称不能为空")
	}

	if len(strings.TrimSpace(prompts)) == 0 {
		c.JSON(http.StatusBadRequest, "prompt不能为空")
	}

	prompt := models.Prompt{
		Uid:         uid,
		ModelName:   modelName,
		Description: description,
		Prompts:     prompts,
		Designer:    1,
	}

	err := utils.DB.Table("prompt").Create(&prompt).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "数据库读写出错")
		return
	}
	c.JSON(http.StatusOK, "OK")
}
