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
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2007"})
			// 进行适当的处理
		}
	}()
	var err error
	var reqBody map[string]interface{}
	reqTemp, ok := c.Get("reqBody")
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}
	reqBody = reqTemp.(map[string]interface{})

	uid := reqBody["apiKey"].(string)
	modelName := reqBody["name"].(string)
	description := reqBody["description"].(string)
	prompts := reqBody["prompts"].(string)

	//检测utf-8编码
	slice := []string{uid, modelName, description, prompts}
	if !utils.Utf8Check(slice) {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1011"})
		return
	}

	uid, err = utils.Encrypt(uid)

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
