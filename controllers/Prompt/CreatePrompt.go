package Prompt

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strings"
)

// CreatePrompt 用户创建prompt
func CreatePrompt(c *gin.Context) {
	var err error
	var reqBody map[string]interface{}
	reqTemp, ok := c.Get("reqBody")
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}
	reqBody = reqTemp.(map[string]interface{})

	apiKey, ok := reqBody["apiKey"].(string)

	modelName, ok := reqBody["name"].(string)

	description, ok := reqBody["description"].(string)

	prompts, ok := reqBody["prompts"].(string)

	//检测utf-8编码

	apiKeyCheck := utils.IsValidApiKey(apiKey)
	if !apiKeyCheck {
		var errCode models.ErrCode
		errCode.ErrCode = "3004"
		c.JSON(http.StatusBadRequest, errCode)
		return
	}

	uid, err := utils.Encrypt(apiKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3006"})
		return
	}

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

	err = utils.DB.Table("prompt").Create(&prompt).Error
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
